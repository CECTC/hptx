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

package model

import "fmt"

// Propagation transaction isolation level
type Propagation byte

const (
	Required Propagation = iota

	RequiresNew

	NotSupported

	Supports

	Never

	Mandatory
)

// String
func (t Propagation) String() string {
	switch t {
	case Required:
		return "Required"
	case RequiresNew:
		return "REQUIRES_NEW"
	case NotSupported:
		return "NOT_SUPPORTED"
	case Supports:
		return "Supports"
	case Never:
		return "Never"
	case Mandatory:
		return "Mandatory"
	default:
		return fmt.Sprintf("%d", t)
	}
}

// TransactionInfo used to configure global transaction parameters
type TransactionInfo struct {
	TimeOut     int32
	Name        string
	Propagation Propagation
}
