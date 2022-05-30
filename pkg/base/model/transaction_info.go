/*
 * Copyright 2022 CECTC, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
