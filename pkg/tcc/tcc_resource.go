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
	"github.com/cectc/hptx/pkg/api"
	"github.com/cectc/hptx/pkg/proxy"
)

type Resource struct {
	ActionName         string
	PrepareMethodName  string
	CommitMethodName   string
	CommitMethod       *proxy.MethodDescriptor
	RollbackMethodName string
	RollbackMethod     *proxy.MethodDescriptor
}

func (resource *Resource) GetResourceID() string {
	return resource.ActionName
}

func (resource *Resource) GetBranchType() api.BranchSession_BranchType {
	return api.TCC
}
