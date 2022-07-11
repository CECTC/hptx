/*
 * This file is part of the hptx distribution (https://github.com/cectc/htpx).
 * Copyright 2022 CECTC, Inc.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
 * PARTICULAR PURPOSE. See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this
 * program. If not, see <https://www.gnu.org/licenses/>.
 */

package tm

import (
	"fmt"

	"github.com/cectc/dbpack/pkg/log"
	"github.com/pkg/errors"

	"github.com/cectc/hptx/pkg/api"
	ctx "github.com/cectc/hptx/pkg/base/context"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/core"
	err2 "github.com/cectc/hptx/pkg/errors"
)

const (
	DefaultGlobalTxTimeout = 60000
	DefaultGlobalTxName    = "default"
)

type SuspendedResourcesHolder struct {
	Xid string
}

type GlobalTransaction interface {
	Begin(ctx *ctx.RootContext) error
	BeginWithTimeout(timeout int32, ctx *ctx.RootContext) error
	BeginWithTimeoutAndName(timeout int32, name string, ctx *ctx.RootContext) error
	Commit(ctx *ctx.RootContext) error
	Rollback(ctx *ctx.RootContext) error
	Suspend(unbindXid bool, ctx *ctx.RootContext) (*SuspendedResourcesHolder, error)
	Resume(suspendedResourcesHolder *SuspendedResourcesHolder, ctx *ctx.RootContext) error
	GetStatus(ctx *ctx.RootContext) (api.GlobalSession_GlobalStatus, error)
	GetXid(ctx *ctx.RootContext) string
	GlobalReport(globalStatus api.GlobalSession_GlobalStatus, ctx *ctx.RootContext) error
	GetLocalStatus() api.GlobalSession_GlobalStatus
}

type GlobalTransactionRole byte

const (
	// The Launcher. The one begins the current global transaction.
	Launcher GlobalTransactionRole = iota

	// The Participant. The one just joins into a existing global transaction.
	Participant
)

func (role GlobalTransactionRole) String() string {
	switch role {
	case Launcher:
		return "Launcher"
	case Participant:
		return "Participant"
	default:
		return fmt.Sprintf("%d", role)
	}
}

type DefaultGlobalTransaction struct {
	conf   config.TMConfig
	XID    string
	Status api.GlobalSession_GlobalStatus
	Role   GlobalTransactionRole
}

func (gtx *DefaultGlobalTransaction) Begin(ctx *ctx.RootContext) error {
	return gtx.BeginWithTimeout(DefaultGlobalTxTimeout, ctx)
}

func (gtx *DefaultGlobalTransaction) BeginWithTimeout(timeout int32, ctx *ctx.RootContext) error {
	return gtx.BeginWithTimeoutAndName(timeout, DefaultGlobalTxName, ctx)
}

func (gtx *DefaultGlobalTransaction) BeginWithTimeoutAndName(timeout int32, name string, ctx *ctx.RootContext) error {
	if gtx.Role != Launcher {
		if gtx.XID == "" {
			return errors.New("xid should not be empty")
		}
		log.Debugf("Ignore Begin(): just involved in global transaction [%s]", gtx.XID)
		return nil
	}
	if gtx.XID != "" {
		return errors.New("xid should be empty")
	}
	if ctx.InGlobalTransaction() {
		return errors.New("xid should be empty")
	}
	xid, err := core.GetDistributedTransactionManager().Begin(ctx, name, timeout)
	if err != nil {
		return errors.WithStack(err)
	}
	gtx.XID = xid
	gtx.Status = api.Begin
	ctx.Bind(xid)
	log.Infof("begin new global transaction [%s]", xid)
	return nil
}

func (gtx *DefaultGlobalTransaction) Commit(ctx *ctx.RootContext) error {
	defer func() {
		ctxXid := ctx.GetXID()
		if ctxXid != "" && gtx.XID == ctxXid {
			ctx.Unbind()
		}
	}()
	if gtx.Role == Participant {
		log.Debugf("ignore Commit(): just involved in global transaction [%s]", gtx.XID)
		return nil
	}
	if gtx.XID == "" {
		return errors.New("xid should not be empty")
	}
	retry := gtx.conf.CommitRetryCount
	for retry > 0 {
		status, err := core.GetDistributedTransactionManager().Commit(ctx, gtx.XID)
		if err != nil {
			if errors.Is(err, err2.GlobalTransactionFinished) {
				return err
			}
			log.Errorf("failed to report global commit [%s],Retry Countdown: %d, reason: %s", gtx.XID, retry, err.Error())
		} else {
			gtx.Status = status
			break
		}
		retry--
		if retry == 0 {
			return errors.New("Failed to report global commit")
		}
	}
	log.Infof("[%s] commit status: %s", gtx.XID, gtx.Status.String())
	return nil
}

func (gtx *DefaultGlobalTransaction) Rollback(ctx *ctx.RootContext) error {
	defer func() {
		ctxXid := ctx.GetXID()
		if ctxXid != "" && gtx.XID == ctxXid {
			ctx.Unbind()
		}
	}()
	if gtx.Role == Participant {
		log.Debugf("ignore Rollback(): just involved in global transaction [%s]", gtx.XID)
		return nil
	}
	if gtx.XID == "" {
		return errors.New("xid should not be empty")
	}
	retry := gtx.conf.RollbackRetryCount
	for retry > 0 {
		status, err := core.GetDistributedTransactionManager().Rollback(ctx, gtx.XID)
		if err != nil {
			if errors.Is(err, err2.GlobalTransactionFinished) {
				return err
			}
			log.Errorf("failed to report global rollback [%s],Retry Countdown: %d, reason: %s", gtx.XID, retry, err.Error())
		} else {
			gtx.Status = status
			break
		}
		retry--
		if retry == 0 {
			return errors.New("Failed to report global rollback")
		}
	}
	log.Infof("[%s] rollback status: %s", gtx.XID, gtx.Status.String())
	return nil
}

func (gtx *DefaultGlobalTransaction) Suspend(unbindXid bool, ctx *ctx.RootContext) (*SuspendedResourcesHolder, error) {
	xid := ctx.GetXID()
	if xid != "" && unbindXid {
		ctx.Unbind()
		log.Debugf("suspending current transaction,xid = %s", xid)
	}
	return &SuspendedResourcesHolder{Xid: xid}, nil
}

func (gtx *DefaultGlobalTransaction) Resume(suspendedResourcesHolder *SuspendedResourcesHolder, ctx *ctx.RootContext) error {
	if suspendedResourcesHolder == nil {
		return nil
	}
	xid := suspendedResourcesHolder.Xid
	if xid != "" {
		ctx.Bind(xid)
		log.Debugf("resuming the transaction,xid = %s", xid)
	}
	return nil
}

func (gtx *DefaultGlobalTransaction) GetXid(ctx *ctx.RootContext) string {
	return gtx.XID
}

func (gtx *DefaultGlobalTransaction) GetLocalStatus() api.GlobalSession_GlobalStatus {
	return gtx.Status
}

func CreateNew() *DefaultGlobalTransaction {
	return &DefaultGlobalTransaction{
		conf:   config.GetTMConfig(),
		XID:    "",
		Status: api.Begin,
		Role:   Launcher,
	}
}

func GetCurrent(ctx *ctx.RootContext) *DefaultGlobalTransaction {
	xid := ctx.GetXID()
	if xid == "" {
		return nil
	}
	return &DefaultGlobalTransaction{
		conf:   config.GetTMConfig(),
		XID:    xid,
		Status: api.Begin,
		Role:   Participant,
	}
}

func GetCurrentOrCreate(ctx *ctx.RootContext) *DefaultGlobalTransaction {
	tx := GetCurrent(ctx)
	if tx == nil {
		return CreateNew()
	}
	return tx
}
