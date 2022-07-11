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

package storage

import (
	"context"

	"github.com/cectc/hptx/pkg/api"
)

var _driver Driver

type Driver interface {
	LeaderElection(applicationID string) bool
	AddGlobalSession(ctx context.Context, globalSession *api.GlobalSession) error
	AddBranchSession(ctx context.Context, branchSession *api.BranchSession) error
	GlobalCommit(ctx context.Context, xid string) (api.GlobalSession_GlobalStatus, error)
	GlobalRollback(ctx context.Context, xid string) (api.GlobalSession_GlobalStatus, error)
	GetGlobalSession(ctx context.Context, xid string) (*api.GlobalSession, error)
	ListGlobalSession(ctx context.Context, applicationID string) ([]*api.GlobalSession, error)
	DeleteGlobalSession(ctx context.Context, xid string) error
	GetBranchSession(ctx context.Context, branchID string) (*api.BranchSession, error)
	ListBranchSession(ctx context.Context, applicationID string) ([]*api.BranchSession, error)
	DeleteBranchSession(ctx context.Context, branchID string) error
	GetBranchSessionKeys(ctx context.Context, xid string) ([]string, error)
	BranchReport(ctx context.Context, branchID string, status api.BranchSession_BranchStatus) error
	IsLockable(ctx context.Context, resourceID string, lockKey string) (bool, error)
	ReleaseLockKeys(ctx context.Context, resourceID string, lockKeys []string) (bool, error)
	WatchGlobalSessions(ctx context.Context, applicationID string) Watcher
	WatchBranchSessions(ctx context.Context, applicationID string) Watcher
}

// Watcher can be implemented by anything that knows how to watch and report changes.
type Watcher interface {
	// Stop watching. Will close the channel returned by ResultChan(). Releases
	// any resources used by the watch.
	Stop()

	// ResultChan return a chan which will receive all the TransactionSessions. If an error occurs
	// or Stop() is called, the implementation will close this channel and
	// release any resources used by the watch.
	ResultChan() <-chan TransactionSession
}

type TransactionSession interface {
	Marshal() (data []byte, err error)
	Unmarshal(data []byte) error
}

func InitStorageDriver(driver Driver) {
	_driver = driver
}

func GetStorageDriver() Driver {
	return _driver
}
