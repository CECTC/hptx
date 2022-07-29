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

package proto

import (
	"context"

	"github.com/cectc/dbpack/pkg/dt/api"
)

type BranchResource interface {
	Commit(ctx context.Context, bs *api.BranchSession) (api.BranchSession_BranchStatus, error)
	Rollback(ctx context.Context, bs *api.BranchSession) (api.BranchSession_BranchStatus, error)
}

type TransactionManager interface {
	// Begin return xid
	Begin(ctx context.Context, transactionName string, timeout int32) (string, error)
	Commit(ctx context.Context, xid string) (api.GlobalSession_GlobalStatus, error)
	Rollback(ctx context.Context, xid string) (api.GlobalSession_GlobalStatus, error)
	BranchRegister(ctx context.Context, in *BranchRegister) (string, int64, error)
	BranchReport(ctx context.Context, branchID string, status api.BranchSession_BranchStatus) error
	ReleaseLockKeys(ctx context.Context, resourceID string, lockKeys []string) (bool, error)
	IsLockable(ctx context.Context, resourceID, lockKey string) (bool, error)
	IsLockableWithXID(ctx context.Context, resourceID, lockKey, xid string) (bool, error)
}
