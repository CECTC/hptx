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

package gin

import (
	"net/http"

	"github.com/cectc/dbpack/pkg/log"
	"github.com/gin-gonic/gin"

	"github.com/cectc/hptx/pkg/core"
)

const XID = "xid"

func GlobalTransaction(timeout int32) gin.HandlerFunc {
	return func(context *gin.Context) {
		var err error
		transactionManager := core.GetDistributedTransactionManager()
		xid, err := transactionManager.Begin(context, context.FullPath(), timeout)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
		}
		context.Set(XID, xid)
		context.Next()
		if context.Writer.Status() == http.StatusOK && len(context.Errors) == 0 {
			_, err = core.GetDistributedTransactionManager().Commit(context, xid)
			if err != nil {
				log.Error(err)
				context.AbortWithError(http.StatusInternalServerError, err)
			}
		} else {
			_, err = core.GetDistributedTransactionManager().Rollback(context, xid)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
