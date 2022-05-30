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

package proto

import (
	"context"

	"github.com/cectc/hptx/pkg/api"
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
}
