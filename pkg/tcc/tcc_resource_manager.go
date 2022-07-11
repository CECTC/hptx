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

package tcc

import (
	"context"
	"encoding/json"

	"github.com/cectc/dbpack/pkg/log"
	"github.com/pkg/errors"

	"github.com/cectc/hptx/pkg/api"
	ctx "github.com/cectc/hptx/pkg/base/context"
	"github.com/cectc/hptx/pkg/proxy"
)

var (
	ActionContext = "actionContext"
)

var tccResourceManager ResourceManager

type ResourceManager struct {
	ResourceCache map[string]*Resource
}

func init() {
	tccResourceManager = ResourceManager{ResourceCache: make(map[string]*Resource)}
}

func GetResourceManager() ResourceManager {
	return tccResourceManager
}

func (resourceManager ResourceManager) Commit(ctx context.Context, bs *api.BranchSession) (api.BranchSession_BranchStatus, error) {
	tccResource := resourceManager.ResourceCache[bs.ResourceID]
	if tccResource == nil {
		log.Errorf("TCC resource is not exist, resourceID: %s", bs.ResourceID)
		return bs.Status, errors.Errorf("TCC resource is not exist, resourceID: %s", bs.ResourceID)
	}
	if tccResource.CommitMethod == nil {
		log.Errorf("TCC resource is not available, resourceID: %s", bs.ResourceID)
		return bs.Status, errors.Errorf("TCC resource is not available, resourceID: %s", bs.ResourceID)
	}

	result := false
	businessActionContext := getBusinessActionContext(bs.XID, bs.BranchSessionID, bs.ResourceID, bs.ApplicationData)
	args := make([]interface{}, 0)
	args = append(args, businessActionContext)
	returnValues := proxy.Invoke(tccResource.CommitMethod, nil, args)
	log.Debugf("TCC resource commit result : %v, xid: %s, branchSessionID: %d, resourceID: %s", returnValues, bs.XID, bs.BranchSessionID, bs.ResourceID)
	if len(returnValues) == 1 {
		result = returnValues[0].Interface().(bool)
	}
	if result {
		return api.Complete, nil
	}
	return bs.Status, nil
}

func (resourceManager ResourceManager) Rollback(ctx context.Context, bs *api.BranchSession) (api.BranchSession_BranchStatus, error) {
	tccResource := resourceManager.ResourceCache[bs.ResourceID]
	if tccResource == nil {
		log.Errorf("TCC resource is not exist, resourceID: %s", bs.ResourceID)
		return bs.Status, errors.Errorf("TCC resource is not exist, resourceID: %s", bs.ResourceID)
	}
	if tccResource.RollbackMethod == nil {
		log.Errorf("TCC resource is not available, resourceID: %s", bs.ResourceID)
		return bs.Status, errors.Errorf("TCC resource is not available, resourceID: %s", bs.ResourceID)
	}

	result := false
	businessActionContext := getBusinessActionContext(bs.XID, bs.BranchSessionID, bs.ResourceID, bs.ApplicationData)
	args := make([]interface{}, 0)
	args = append(args, businessActionContext)
	returnValues := proxy.Invoke(tccResource.RollbackMethod, nil, args)
	log.Debugf("TCC resource rollback result : %v, xid: %s, branchSessionID: %d, resourceID: %s", returnValues, bs.XID, bs.BranchSessionID, bs.ResourceID)
	if len(returnValues) == 1 {
		result = returnValues[0].Interface().(bool)
	}
	if result {
		return api.Complete, nil
	}
	return bs.Status, nil
}

func getBusinessActionContext(xid string, branchSessionID int64, resourceID string, applicationData []byte) *ctx.BusinessActionContext {
	var (
		tccContext       = make(map[string]interface{})
		actionContextMap = make(map[string]interface{})
	)
	if len(applicationData) > 0 {
		err := json.Unmarshal(applicationData, &tccContext)
		if err != nil {
			log.Errorf("getBusinessActionContext, unmarshal applicationData err=%v", err)
		}
	}

	acMap := tccContext[ActionContext]
	if acMap != nil {
		actionContextMap = acMap.(map[string]interface{})
	}

	businessActionContext := &ctx.BusinessActionContext{
		XID:             xid,
		BranchSessionID: branchSessionID,
		ActionName:      resourceID,
		ActionContext:   actionContextMap,
	}
	return businessActionContext
}

func (resourceManager ResourceManager) RegisterResource(resource *Resource) {
	resourceManager.ResourceCache[resource.GetResourceID()] = resource
}

func (resourceManager ResourceManager) UnregisterResource(resource *Resource) {
	delete(resourceManager.ResourceCache, resource.GetResourceID())
}

func (resourceManager ResourceManager) GetBranchType() api.BranchSession_BranchType {
	return api.TCC
}
