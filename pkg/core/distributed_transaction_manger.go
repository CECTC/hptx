/*
* Copyright 2022 CECTC, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package core

import (
	"context"
	"fmt"
	"time"

	"github.com/cectc/dbpack/pkg/log"
	"github.com/cectc/dbpack/pkg/misc"
	"github.com/cectc/dbpack/pkg/misc/uuid"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"

	"github.com/cectc/hptx/pkg/api"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/proto"
	"github.com/cectc/hptx/pkg/resource"
	"github.com/cectc/hptx/pkg/storage"
)

const DefaultRetryDeadThreshold = 130 * 1000

var manager *DistributedTransactionManager

func InitDistributedTransactionManager(conf *config.DistributedTransaction) {
	if conf.RetryDeadThreshold == 0 {
		conf.RetryDeadThreshold = DefaultRetryDeadThreshold
	}
	manager = &DistributedTransactionManager{
		applicationID:                    conf.ApplicationID,
		retryDeadThreshold:               conf.RetryDeadThreshold,
		rollbackRetryTimeoutUnlockEnable: conf.RollbackRetryTimeoutUnlockEnable,

		globalSessionQueue: workqueue.NewDelayingQueue(),
		branchSessionQueue: workqueue.New(),
	}
	driver := storage.GetStorageDriver()
	if driver == nil {
		log.Fatal("must init storage driver first")
	}
	manager.storageDriver = driver
	go func() {
		if driver.LeaderElection(manager.applicationID) {
			if err := manager.processGlobalSessions(); err != nil {
				log.Fatal(err)
			}
			if err := manager.processBranchSessions(); err != nil {
				log.Fatal(err)
			}
			go manager.processGlobalSessionQueue()
			go manager.processBranchSessionQueue()
			go manager.watchBranchSession()
		}
	}()
}

func GetDistributedTransactionManager() proto.TransactionManager {
	return manager
}

type DistributedTransactionManager struct {
	applicationID                    string
	retryDeadThreshold               int64
	rollbackRetryTimeoutUnlockEnable bool

	storageDriver storage.Driver

	globalSessionQueue workqueue.DelayingInterface
	branchSessionQueue workqueue.Interface
}

func (manager *DistributedTransactionManager) Begin(ctx context.Context, transactionName string, timeout int32) (string, error) {
	transactionID := uuid.NextID()
	xid := fmt.Sprintf("gs/%s/%d", manager.applicationID, transactionID)
	gt := &api.GlobalSession{
		XID:             xid,
		ApplicationID:   manager.applicationID,
		TransactionID:   transactionID,
		TransactionName: transactionName,
		Timeout:         timeout,
		BeginTime:       int64(misc.CurrentTimeMillis()),
		Status:          api.Begin,
	}
	if err := manager.storageDriver.AddGlobalSession(ctx, gt); err != nil {
		return "", err
	}
	manager.globalSessionQueue.AddAfter(gt, time.Duration(timeout)*time.Millisecond)
	log.Infof("successfully begin global transaction xid = {}", gt.XID)
	return xid, nil
}

func (manager *DistributedTransactionManager) Commit(ctx context.Context, xid string) (api.GlobalSession_GlobalStatus, error) {
	return manager.storageDriver.GlobalCommit(ctx, xid)
}

func (manager *DistributedTransactionManager) Rollback(ctx context.Context, xid string) (api.GlobalSession_GlobalStatus, error) {
	return manager.storageDriver.GlobalRollback(ctx, xid)
}

func (manager *DistributedTransactionManager) BranchRegister(ctx context.Context, in *proto.BranchRegister) (string, int64, error) {
	branchSessionID := uuid.NextID()
	branchID := fmt.Sprintf("bs/%s/%d", manager.applicationID, branchSessionID)
	transactionID := misc.GetTransactionID(in.XID)
	bs := &api.BranchSession{
		BranchID:        branchID,
		ApplicationID:   manager.applicationID,
		BranchSessionID: branchSessionID,
		XID:             in.XID,
		TransactionID:   transactionID,
		ResourceID:      in.ResourceID,
		LockKey:         in.LockKey,
		Type:            in.BranchType,
		Status:          api.Registered,
		ApplicationData: in.ApplicationData,
		BeginTime:       int64(misc.CurrentTimeMillis()),
	}

	if err := manager.storageDriver.AddBranchSession(ctx, bs); err != nil {
		return "", 0, err
	}
	return branchID, branchSessionID, nil
}

func (manager *DistributedTransactionManager) BranchReport(ctx context.Context, branchID string, status api.BranchSession_BranchStatus) error {
	return manager.storageDriver.BranchReport(ctx, branchID, status)
}

func (manager *DistributedTransactionManager) ReleaseLockKeys(ctx context.Context, resourceID string, lockKeys []string) (bool, error) {
	return manager.storageDriver.ReleaseLockKeys(ctx, resourceID, lockKeys)
}

func (manager *DistributedTransactionManager) IsLockable(ctx context.Context, resourceID, lockKey string) (bool, error) {
	return manager.storageDriver.IsLockable(ctx, resourceID, lockKey)
}

func (manager *DistributedTransactionManager) branchCommit(ctx context.Context, bs *api.BranchSession) (api.BranchSession_BranchStatus, error) {
	if bs.Type == api.TCC {
		return resource.GetTCCBranchResource().Commit(ctx, bs)
	}
	if bs.Type == api.AT {
		status, err := resource.GetATBranchResource().Commit(ctx, bs)
		if status == api.Complete {
			if err := manager.storageDriver.DeleteBranchSession(context.Background(), bs.BranchID); err != nil {
				log.Error(err)
			}
		}
		return status, err
	}
	return bs.Status, errors.Errorf("unsupport branch type: %s", bs.Type)
}

func (manager *DistributedTransactionManager) branchRollback(ctx context.Context, bs *api.BranchSession) (api.BranchSession_BranchStatus, error) {
	if bs.Type == api.TCC {
		return resource.GetTCCBranchResource().Rollback(ctx, bs)
	}
	if bs.Type == api.AT {
		status, err := resource.GetATBranchResource().Rollback(ctx, bs)
		if status == api.Complete {
			if err := manager.storageDriver.DeleteBranchSession(context.Background(), bs.BranchID); err != nil {
				log.Error(err)
			}
		}
		return status, err
	}
	return bs.Status, errors.Errorf("unsupport branch type: %s", bs.Type)
}

func (manager *DistributedTransactionManager) processGlobalSessions() error {
	globalSessions, err := manager.storageDriver.ListGlobalSession(context.Background(), manager.applicationID)
	if err != nil {
		return err
	}
	for _, gs := range globalSessions {
		if gs.Status == api.Begin {
			if isGlobalSessionTimeout(gs) {
				if _, err := manager.Rollback(context.Background(), gs.XID); err != nil {
					return err
				}
			}
			manager.globalSessionQueue.AddAfter(gs, time.Duration(misc.CurrentTimeMillis()-uint64(gs.BeginTime))*time.Millisecond)
		}
		if gs.Status == api.Committing || gs.Status == api.Rollbacking {
			bsKeys, err := manager.storageDriver.GetBranchSessionKeys(context.Background(), gs.XID)
			if err != nil {
				return err
			}
			if len(bsKeys) == 0 {
				if err := manager.storageDriver.DeleteGlobalSession(context.Background(), gs.XID); err != nil {
					return err
				}
				log.Debugf("global session finished, key: %s", gs.XID)
			}
		}
	}
	return nil
}

func (manager *DistributedTransactionManager) processGlobalSessionQueue() {
	for manager.processNextGlobalSession(context.Background()) {
	}
}

func (manager *DistributedTransactionManager) processNextGlobalSession(ctx context.Context) bool {
	obj, shutdown := manager.globalSessionQueue.Get()
	if shutdown {
		// Stop working
		return false
	}

	// We call Done here so the workqueue knows we have finished
	// processing this item. We also must remember to call Forget if we
	// do not want this work item being re-queued. For example, we do
	// not call Forget if a transient error occurs, instead the item is
	// put back on the workqueue and attempted again after a back-off
	// period.
	defer manager.globalSessionQueue.Done(obj)

	gs := obj.(*api.GlobalSession)
	newGlobalSession, err := manager.storageDriver.GetGlobalSession(ctx, gs.XID)
	if err != nil {
		log.Error(err)
		return true
	}
	if newGlobalSession.Status == api.Begin {
		if isGlobalSessionTimeout(newGlobalSession) {
			_, err := manager.Rollback(context.Background(), newGlobalSession.XID)
			if err != nil {
				log.Error(err)
			}
		}
	}
	if newGlobalSession.Status == api.Committing || newGlobalSession.Status == api.Rollbacking {
		bsKeys, err := manager.storageDriver.GetBranchSessionKeys(context.Background(), newGlobalSession.XID)
		if err != nil {
			log.Error(err)
		}
		if len(bsKeys) == 0 {
			if err := manager.storageDriver.DeleteGlobalSession(context.Background(), newGlobalSession.XID); err != nil {
				log.Error(err)
			}
			log.Debugf("global session finished, key: %s", newGlobalSession.XID)
		}
	}
	return true
}

func (manager *DistributedTransactionManager) processBranchSessions() error {
	branchSessions, err := manager.storageDriver.ListBranchSession(context.Background(), manager.applicationID)
	if err != nil {
		return err
	}
	for _, bs := range branchSessions {
		switch bs.Status {
		case api.Registered:
		case api.PhaseOneFailed:
			if err := manager.storageDriver.DeleteBranchSession(context.Background(), bs.BranchID); err != nil {
				return err
			}
		case api.PhaseTwoCommitting:
			manager.branchSessionQueue.Add(bs)
		case api.PhaseTwoRollbacking:
			if manager.IsRollingBackDead(bs) {
				log.Debugf("branch session rollback dead, key: %s, lock key: %s", bs.BranchID, bs.LockKey)
				if manager.rollbackRetryTimeoutUnlockEnable {
					log.Debugf("lock key: %s released", bs.BranchID, bs.LockKey)
					if _, err := manager.storageDriver.ReleaseLockKeys(context.Background(), bs.ResourceID, []string{bs.LockKey}); err != nil {
						return err
					}
				}
			} else {
				manager.branchSessionQueue.Add(bs)
			}
		}
	}
	return nil
}

func (manager *DistributedTransactionManager) processBranchSessionQueue() {
	for manager.processNextBranchSession(context.Background()) {
	}
}

func (manager *DistributedTransactionManager) processNextBranchSession(ctx context.Context) bool {
	obj, shutdown := manager.branchSessionQueue.Get()
	if shutdown {
		// Stop working
		return false
	}

	// We call Done here so the workqueue knows we have finished
	// processing this item. We also must remember to call Forget if we
	// do not want this work item being re-queued. For example, we do
	// not call Forget if a transient error occurs, instead the item is
	// put back on the workqueue and attempted again after a back-off
	// period.
	defer manager.branchSessionQueue.Done(obj)

	bs := obj.(*api.BranchSession)
	if bs.Status == api.PhaseTwoCommitting {
		status, err := manager.branchCommit(ctx, bs)
		if err != nil {
			log.Error(err)
			manager.branchSessionQueue.Add(obj)
		}
		if status != api.Complete {
			manager.branchSessionQueue.Add(obj)
		}
	}
	if bs.Status == api.PhaseTwoRollbacking {
		if manager.IsRollingBackDead(bs) {
			if manager.rollbackRetryTimeoutUnlockEnable {
				if _, err := manager.storageDriver.ReleaseLockKeys(context.Background(), bs.ResourceID, []string{bs.LockKey}); err != nil {
					log.Error(err)
				}
			}
		} else {
			status, err := manager.branchRollback(ctx, bs)
			if err != nil {
				log.Error(err)
				manager.branchSessionQueue.Add(obj)
			}
			if status != api.Complete {
				manager.branchSessionQueue.Add(obj)
			}
		}
	}
	return true
}

func (manager *DistributedTransactionManager) watchBranchSession() {
	watcher := manager.storageDriver.WatchBranchSessions(context.Background(), manager.applicationID)
	for {
		bs := <-watcher.ResultChan()
		manager.branchSessionQueue.Add(bs)
	}
}

func isGlobalSessionTimeout(gs *api.GlobalSession) bool {
	return (misc.CurrentTimeMillis() - uint64(gs.BeginTime)) > uint64(gs.Timeout)
}

func (manager *DistributedTransactionManager) IsRollingBackDead(bs *api.BranchSession) bool {
	return (misc.CurrentTimeMillis() - uint64(bs.BeginTime)) > uint64(manager.retryDeadThreshold)
}
