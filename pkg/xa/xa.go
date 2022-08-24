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

package xa

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cectc/dbpack/pkg/dt/api"
	"github.com/cectc/dbpack/pkg/log"
	"github.com/pkg/errors"

	"github.com/cectc/hptx/pkg/constant"
	"github.com/cectc/hptx/pkg/core"
	"github.com/cectc/hptx/pkg/proto"
)

func HandleWithXA(ctx context.Context, db *sql.DB, appid string, businessFn func(conn *sql.Conn) error) (err error) {
	xid := ctx.Value(constant.XID)
	if xid == nil {
		return errors.New("ctx must with value xid")
	}

	var branchID string
	branchID, _, err = core.GetDistributedTransactionManager().BranchRegister(ctx, &proto.BranchRegister{
		XID:             xid.(string),
		ResourceID:      appid,
		LockKey:         "",
		BranchType:      api.XA,
		ApplicationData: nil,
	})
	if err != nil {
		log.Errorf("XA branch Register error, xid: %s", xid.(string))
		return errors.WithStack(err)
	}
	defer func() {
		if err != nil {
			if reportErr := core.GetDistributedTransactionManager().BranchReport(ctx, branchID, api.PhaseOneFailed); reportErr != nil {
				log.Error(reportErr)
			}
		}
	}()

	var conn *sql.Conn
	conn, err = db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err = conn.ExecContext(ctx, fmt.Sprintf("XA START '%s'", branchID)); err != nil {
		return err
	}
	defer func() {
		if err == nil {
			_, err = conn.ExecContext(ctx, fmt.Sprintf("XA PREPARE '%s'", branchID))
		}
	}()
	defer func() {
		_, err = conn.ExecContext(ctx, fmt.Sprintf("XA END '%s'", branchID))
	}()
	if err = businessFn(conn); err != nil {
		return err
	}
	return nil
}
