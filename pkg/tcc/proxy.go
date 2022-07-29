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
	"encoding/json"
	"reflect"

	"github.com/cectc/dbpack/pkg/dt/api"
	"github.com/cectc/dbpack/pkg/log"
	"github.com/cectc/dbpack/pkg/misc"
	gxnet "github.com/dubbogo/gost/net"
	"github.com/pkg/errors"

	ctx "github.com/cectc/hptx/pkg/base/context"
	"github.com/cectc/hptx/pkg/core"
	"github.com/cectc/hptx/pkg/proto"
	"github.com/cectc/hptx/pkg/proxy"
)

var (
	ActionNameTag = "TccActionName"

	TryMethod     = "Try"
	ConfirmMethod = "Confirm"
	CancelMethod  = "Cancel"

	ActionStartTime = "action-start-time"
	ActionName      = "actionName"
	PrepareMethod   = "sys::prepare"
	CommitMethod    = "sys::commit"
	RollbackMethod  = "sys::rollback"
	HostName        = "host-name"

	businessActionContextType = reflect.TypeOf(&ctx.BusinessActionContext{})
)

type Service interface {
	Try(ctx *ctx.BusinessActionContext) (bool, error)
	Confirm(ctx *ctx.BusinessActionContext) bool
	Cancel(ctx *ctx.BusinessActionContext) bool
}

type ProxyService interface {
	GetService() Service
}

func ImplementTCC(v ProxyService) {
	valueOf := reflect.ValueOf(v)
	log.Debugf("[implement] reflect.TypeOf: %s", valueOf.String())

	valueOfElem := valueOf.Elem()
	typeOf := valueOfElem.Type()

	// check incoming interface, incoming interface's elem must be a struct.
	if typeOf.Kind() != reflect.Struct {
		log.Errorf("%s must be a struct ptr", valueOf.String())
		return
	}
	proxyService := v.GetService()
	makeCallProxy := func(methodDesc *proxy.MethodDescriptor, resource *Resource) func(in []reflect.Value) []reflect.Value {
		return func(in []reflect.Value) []reflect.Value {
			businessContextValue := in[0]
			businessActionContext := businessContextValue.Interface().(*ctx.BusinessActionContext)
			rootContext := businessActionContext.RootContext
			businessActionContext.XID = rootContext.GetXID()
			businessActionContext.ActionName = resource.ActionName
			if !rootContext.InGlobalTransaction() {
				args := make([]interface{}, 0)
				args = append(args, businessActionContext)
				return proxy.Invoke(methodDesc, nil, args)
			}

			returnValues, err := proceed(methodDesc, businessActionContext, resource)
			if err != nil {
				return proxy.ReturnWithError(methodDesc, errors.WithStack(err))
			}
			return returnValues
		}
	}

	numField := valueOfElem.NumField()
	for i := 0; i < numField; i++ {
		t := typeOf.Field(i)
		methodName := t.Name
		f := valueOfElem.Field(i)
		if f.Kind() == reflect.Func && f.IsValid() && f.CanSet() && methodName == TryMethod {
			if t.Type.NumIn() != 1 && t.Type.In(0) != businessActionContextType {
				panic("prepare method argument is not BusinessActionContext")
			}

			actionName := t.Tag.Get(ActionNameTag)
			if actionName == "" {
				panic("must tag TccActionName")
			}

			commitMethodDesc := proxy.Register(proxyService, ConfirmMethod)
			cancelMethodDesc := proxy.Register(proxyService, CancelMethod)
			tryMethodDesc := proxy.Register(proxyService, methodName)

			tccResource := &Resource{
				ActionName:         actionName,
				PrepareMethodName:  TryMethod,
				CommitMethodName:   ConfirmMethod,
				CommitMethod:       commitMethodDesc,
				RollbackMethodName: CancelMethod,
				RollbackMethod:     cancelMethodDesc,
			}

			tccResourceManager.RegisterResource(tccResource)

			// do method proxy here:
			f.Set(reflect.MakeFunc(f.Type(), makeCallProxy(tryMethodDesc, tccResource)))
			log.Debugf("set method [%s]", methodName)
		}
	}
}

func proceed(methodDesc *proxy.MethodDescriptor, ctx *ctx.BusinessActionContext, resource *Resource) ([]reflect.Value, error) {
	var (
		args = make([]interface{}, 0)
	)

	branchID, branchSessionID, err := doTccActionLogStore(ctx, resource)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ctx.BranchSessionID = branchSessionID

	args = append(args, ctx)
	returnValues := proxy.Invoke(methodDesc, nil, args)
	errValue := returnValues[len(returnValues)-1]
	if errValue.IsValid() && !errValue.IsNil() {
		err := core.GetDistributedTransactionManager().BranchReport(ctx.Context, branchID, api.PhaseOneFailed)
		if err != nil {
			log.Errorf("branch report err: %v", err)
		}
	}
	return returnValues, nil
}

func doTccActionLogStore(ctx *ctx.BusinessActionContext, resource *Resource) (string, int64, error) {
	ctx.ActionContext[ActionStartTime] = misc.CurrentTimeMillis()
	ctx.ActionContext[ActionName] = ctx.ActionName
	ctx.ActionContext[PrepareMethod] = resource.PrepareMethodName
	ctx.ActionContext[CommitMethod] = resource.CommitMethodName
	ctx.ActionContext[RollbackMethod] = resource.RollbackMethodName
	ip, err := gxnet.GetLocalIP()
	if err == nil {
		ctx.ActionContext[HostName] = ip
	} else {
		log.Warn("getLocalIP error")
	}

	applicationContext := make(map[string]interface{})
	applicationContext[ActionContext] = ctx.ActionContext

	applicationData, err := json.Marshal(applicationContext)
	if err != nil {
		log.Errorf("marshal applicationContext failed:%v", applicationContext)
		return "", 0, err
	}

	branchID, branchSessionID, err := core.GetDistributedTransactionManager().BranchRegister(ctx.Context, &proto.BranchRegister{
		XID:             ctx.XID,
		ResourceID:      resource.GetResourceID(),
		LockKey:         "",
		BranchType:      resource.GetBranchType(),
		ApplicationData: applicationData,
	})
	if err != nil {
		log.Errorf("TCC branch Register error, xid: %s", ctx.XID)
		return "", 0, errors.WithStack(err)
	}
	return branchID, branchSessionID, nil
}
