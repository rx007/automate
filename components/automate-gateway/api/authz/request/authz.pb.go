// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/authz/request/authz.proto

package request // import "github.com/chef/automate/components/automate-gateway/api/authz/request"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IsAuthorizedReq struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedReq) Reset()         { *m = IsAuthorizedReq{} }
func (m *IsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReq) ProtoMessage()    {}
func (*IsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{0}
}
func (m *IsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReq.Unmarshal(m, b)
}
func (m *IsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReq.Marshal(b, m, deterministic)
}
func (dst *IsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReq.Merge(dst, src)
}
func (m *IsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReq.Size(m)
}
func (m *IsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReq proto.InternalMessageInfo

func (m *IsAuthorizedReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *IsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *IsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type CreatePolicyReq struct {
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Subjects             []string `protobuf:"bytes,3,rep,name=subjects,proto3" json:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePolicyReq) Reset()         { *m = CreatePolicyReq{} }
func (m *CreatePolicyReq) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyReq) ProtoMessage()    {}
func (*CreatePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{1}
}
func (m *CreatePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyReq.Unmarshal(m, b)
}
func (m *CreatePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyReq.Marshal(b, m, deterministic)
}
func (dst *CreatePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyReq.Merge(dst, src)
}
func (m *CreatePolicyReq) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyReq.Size(m)
}
func (m *CreatePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyReq proto.InternalMessageInfo

func (m *CreatePolicyReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreatePolicyReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *CreatePolicyReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type ListPoliciesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPoliciesReq) Reset()         { *m = ListPoliciesReq{} }
func (m *ListPoliciesReq) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesReq) ProtoMessage()    {}
func (*ListPoliciesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{2}
}
func (m *ListPoliciesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesReq.Unmarshal(m, b)
}
func (m *ListPoliciesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesReq.Marshal(b, m, deterministic)
}
func (dst *ListPoliciesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesReq.Merge(dst, src)
}
func (m *ListPoliciesReq) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesReq.Size(m)
}
func (m *ListPoliciesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesReq proto.InternalMessageInfo

type DeletePolicyReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyReq) Reset()         { *m = DeletePolicyReq{} }
func (m *DeletePolicyReq) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyReq) ProtoMessage()    {}
func (*DeletePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{3}
}
func (m *DeletePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyReq.Unmarshal(m, b)
}
func (m *DeletePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyReq.Marshal(b, m, deterministic)
}
func (dst *DeletePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyReq.Merge(dst, src)
}
func (m *DeletePolicyReq) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyReq.Size(m)
}
func (m *DeletePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyReq proto.InternalMessageInfo

func (m *DeletePolicyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type IntrospectAllReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectAllReq) Reset()         { *m = IntrospectAllReq{} }
func (m *IntrospectAllReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectAllReq) ProtoMessage()    {}
func (*IntrospectAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{4}
}
func (m *IntrospectAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectAllReq.Unmarshal(m, b)
}
func (m *IntrospectAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectAllReq.Marshal(b, m, deterministic)
}
func (dst *IntrospectAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectAllReq.Merge(dst, src)
}
func (m *IntrospectAllReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectAllReq.Size(m)
}
func (m *IntrospectAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectAllReq proto.InternalMessageInfo

type IntrospectSomeReq struct {
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectSomeReq) Reset()         { *m = IntrospectSomeReq{} }
func (m *IntrospectSomeReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectSomeReq) ProtoMessage()    {}
func (*IntrospectSomeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{5}
}
func (m *IntrospectSomeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectSomeReq.Unmarshal(m, b)
}
func (m *IntrospectSomeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectSomeReq.Marshal(b, m, deterministic)
}
func (dst *IntrospectSomeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectSomeReq.Merge(dst, src)
}
func (m *IntrospectSomeReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectSomeReq.Size(m)
}
func (m *IntrospectSomeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectSomeReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectSomeReq proto.InternalMessageInfo

func (m *IntrospectSomeReq) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type IntrospectReq struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Parameters           []string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectReq) Reset()         { *m = IntrospectReq{} }
func (m *IntrospectReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectReq) ProtoMessage()    {}
func (*IntrospectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{6}
}
func (m *IntrospectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectReq.Unmarshal(m, b)
}
func (m *IntrospectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectReq.Marshal(b, m, deterministic)
}
func (dst *IntrospectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectReq.Merge(dst, src)
}
func (m *IntrospectReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectReq.Size(m)
}
func (m *IntrospectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectReq proto.InternalMessageInfo

func (m *IntrospectReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IntrospectReq) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type IntrospectAllProjectsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectAllProjectsReq) Reset()         { *m = IntrospectAllProjectsReq{} }
func (m *IntrospectAllProjectsReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectAllProjectsReq) ProtoMessage()    {}
func (*IntrospectAllProjectsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_c6fc09cba565b7a1, []int{7}
}
func (m *IntrospectAllProjectsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectAllProjectsReq.Unmarshal(m, b)
}
func (m *IntrospectAllProjectsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectAllProjectsReq.Marshal(b, m, deterministic)
}
func (dst *IntrospectAllProjectsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectAllProjectsReq.Merge(dst, src)
}
func (m *IntrospectAllProjectsReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectAllProjectsReq.Size(m)
}
func (m *IntrospectAllProjectsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectAllProjectsReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectAllProjectsReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IsAuthorizedReq)(nil), "chef.automate.api.authz.request.IsAuthorizedReq")
	proto.RegisterType((*CreatePolicyReq)(nil), "chef.automate.api.authz.request.CreatePolicyReq")
	proto.RegisterType((*ListPoliciesReq)(nil), "chef.automate.api.authz.request.ListPoliciesReq")
	proto.RegisterType((*DeletePolicyReq)(nil), "chef.automate.api.authz.request.DeletePolicyReq")
	proto.RegisterType((*IntrospectAllReq)(nil), "chef.automate.api.authz.request.IntrospectAllReq")
	proto.RegisterType((*IntrospectSomeReq)(nil), "chef.automate.api.authz.request.IntrospectSomeReq")
	proto.RegisterType((*IntrospectReq)(nil), "chef.automate.api.authz.request.IntrospectReq")
	proto.RegisterType((*IntrospectAllProjectsReq)(nil), "chef.automate.api.authz.request.IntrospectAllProjectsReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/authz/request/authz.proto", fileDescriptor_authz_c6fc09cba565b7a1)
}

var fileDescriptor_authz_c6fc09cba565b7a1 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbf, 0x4f, 0xf3, 0x30,
	0x10, 0x55, 0xd3, 0x7e, 0xfd, 0xbe, 0xef, 0x24, 0x08, 0x8d, 0x10, 0xb2, 0x3a, 0x40, 0xc9, 0x04,
	0x03, 0xc9, 0xc0, 0xc4, 0x58, 0x8a, 0x10, 0x95, 0x18, 0xaa, 0xb0, 0xb1, 0x20, 0xd7, 0x3d, 0x1a,
	0xa3, 0x24, 0x76, 0xec, 0x8b, 0x50, 0xfb, 0xd7, 0x23, 0xbb, 0xe9, 0xaf, 0x91, 0xcd, 0xef, 0xde,
	0x7b, 0x7e, 0x77, 0xba, 0x83, 0x07, 0xa1, 0x4a, 0xad, 0x2a, 0xac, 0xc8, 0xa6, 0xbc, 0x21, 0x55,
	0x72, 0xc2, 0xbb, 0x25, 0x27, 0xfc, 0xe6, 0xab, 0x94, 0x6b, 0xe9, 0x8a, 0xf9, 0x3a, 0x35, 0x58,
	0x37, 0x68, 0x69, 0x83, 0x12, 0x6d, 0x14, 0xa9, 0xe8, 0x4a, 0xe4, 0xf8, 0x99, 0x6c, 0x4d, 0x09,
	0xd7, 0x32, 0xd9, 0xd0, 0xad, 0x38, 0xfe, 0x80, 0x70, 0x6a, 0xc7, 0x0d, 0xe5, 0xca, 0xc8, 0x35,
	0x2e, 0x32, 0xac, 0x23, 0x06, 0x7f, 0x6d, 0x33, 0xff, 0x42, 0x41, 0xac, 0x33, 0xea, 0xdc, 0xfc,
	0xcf, 0xb6, 0x30, 0x1a, 0xc2, 0x3f, 0x83, 0x56, 0x35, 0x46, 0x20, 0x0b, 0x3c, 0xb5, 0xc3, 0xd1,
	0x05, 0xf4, 0xb9, 0x20, 0xa9, 0x2a, 0xd6, 0xf5, 0x4c, 0x8b, 0x62, 0x0e, 0xe1, 0xc4, 0x20, 0x27,
	0x9c, 0xa9, 0x42, 0x8a, 0x95, 0x0b, 0xd8, 0x4b, 0x83, 0x43, 0xa9, 0xfb, 0xbe, 0x4d, 0xb2, 0xac,
	0x3b, 0xea, 0xba, 0xef, 0xb7, 0xf8, 0x28, 0xba, 0x77, 0x1c, 0x1d, 0x0f, 0x20, 0x7c, 0x95, 0x96,
	0x7c, 0x80, 0x44, 0x9b, 0x61, 0x1d, 0x5f, 0x43, 0xf8, 0x84, 0x05, 0x1e, 0xa6, 0x9e, 0x42, 0x20,
	0x17, 0xed, 0x44, 0x81, 0x5c, 0xc4, 0x11, 0x9c, 0x4d, 0x2b, 0x32, 0xca, 0x6a, 0x14, 0x34, 0x2e,
	0x0a, 0x67, 0xbb, 0x85, 0xc1, 0xbe, 0xf6, 0xa6, 0x4a, 0x74, 0xc6, 0x73, 0xf8, 0xa3, 0x39, 0xe5,
	0x96, 0x75, 0x7c, 0x4f, 0x1b, 0x10, 0x4f, 0xe0, 0x64, 0x2f, 0x75, 0xb2, 0x08, 0x7a, 0x8e, 0x69,
	0x13, 0xfc, 0x3b, 0xba, 0x04, 0xd0, 0xdc, 0xf0, 0x12, 0x09, 0x8d, 0x65, 0x81, 0xf7, 0x1f, 0x54,
	0xe2, 0x21, 0xb0, 0xa3, 0x1e, 0x66, 0x46, 0xf9, 0x71, 0x33, 0xac, 0x1f, 0x5f, 0xde, 0x9f, 0x97,
	0x92, 0xf2, 0x66, 0x9e, 0x08, 0x55, 0xa6, 0x6e, 0x8f, 0xbb, 0xe5, 0xa7, 0xbf, 0x3a, 0x88, 0x79,
	0xdf, 0xdf, 0xc2, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xb9, 0x99, 0x13, 0x48, 0x02,
	0x00, 0x00,
}
