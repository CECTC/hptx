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

package context

import (
	"context"
	"fmt"
	"strings"

	"github.com/cectc/dbpack/pkg/log"

	"github.com/cectc/hptx/pkg/api"
)

const (
	KeyXID                = "TX_XID"
	KeyXIDInterceptorType = "tx-xid-interceptor-type"
	KeyGlobalLockFlag     = "TX_LOCK"
)

// RootContext store the global transaction context
type RootContext struct {
	context.Context

	// like thread local map
	localMap map[string]interface{}
}

// NewRootContext return a pointer to RootContext
func NewRootContext(ctx context.Context) *RootContext {
	rootCtx := &RootContext{
		Context:  ctx,
		localMap: make(map[string]interface{}),
	}

	xID := ctx.Value(KeyXID)
	if xID != nil {
		xid := xID.(string)
		rootCtx.Bind(xid)
	}
	return rootCtx
}

// Set store key value to RootContext
func (c *RootContext) Set(key string, value interface{}) {
	if c.localMap == nil {
		c.localMap = make(map[string]interface{})
	}
	c.localMap[key] = value
}

// Get get a value by given key from RootContext
func (c *RootContext) Get(key string) (value interface{}, exists bool) {
	value, exists = c.localMap[key]
	return
}

// GetXID from RootContext get xid
func (c *RootContext) GetXID() string {
	xID := c.localMap[KeyXID]
	xid, ok := xID.(string)
	if ok && xid != "" {
		return xid
	}

	xIDType := c.localMap[KeyXIDInterceptorType]
	xidType, success := xIDType.(string)

	if success && xidType != "" && strings.Contains(xidType, "_") {
		return strings.Split(xidType, "_")[0]
	}

	return ""
}

// GetXIDInterceptorType from RootContext get xid interceptor type
func (c *RootContext) GetXIDInterceptorType() string {
	xIDType := c.localMap[KeyXIDInterceptorType]
	xidType, _ := xIDType.(string)
	return xidType
}

// Bind bind xid with RootContext
func (c *RootContext) Bind(xid string) {
	log.Debugf("bind %s", xid)
	c.Set(KeyXID, xid)
}

// BindInterceptorType bind interceptor type with RootContext
func (c *RootContext) BindInterceptorType(xidType string) {
	if xidType != "" {
		xidTypes := strings.Split(xidType, "_")

		if len(xidTypes) == 2 {
			c.BindInterceptorTypeWithBranchType(xidTypes[0],
				api.BranchSession_BranchType(api.BranchSession_BranchType_value[xidTypes[1]]))
		}
	}
}

// BindInterceptorTypeWithBranchType bind interceptor type and branch type with RootContext
func (c *RootContext) BindInterceptorTypeWithBranchType(xid string, branchType api.BranchSession_BranchType) {
	xidType := fmt.Sprintf("%s_%s", xid, branchType.String())
	log.Debugf("bind interceptor type xid=%s branchType=%s", xid, branchType.String())
	c.Set(KeyXIDInterceptorType, xidType)
}

// BindGlobalLockFlag bind global lock flag with RootContext
func (c *RootContext) BindGlobalLockFlag() {
	log.Debug("local transaction global lock support enabled")
	c.Set(KeyGlobalLockFlag, KeyGlobalLockFlag)
}

// Unbind unbind xid with RootContext
func (c *RootContext) Unbind() string {
	xID := c.localMap[KeyXID]
	xid, ok := xID.(string)
	if ok && xid != "" {
		log.Debugf("unbind %s", xid)
		delete(c.localMap, KeyXID)
		return xid
	}
	return ""

}

// UnbindInterceptorType unbind interceptor type with RootContext
func (c *RootContext) UnbindInterceptorType() string {
	xidType := c.localMap[KeyXIDInterceptorType]
	xt, ok := xidType.(string)
	if ok && xt != "" {
		log.Debugf("unbind inteceptor type %s", xidType)
		delete(c.localMap, KeyXIDInterceptorType)
		return xt
	}
	return ""
}

// UnbindGlobalLockFlag unbind global lock flag with RootContext
func (c *RootContext) UnbindGlobalLockFlag() {
	log.Debug("unbind global lock flag")
	delete(c.localMap, KeyGlobalLockFlag)
}

// InGlobalTransaction determine whether the context is in global transaction
func (c *RootContext) InGlobalTransaction() bool {
	return c.localMap[KeyXID] != nil
}

// RequireGlobalLock return global lock flag
func (c *RootContext) RequireGlobalLock() bool {
	_, exists := c.localMap[KeyGlobalLockFlag]
	return exists
}
