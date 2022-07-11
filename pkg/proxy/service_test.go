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

package proxy

import (
	"context"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/cectc/hptx/pkg/base/model"
)

type Svc struct{}

var service = &Svc{}

type ProxyService struct {
	*Svc
	CreateSo func(ctx context.Context, rollback bool) error
}

var methodTransactionInfo = make(map[string]*model.TransactionInfo)

func (svc *ProxyService) GetMethodTransactionInfo(methodName string) *model.TransactionInfo {
	return methodTransactionInfo[methodName]
}
func (svc *ProxyService) GetProxyService() interface{} {
	return svc.Svc
}
func (svc *Svc) CreateSo(ctx context.Context, rollback bool) error {
	return nil
}

var ProxySvc = &ProxyService{
	Svc: service,
}

func TestRegister(t *testing.T) {
	ps := ProxySvc.GetProxyService()
	method, _ := reflect.TypeOf(ps).MethodByName("CreateSo")

	tests := []struct {
		name           string
		service        interface{}
		methodName     string
		expectedResult *MethodDescriptor
	}{
		{
			name:       "test register",
			service:    ps,
			methodName: "CreateSo",
			expectedResult: &MethodDescriptor{
				ArgsNum:         3,
				Method:          reflect.Method{Name: "CreateSo"},
				ReturnValuesNum: 1,
				ReturnValuesType: []reflect.Type{
					method.Type.Out(0),
				},
				ArgsType: []reflect.Type{
					method.Type.In(1),
					method.Type.In(2),
				},
				CtxType:     method.Type.In(1),
				CallerValue: reflect.ValueOf(ps),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := Register(tt.service, tt.methodName)

			assert.Equal(t, tt.expectedResult.ArgsNum, actualResult.ArgsNum)
			assert.Equal(t, tt.expectedResult.Method.Name, actualResult.Method.Name)
			assert.Equal(t, tt.expectedResult.ReturnValuesNum, actualResult.ReturnValuesNum)
			assert.Equal(t, tt.expectedResult.ReturnValuesType, actualResult.ReturnValuesType)
		})
	}
}

func TestSuiteContext(t *testing.T) {
	type key string
	const aa key = "aa"
	const bb key = "bb"

	emptyContext := context.TODO()
	validContext := context.WithValue(context.TODO(), aa, bb)

	tests := []struct {
		name           string
		methodDesc     *MethodDescriptor
		ctx            context.Context
		expectedResult reflect.Value
	}{
		{
			name:           "test suite valid context",
			methodDesc:     &MethodDescriptor{},
			ctx:            emptyContext,
			expectedResult: reflect.ValueOf(emptyContext),
		},
		{
			name: "test suite invalid context",
			methodDesc: &MethodDescriptor{
				CtxType: reflect.TypeOf(validContext),
			},
			ctx:            nil,
			expectedResult: reflect.Zero(reflect.TypeOf(validContext)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := SuiteContext(tt.ctx, tt.methodDesc)

			assert.Equal(t, tt.expectedResult, actualResult)
		})
	}
}

func TestReturnWithError(t *testing.T) {
	testingErr := errors.New("testing err")

	tests := []struct {
		name           string
		methodDesc     *MethodDescriptor
		err            error
		expectedResult []reflect.Value
	}{
		{
			name: "test return with error when methodDesc.ReturnValuesNum is 2",
			methodDesc: &MethodDescriptor{
				ReturnValuesNum: 2,
				ReturnValuesType: []reflect.Type{
					reflect.TypeOf("1"),
					reflect.TypeOf(testingErr),
				},
			},
			err: testingErr,
			expectedResult: []reflect.Value{
				reflect.Zero(reflect.TypeOf("1")),
				reflect.ValueOf(testingErr),
			},
		},
		{
			name: "test return with error when methodDesc.ReturnValuesNum is 1",
			methodDesc: &MethodDescriptor{
				ReturnValuesNum: 1,
				ReturnValuesType: []reflect.Type{
					reflect.TypeOf(testingErr),
				},
			},
			err: testingErr,
			expectedResult: []reflect.Value{
				reflect.ValueOf(testingErr),
			},
		},
		{
			name: "test return with error when methodDesc.ReturnValuesNum is 0",
			methodDesc: &MethodDescriptor{
				ReturnValuesNum:  0,
				ReturnValuesType: []reflect.Type{},
			},
			err: testingErr,
			expectedResult: []reflect.Value{
				reflect.ValueOf(testingErr),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := ReturnWithError(tt.methodDesc, tt.err)

			for idx := range tt.expectedResult {
				assert.Equal(t, tt.expectedResult[idx].Type(), actualResult[idx].Type())
			}
		})
	}
}

func TestIsExportedOrBuiltinType(t *testing.T) {
	type ExportedStruct struct {
		Abb string
	}
	type notExportedStruct struct {
		Cdd string
	}
	tests := []struct {
		name           string
		typeOf         reflect.Type
		expectedResult bool
	}{
		{
			name: "test type is exported and not builtin",
			typeOf: reflect.TypeOf(ExportedStruct{
				Abb: "testing",
			}),
			expectedResult: true,
		},
		{
			name: "test type is not exported and not builtin",
			typeOf: reflect.TypeOf(notExportedStruct{
				Cdd: "testing",
			}),
			expectedResult: false,
		},
		{
			name:           "test type is builtin",
			typeOf:         reflect.TypeOf("Abb"), // string type
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := isExportedOrBuiltinType(tt.typeOf)

			assert.Equal(t, tt.expectedResult, actualResult)
		})
	}
}
