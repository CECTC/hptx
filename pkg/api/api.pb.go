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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ResultCode int32

const (
	ResultCodeFailed  ResultCode = 0
	ResultCodeSuccess ResultCode = 1
)

var ResultCode_name = map[int32]string{
	0: "ResultCodeFailed",
	1: "ResultCodeSuccess",
}

var ResultCode_value = map[string]int32{
	"ResultCodeFailed":  0,
	"ResultCodeSuccess": 1,
}

func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type GlobalSession_GlobalStatus int32

const (
	Begin       GlobalSession_GlobalStatus = 0
	Committing  GlobalSession_GlobalStatus = 1
	Rollbacking GlobalSession_GlobalStatus = 2
	Finished    GlobalSession_GlobalStatus = 3
)

var GlobalSession_GlobalStatus_name = map[int32]string{
	0: "Begin",
	1: "Committing",
	2: "Rollbacking",
	3: "Finished",
}

var GlobalSession_GlobalStatus_value = map[string]int32{
	"Begin":       0,
	"Committing":  1,
	"Rollbacking": 2,
	"Finished":    3,
}

func (GlobalSession_GlobalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0, 0}
}

type BranchSession_BranchType int32

const (
	AT   BranchSession_BranchType = 0
	TCC  BranchSession_BranchType = 1
	SAGA BranchSession_BranchType = 2
	XA   BranchSession_BranchType = 3
)

var BranchSession_BranchType_name = map[int32]string{
	0: "AT",
	1: "TCC",
	2: "SAGA",
	3: "XA",
}

var BranchSession_BranchType_value = map[string]int32{
	"AT":   0,
	"TCC":  1,
	"SAGA": 2,
	"XA":   3,
}

func (BranchSession_BranchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1, 0}
}

type BranchSession_BranchStatus int32

const (
	Registered          BranchSession_BranchStatus = 0
	PhaseOneFailed      BranchSession_BranchStatus = 1
	PhaseTwoCommitting  BranchSession_BranchStatus = 2
	PhaseTwoRollbacking BranchSession_BranchStatus = 3
	Complete            BranchSession_BranchStatus = 4
)

var BranchSession_BranchStatus_name = map[int32]string{
	0: "Registered",
	1: "PhaseOneFailed",
	2: "PhaseTwoCommitting",
	3: "PhaseTwoRollbacking",
	4: "Complete",
}

var BranchSession_BranchStatus_value = map[string]int32{
	"Registered":          0,
	"PhaseOneFailed":      1,
	"PhaseTwoCommitting":  2,
	"PhaseTwoRollbacking": 3,
	"Complete":            4,
}

func (BranchSession_BranchStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1, 1}
}

type GlobalSession struct {
	XID             string                     `protobuf:"bytes,1,opt,name=XID,proto3" json:"XID,omitempty"`
	ApplicationID   string                     `protobuf:"bytes,2,opt,name=ApplicationID,proto3" json:"ApplicationID,omitempty"`
	TransactionID   int64                      `protobuf:"varint,3,opt,name=TransactionID,proto3" json:"TransactionID,omitempty"`
	TransactionName string                     `protobuf:"bytes,4,opt,name=TransactionName,proto3" json:"TransactionName,omitempty"`
	Timeout         int32                      `protobuf:"varint,5,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	BeginTime       int64                      `protobuf:"varint,6,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
	Status          GlobalSession_GlobalStatus `protobuf:"varint,7,opt,name=Status,proto3,enum=api.GlobalSession_GlobalStatus" json:"Status,omitempty"`
}

func (m *GlobalSession) Reset()      { *m = GlobalSession{} }
func (*GlobalSession) ProtoMessage() {}
func (*GlobalSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *GlobalSession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GlobalSession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GlobalSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSession.Merge(m, src)
}
func (m *GlobalSession) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSession) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSession.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSession proto.InternalMessageInfo

func (m *GlobalSession) GetXID() string {
	if m != nil {
		return m.XID
	}
	return ""
}

func (m *GlobalSession) GetApplicationID() string {
	if m != nil {
		return m.ApplicationID
	}
	return ""
}

func (m *GlobalSession) GetTransactionID() int64 {
	if m != nil {
		return m.TransactionID
	}
	return 0
}

func (m *GlobalSession) GetTransactionName() string {
	if m != nil {
		return m.TransactionName
	}
	return ""
}

func (m *GlobalSession) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *GlobalSession) GetBeginTime() int64 {
	if m != nil {
		return m.BeginTime
	}
	return 0
}

func (m *GlobalSession) GetStatus() GlobalSession_GlobalStatus {
	if m != nil {
		return m.Status
	}
	return Begin
}

type BranchSession struct {
	BranchID        string                     `protobuf:"bytes,1,opt,name=BranchID,proto3" json:"BranchID,omitempty"`
	ApplicationID   string                     `protobuf:"bytes,2,opt,name=ApplicationID,proto3" json:"ApplicationID,omitempty"`
	BranchSessionID int64                      `protobuf:"varint,3,opt,name=BranchSessionID,proto3" json:"BranchSessionID,omitempty"`
	XID             string                     `protobuf:"bytes,4,opt,name=XID,proto3" json:"XID,omitempty"`
	TransactionID   int64                      `protobuf:"varint,5,opt,name=TransactionID,proto3" json:"TransactionID,omitempty"`
	ResourceID      string                     `protobuf:"bytes,6,opt,name=ResourceID,proto3" json:"ResourceID,omitempty"`
	LockKey         string                     `protobuf:"bytes,7,opt,name=LockKey,proto3" json:"LockKey,omitempty"`
	Type            BranchSession_BranchType   `protobuf:"varint,8,opt,name=Type,proto3,enum=api.BranchSession_BranchType" json:"Type,omitempty"`
	Status          BranchSession_BranchStatus `protobuf:"varint,9,opt,name=Status,proto3,enum=api.BranchSession_BranchStatus" json:"Status,omitempty"`
	ApplicationData []byte                     `protobuf:"bytes,10,opt,name=ApplicationData,proto3" json:"ApplicationData,omitempty"`
	BeginTime       int64                      `protobuf:"varint,11,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
}

func (m *BranchSession) Reset()      { *m = BranchSession{} }
func (*BranchSession) ProtoMessage() {}
func (*BranchSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *BranchSession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BranchSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BranchSession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BranchSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchSession.Merge(m, src)
}
func (m *BranchSession) XXX_Size() int {
	return m.Size()
}
func (m *BranchSession) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchSession.DiscardUnknown(m)
}

var xxx_messageInfo_BranchSession proto.InternalMessageInfo

func (m *BranchSession) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

func (m *BranchSession) GetApplicationID() string {
	if m != nil {
		return m.ApplicationID
	}
	return ""
}

func (m *BranchSession) GetBranchSessionID() int64 {
	if m != nil {
		return m.BranchSessionID
	}
	return 0
}

func (m *BranchSession) GetXID() string {
	if m != nil {
		return m.XID
	}
	return ""
}

func (m *BranchSession) GetTransactionID() int64 {
	if m != nil {
		return m.TransactionID
	}
	return 0
}

func (m *BranchSession) GetResourceID() string {
	if m != nil {
		return m.ResourceID
	}
	return ""
}

func (m *BranchSession) GetLockKey() string {
	if m != nil {
		return m.LockKey
	}
	return ""
}

func (m *BranchSession) GetType() BranchSession_BranchType {
	if m != nil {
		return m.Type
	}
	return AT
}

func (m *BranchSession) GetStatus() BranchSession_BranchStatus {
	if m != nil {
		return m.Status
	}
	return Registered
}

func (m *BranchSession) GetApplicationData() []byte {
	if m != nil {
		return m.ApplicationData
	}
	return nil
}

func (m *BranchSession) GetBeginTime() int64 {
	if m != nil {
		return m.BeginTime
	}
	return 0
}

func init() {
	proto.RegisterEnum("api.ResultCode", ResultCode_name, ResultCode_value)
	proto.RegisterEnum("api.GlobalSession_GlobalStatus", GlobalSession_GlobalStatus_name, GlobalSession_GlobalStatus_value)
	proto.RegisterEnum("api.BranchSession_BranchType", BranchSession_BranchType_name, BranchSession_BranchType_value)
	proto.RegisterEnum("api.BranchSession_BranchStatus", BranchSession_BranchStatus_name, BranchSession_BranchStatus_value)
	proto.RegisterType((*GlobalSession)(nil), "api.GlobalSession")
	proto.RegisterType((*BranchSession)(nil), "api.BranchSession")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xbb, 0x8e, 0xd3, 0x40,
	0x18, 0x85, 0x3d, 0xb6, 0x93, 0x8d, 0xff, 0xcd, 0x6e, 0x86, 0xe1, 0x66, 0x21, 0x18, 0xa2, 0x88,
	0xc2, 0xa2, 0x08, 0x02, 0x0a, 0x84, 0xb6, 0xca, 0x45, 0xbb, 0x44, 0x20, 0x40, 0x8e, 0x8b, 0x15,
	0xdd, 0xc4, 0x19, 0x25, 0xa3, 0x75, 0x3c, 0x26, 0x76, 0x84, 0xb6, 0xe3, 0x11, 0x78, 0x0c, 0x5e,
	0x04, 0x89, 0x32, 0xa2, 0xda, 0x92, 0x38, 0x0d, 0xe5, 0x3e, 0x02, 0xf2, 0x10, 0x27, 0x76, 0xb4,
	0x05, 0x5d, 0xce, 0x37, 0x7f, 0xe6, 0x72, 0xce, 0x91, 0xc1, 0x62, 0x91, 0x68, 0x47, 0x73, 0x99,
	0x48, 0x62, 0xb0, 0x48, 0xb4, 0x7e, 0xe9, 0x70, 0x74, 0x16, 0xc8, 0x11, 0x0b, 0x86, 0x3c, 0x8e,
	0x85, 0x0c, 0x09, 0x06, 0xe3, 0x7c, 0xd0, 0xb7, 0x51, 0x13, 0x39, 0x96, 0x9b, 0xfd, 0x24, 0x4f,
	0xe0, 0xa8, 0x13, 0x45, 0x81, 0xf0, 0x59, 0x22, 0x64, 0x38, 0xe8, 0xdb, 0xba, 0x5a, 0x2b, 0xc3,
	0x6c, 0xca, 0x9b, 0xb3, 0x30, 0x66, 0xfe, 0x66, 0xca, 0x68, 0x22, 0xc7, 0x70, 0xcb, 0x90, 0x38,
	0xd0, 0x28, 0x80, 0xf7, 0x6c, 0xc6, 0x6d, 0x53, 0xed, 0xb6, 0x8f, 0x89, 0x0d, 0x07, 0x9e, 0x98,
	0x71, 0xb9, 0x48, 0xec, 0x4a, 0x13, 0x39, 0x15, 0x37, 0x97, 0xe4, 0x21, 0x58, 0x5d, 0x3e, 0x11,
	0x61, 0xa6, 0xed, 0xaa, 0x3a, 0x65, 0x07, 0xc8, 0x2b, 0xa8, 0x0e, 0x13, 0x96, 0x2c, 0x62, 0xfb,
	0xa0, 0x89, 0x9c, 0xe3, 0x17, 0x8f, 0xdb, 0xd9, 0x93, 0x4b, 0x6f, 0xcc, 0x95, 0x1a, 0x73, 0x37,
	0xe3, 0xad, 0x37, 0x50, 0x2f, 0x72, 0x62, 0x41, 0x45, 0xed, 0x8a, 0x35, 0x72, 0x0c, 0xd0, 0x93,
	0xb3, 0x99, 0x48, 0x12, 0x11, 0x4e, 0x30, 0x22, 0x0d, 0x38, 0x74, 0x65, 0x10, 0x8c, 0x98, 0x7f,
	0x91, 0x01, 0x9d, 0xd4, 0xa1, 0x76, 0x2a, 0x42, 0x11, 0x4f, 0xf9, 0x18, 0x1b, 0xad, 0x1f, 0x26,
	0x1c, 0x75, 0xe7, 0x2c, 0xf4, 0xa7, 0xb9, 0xa9, 0x0f, 0xa0, 0xf6, 0x0f, 0x6c, 0x9d, 0xdd, 0xea,
	0xff, 0xb4, 0xd7, 0x81, 0x46, 0x69, 0xcb, 0xad, 0xc1, 0xfb, 0x38, 0x0f, 0xd0, 0x2c, 0x05, 0x58,
	0x8e, 0xa6, 0x72, 0x53, 0x34, 0x14, 0xc0, 0xe5, 0xb1, 0x5c, 0xcc, 0x7d, 0x3e, 0xe8, 0x2b, 0x5f,
	0x2d, 0xb7, 0x40, 0xb2, 0x40, 0xde, 0x49, 0xff, 0xe2, 0x2d, 0xbf, 0x54, 0xce, 0x5a, 0x6e, 0x2e,
	0xc9, 0x73, 0x30, 0xbd, 0xcb, 0x88, 0xdb, 0x35, 0x65, 0xf8, 0x23, 0x65, 0x78, 0xe9, 0x56, 0x1b,
	0x95, 0x0d, 0xb9, 0x6a, 0xb4, 0x90, 0x92, 0x55, 0x48, 0xe9, 0xa6, 0x3f, 0x95, 0x53, 0xca, 0x7c,
	0x28, 0x18, 0xd3, 0x67, 0x09, 0xb3, 0xa1, 0x89, 0x9c, 0xba, 0xbb, 0x8f, 0xcb, 0x35, 0x39, 0xdc,
	0xab, 0x49, 0xeb, 0x19, 0xc0, 0xee, 0x52, 0xa4, 0x0a, 0x7a, 0xc7, 0xc3, 0x1a, 0x39, 0x00, 0xc3,
	0xeb, 0xf5, 0x30, 0x22, 0x35, 0x30, 0x87, 0x9d, 0xb3, 0x0e, 0xd6, 0xb3, 0xa5, 0xf3, 0x0e, 0x36,
	0x5a, 0x9f, 0xa1, 0x5e, 0xbc, 0x50, 0xd6, 0x09, 0x97, 0x4f, 0x44, 0x9c, 0xf0, 0x39, 0x1f, 0x63,
	0x8d, 0x10, 0x38, 0xfe, 0x38, 0x65, 0x31, 0xff, 0x10, 0xf2, 0x53, 0x26, 0x02, 0x3e, 0xc6, 0x88,
	0xdc, 0x03, 0xa2, 0x98, 0xf7, 0x45, 0x16, 0xfa, 0xa3, 0x93, 0xfb, 0x70, 0x3b, 0xe7, 0xc5, 0x1e,
	0x19, 0x59, 0x8f, 0x7a, 0x72, 0x16, 0x05, 0x3c, 0xe1, 0xd8, 0x7c, 0xfa, 0x5a, 0x25, 0xb2, 0x08,
	0x92, 0x9e, 0x1c, 0x73, 0x72, 0x07, 0xf0, 0x4e, 0x6d, 0x8e, 0xd0, 0xc8, 0x5d, 0xb8, 0xb5, 0xa3,
	0xc3, 0x85, 0xef, 0xf3, 0x38, 0xc6, 0xa8, 0x7b, 0xb2, 0x5c, 0x51, 0xed, 0x6a, 0x45, 0xb5, 0xeb,
	0x15, 0x45, 0x5f, 0x53, 0x8a, 0xbe, 0xa7, 0x14, 0xfd, 0x4c, 0x29, 0x5a, 0xa6, 0x14, 0xfd, 0x4e,
	0x29, 0xfa, 0x93, 0x52, 0xed, 0x3a, 0xa5, 0xe8, 0xdb, 0x9a, 0x6a, 0xcb, 0x35, 0xd5, 0xae, 0xd6,
	0x54, 0xfb, 0x54, 0x69, 0x9f, 0xb0, 0x48, 0x8c, 0xaa, 0xea, 0x03, 0xf1, 0xf2, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x09, 0xa8, 0xa2, 0xa7, 0x2d, 0x04, 0x00, 0x00,
}

func (x ResultCode) String() string {
	s, ok := ResultCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x GlobalSession_GlobalStatus) String() string {
	s, ok := GlobalSession_GlobalStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x BranchSession_BranchType) String() string {
	s, ok := BranchSession_BranchType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x BranchSession_BranchStatus) String() string {
	s, ok := BranchSession_BranchStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *GlobalSession) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSession)
	if !ok {
		that2, ok := that.(GlobalSession)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.XID != that1.XID {
		return false
	}
	if this.ApplicationID != that1.ApplicationID {
		return false
	}
	if this.TransactionID != that1.TransactionID {
		return false
	}
	if this.TransactionName != that1.TransactionName {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if this.BeginTime != that1.BeginTime {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *BranchSession) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BranchSession)
	if !ok {
		that2, ok := that.(BranchSession)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BranchID != that1.BranchID {
		return false
	}
	if this.ApplicationID != that1.ApplicationID {
		return false
	}
	if this.BranchSessionID != that1.BranchSessionID {
		return false
	}
	if this.XID != that1.XID {
		return false
	}
	if this.TransactionID != that1.TransactionID {
		return false
	}
	if this.ResourceID != that1.ResourceID {
		return false
	}
	if this.LockKey != that1.LockKey {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !bytes.Equal(this.ApplicationData, that1.ApplicationData) {
		return false
	}
	if this.BeginTime != that1.BeginTime {
		return false
	}
	return true
}
func (this *GlobalSession) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&api.GlobalSession{")
	s = append(s, "XID: "+fmt.Sprintf("%#v", this.XID)+",\n")
	s = append(s, "ApplicationID: "+fmt.Sprintf("%#v", this.ApplicationID)+",\n")
	s = append(s, "TransactionID: "+fmt.Sprintf("%#v", this.TransactionID)+",\n")
	s = append(s, "TransactionName: "+fmt.Sprintf("%#v", this.TransactionName)+",\n")
	s = append(s, "Timeout: "+fmt.Sprintf("%#v", this.Timeout)+",\n")
	s = append(s, "BeginTime: "+fmt.Sprintf("%#v", this.BeginTime)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BranchSession) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
	s = append(s, "&api.BranchSession{")
	s = append(s, "BranchID: "+fmt.Sprintf("%#v", this.BranchID)+",\n")
	s = append(s, "ApplicationID: "+fmt.Sprintf("%#v", this.ApplicationID)+",\n")
	s = append(s, "BranchSessionID: "+fmt.Sprintf("%#v", this.BranchSessionID)+",\n")
	s = append(s, "XID: "+fmt.Sprintf("%#v", this.XID)+",\n")
	s = append(s, "TransactionID: "+fmt.Sprintf("%#v", this.TransactionID)+",\n")
	s = append(s, "ResourceID: "+fmt.Sprintf("%#v", this.ResourceID)+",\n")
	s = append(s, "LockKey: "+fmt.Sprintf("%#v", this.LockKey)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "ApplicationData: "+fmt.Sprintf("%#v", this.ApplicationData)+",\n")
	s = append(s, "BeginTime: "+fmt.Sprintf("%#v", this.BeginTime)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringApi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if m.BeginTime != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.BeginTime))
		i--
		dAtA[i] = 0x30
	}
	if m.Timeout != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TransactionName) > 0 {
		i -= len(m.TransactionName)
		copy(dAtA[i:], m.TransactionName)
		i = encodeVarintApi(dAtA, i, uint64(len(m.TransactionName)))
		i--
		dAtA[i] = 0x22
	}
	if m.TransactionID != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.TransactionID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ApplicationID) > 0 {
		i -= len(m.ApplicationID)
		copy(dAtA[i:], m.ApplicationID)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ApplicationID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XID) > 0 {
		i -= len(m.XID)
		copy(dAtA[i:], m.XID)
		i = encodeVarintApi(dAtA, i, uint64(len(m.XID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BranchSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BranchSession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BranchSession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BeginTime != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.BeginTime))
		i--
		dAtA[i] = 0x58
	}
	if len(m.ApplicationData) > 0 {
		i -= len(m.ApplicationData)
		copy(dAtA[i:], m.ApplicationData)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ApplicationData)))
		i--
		dAtA[i] = 0x52
	}
	if m.Status != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x48
	}
	if m.Type != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x40
	}
	if len(m.LockKey) > 0 {
		i -= len(m.LockKey)
		copy(dAtA[i:], m.LockKey)
		i = encodeVarintApi(dAtA, i, uint64(len(m.LockKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ResourceID) > 0 {
		i -= len(m.ResourceID)
		copy(dAtA[i:], m.ResourceID)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ResourceID)))
		i--
		dAtA[i] = 0x32
	}
	if m.TransactionID != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.TransactionID))
		i--
		dAtA[i] = 0x28
	}
	if len(m.XID) > 0 {
		i -= len(m.XID)
		copy(dAtA[i:], m.XID)
		i = encodeVarintApi(dAtA, i, uint64(len(m.XID)))
		i--
		dAtA[i] = 0x22
	}
	if m.BranchSessionID != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.BranchSessionID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ApplicationID) > 0 {
		i -= len(m.ApplicationID)
		copy(dAtA[i:], m.ApplicationID)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ApplicationID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BranchID) > 0 {
		i -= len(m.BranchID)
		copy(dAtA[i:], m.BranchID)
		i = encodeVarintApi(dAtA, i, uint64(len(m.BranchID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GlobalSession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ApplicationID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.TransactionID != 0 {
		n += 1 + sovApi(uint64(m.TransactionID))
	}
	l = len(m.TransactionName)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovApi(uint64(m.Timeout))
	}
	if m.BeginTime != 0 {
		n += 1 + sovApi(uint64(m.BeginTime))
	}
	if m.Status != 0 {
		n += 1 + sovApi(uint64(m.Status))
	}
	return n
}

func (m *BranchSession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BranchID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ApplicationID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.BranchSessionID != 0 {
		n += 1 + sovApi(uint64(m.BranchSessionID))
	}
	l = len(m.XID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.TransactionID != 0 {
		n += 1 + sovApi(uint64(m.TransactionID))
	}
	l = len(m.ResourceID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.LockKey)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovApi(uint64(m.Type))
	}
	if m.Status != 0 {
		n += 1 + sovApi(uint64(m.Status))
	}
	l = len(m.ApplicationData)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.BeginTime != 0 {
		n += 1 + sovApi(uint64(m.BeginTime))
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSession) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSession{`,
		`XID:` + fmt.Sprintf("%v", this.XID) + `,`,
		`ApplicationID:` + fmt.Sprintf("%v", this.ApplicationID) + `,`,
		`TransactionID:` + fmt.Sprintf("%v", this.TransactionID) + `,`,
		`TransactionName:` + fmt.Sprintf("%v", this.TransactionName) + `,`,
		`Timeout:` + fmt.Sprintf("%v", this.Timeout) + `,`,
		`BeginTime:` + fmt.Sprintf("%v", this.BeginTime) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BranchSession) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BranchSession{`,
		`BranchID:` + fmt.Sprintf("%v", this.BranchID) + `,`,
		`ApplicationID:` + fmt.Sprintf("%v", this.ApplicationID) + `,`,
		`BranchSessionID:` + fmt.Sprintf("%v", this.BranchSessionID) + `,`,
		`XID:` + fmt.Sprintf("%v", this.XID) + `,`,
		`TransactionID:` + fmt.Sprintf("%v", this.TransactionID) + `,`,
		`ResourceID:` + fmt.Sprintf("%v", this.ResourceID) + `,`,
		`LockKey:` + fmt.Sprintf("%v", this.LockKey) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`ApplicationData:` + fmt.Sprintf("%v", this.ApplicationData) + `,`,
		`BeginTime:` + fmt.Sprintf("%v", this.BeginTime) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GlobalSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionID", wireType)
			}
			m.TransactionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginTime", wireType)
			}
			m.BeginTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= GlobalSession_GlobalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BranchSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BranchSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BranchSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BranchID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BranchID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BranchSessionID", wireType)
			}
			m.BranchSessionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BranchSessionID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionID", wireType)
			}
			m.TransactionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= BranchSession_BranchType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= BranchSession_BranchStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationData = append(m.ApplicationData[:0], dAtA[iNdEx:postIndex]...)
			if m.ApplicationData == nil {
				m.ApplicationData = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginTime", wireType)
			}
			m.BeginTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
