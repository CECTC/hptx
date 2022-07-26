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

package grpc

import (
	"context"
	"strings"

	"github.com/cectc/dbpack/pkg/log"
	"google.golang.org/grpc"

	"github.com/cectc/hptx/pkg/core"
)

const XID = "xid"

type GlobalTransactionInfo struct {
	FullMethod string
	Timeout    int32
}

func GlobalTransactionInterceptor(globalTransactionInfos []*GlobalTransactionInfo) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var xid string
		transactionManager := core.GetDistributedTransactionManager()
		for _, gs := range globalTransactionInfos {
			if strings.EqualFold(gs.FullMethod, info.FullMethod) {
				xid, err = transactionManager.Begin(ctx, gs.FullMethod, gs.Timeout)
				if err != nil {
					return nil, err
				}
				ctx = context.WithValue(ctx, XID, xid)
				resp, err = handler(ctx, req)
				if err == nil {
					_, commitErr := core.GetDistributedTransactionManager().Commit(ctx, xid)
					if err != nil {
						log.Error(err)
						return resp, commitErr
					}
				} else {
					_, rollbackErr := core.GetDistributedTransactionManager().Rollback(ctx, xid)
					if rollbackErr != nil {
						log.Error(rollbackErr)
					}
				}
				return resp, err
			}
		}
		return handler(ctx, req)
	}
}
