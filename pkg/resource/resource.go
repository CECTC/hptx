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

package resource

import (
	"github.com/cectc/hptx/pkg/api"
	"github.com/cectc/hptx/pkg/proto"
)

var branches = make(map[api.BranchSession_BranchType]proto.BranchResource)

func GetTCCBranchResource() proto.BranchResource {
	return branches[api.TCC]
}

func GetATBranchResource() proto.BranchResource {
	return branches[api.AT]
}

func InitTCCBranchResource(resource proto.BranchResource) {
	branches[api.TCC] = resource
}

func InitATBranchResource(resource proto.BranchResource) {
	branches[api.AT] = resource
}
