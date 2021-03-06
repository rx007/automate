// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/common/policy.proto

package common // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"

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

type Type int32

const (
	Type_CHEF_MANAGED Type = 0
	Type_CUSTOM       Type = 1
)

var Type_name = map[int32]string{
	0: "CHEF_MANAGED",
	1: "CUSTOM",
}
var Type_value = map[string]int32{
	"CHEF_MANAGED": 0,
	"CUSTOM":       1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{0}
}

// passed to UpgradeToV2 to set version
type Flag int32

const (
	Flag_VERSION_2_0 Flag = 0
	Flag_VERSION_2_1 Flag = 1
)

var Flag_name = map[int32]string{
	0: "VERSION_2_0",
	1: "VERSION_2_1",
}
var Flag_value = map[string]int32{
	"VERSION_2_0": 0,
	"VERSION_2_1": 1,
}

func (x Flag) String() string {
	return proto.EnumName(Flag_name, int32(x))
}
func (Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{1}
}

type Statement_Effect int32

const (
	Statement_ALLOW Statement_Effect = 0
	Statement_DENY  Statement_Effect = 1
)

var Statement_Effect_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}
var Statement_Effect_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x Statement_Effect) String() string {
	return proto.EnumName(Statement_Effect_name, int32(x))
}
func (Statement_Effect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{1, 0}
}

type Version_VersionNumber int32

const (
	Version_V0 Version_VersionNumber = 0
	Version_V1 Version_VersionNumber = 1
	Version_V2 Version_VersionNumber = 2
)

var Version_VersionNumber_name = map[int32]string{
	0: "V0",
	1: "V1",
	2: "V2",
}
var Version_VersionNumber_value = map[string]int32{
	"V0": 0,
	"V1": 1,
	"V2": 2,
}

func (x Version_VersionNumber) String() string {
	return proto.EnumName(Version_VersionNumber_name, int32(x))
}
func (Version_VersionNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{4, 0}
}

type Policy struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Type         `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2beta.Type" json:"type,omitempty"`
	Members              []string     `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	Statements           []*Statement `protobuf:"bytes,5,rep,name=statements,proto3" json:"statements,omitempty"`
	Projects             []string     `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Policy) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Policy) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Policy) GetStatements() []*Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *Policy) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Statement struct {
	Effect Statement_Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=chef.automate.api.iam.v2beta.Statement_Effect" json:"effect,omitempty"`
	// inline definitions
	Actions []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	// references
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// Note: these are for display only, not to be set in CreatePolicy/UpdatePolicy
	Resources            []string `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{1}
}
func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (dst *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(dst, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetEffect() Statement_Effect {
	if m != nil {
		return m.Effect
	}
	return Statement_ALLOW
}

func (m *Statement) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Statement) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Statement) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Statement) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Role struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Type     `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2beta.Type" json:"type,omitempty"`
	Actions              []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{2}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Role) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Role) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Project struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Type     `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2beta.Type" json:"type,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{3}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (dst *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(dst, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Project) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

// the only values that may be returned by GetPolicyVersion
type Version struct {
	Major                Version_VersionNumber `protobuf:"varint,1,opt,name=major,proto3,enum=chef.automate.api.iam.v2beta.Version_VersionNumber" json:"major,omitempty"`
	Minor                Version_VersionNumber `protobuf:"varint,2,opt,name=minor,proto3,enum=chef.automate.api.iam.v2beta.Version_VersionNumber" json:"minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_3901461ee0db5b96, []int{4}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (dst *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(dst, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() Version_VersionNumber {
	if m != nil {
		return m.Major
	}
	return Version_V0
}

func (m *Version) GetMinor() Version_VersionNumber {
	if m != nil {
		return m.Minor
	}
	return Version_V0
}

func init() {
	proto.RegisterType((*Policy)(nil), "chef.automate.api.iam.v2beta.Policy")
	proto.RegisterType((*Statement)(nil), "chef.automate.api.iam.v2beta.Statement")
	proto.RegisterType((*Role)(nil), "chef.automate.api.iam.v2beta.Role")
	proto.RegisterType((*Project)(nil), "chef.automate.api.iam.v2beta.Project")
	proto.RegisterType((*Version)(nil), "chef.automate.api.iam.v2beta.Version")
	proto.RegisterEnum("chef.automate.api.iam.v2beta.Type", Type_name, Type_value)
	proto.RegisterEnum("chef.automate.api.iam.v2beta.Flag", Flag_name, Flag_value)
	proto.RegisterEnum("chef.automate.api.iam.v2beta.Statement_Effect", Statement_Effect_name, Statement_Effect_value)
	proto.RegisterEnum("chef.automate.api.iam.v2beta.Version_VersionNumber", Version_VersionNumber_name, Version_VersionNumber_value)
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/common/policy.proto", fileDescriptor_policy_3901461ee0db5b96)
}

var fileDescriptor_policy_3901461ee0db5b96 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0xc5, 0x10, 0x42, 0x73, 0xbb, 0xb1, 0xc8, 0x4f, 0xd1, 0xd4, 0x49, 0x08, 0x4d, 0x2a, 0xaa,
	0x34, 0x67, 0x4d, 0xa5, 0x3d, 0x4e, 0x62, 0x2d, 0x74, 0x9d, 0x28, 0x54, 0xa1, 0x63, 0xda, 0x5e,
	0x90, 0x49, 0x0d, 0x75, 0x85, 0xe3, 0x28, 0x31, 0x9b, 0x78, 0xdc, 0x5f, 0xec, 0x93, 0xf6, 0x0d,
	0x7b, 0xdd, 0x8f, 0x4c, 0x76, 0x08, 0x83, 0x3e, 0xa0, 0xad, 0x52, 0x9f, 0xae, 0xef, 0xd5, 0x3d,
	0xe7, 0x9e, 0x63, 0x5b, 0x17, 0xde, 0x46, 0x52, 0x24, 0x32, 0x66, 0xb1, 0xca, 0x7c, 0xba, 0x50,
	0x52, 0x50, 0xc5, 0x5e, 0xcd, 0xa8, 0x62, 0xdf, 0xe8, 0xd2, 0xa7, 0x09, 0xf7, 0x39, 0x15, 0xfe,
	0xd7, 0x60, 0xc2, 0x14, 0xf5, 0x23, 0x29, 0x84, 0x8c, 0xfd, 0x44, 0xce, 0x79, 0xb4, 0x24, 0x49,
	0x2a, 0x95, 0xc4, 0x07, 0xd1, 0x2d, 0x9b, 0x92, 0x02, 0x49, 0x68, 0xc2, 0x09, 0xa7, 0x82, 0xe4,
	0x88, 0xe6, 0x6f, 0x04, 0xf6, 0x95, 0x69, 0xc7, 0x18, 0xac, 0x98, 0x0a, 0xe6, 0xa1, 0x06, 0x6a,
	0x39, 0xa1, 0x39, 0xe3, 0x3a, 0x94, 0xf9, 0x8d, 0x57, 0x36, 0x95, 0x32, 0xbf, 0xc1, 0x6f, 0xc0,
	0x52, 0xcb, 0x84, 0x79, 0x95, 0x06, 0x6a, 0xd5, 0x83, 0x26, 0xd9, 0xc5, 0x4d, 0xae, 0x97, 0x09,
	0x0b, 0x4d, 0x3f, 0xf6, 0xa0, 0x26, 0x98, 0x98, 0xb0, 0x34, 0xf3, 0xac, 0x46, 0xa5, 0xe5, 0x84,
	0x45, 0x8a, 0xcf, 0x01, 0x32, 0x45, 0x15, 0x13, 0xda, 0xa0, 0x57, 0x6d, 0x54, 0x5a, 0xfb, 0xc1,
	0xe1, 0x6e, 0xde, 0x61, 0xd1, 0x1f, 0x6e, 0x40, 0xf1, 0x73, 0xd8, 0x4b, 0x52, 0x79, 0xc7, 0x22,
	0x95, 0x79, 0xb6, 0x99, 0xb1, 0xce, 0x9b, 0xbf, 0x10, 0x38, 0x6b, 0x14, 0xee, 0x82, 0xcd, 0xa6,
	0x53, 0x16, 0x29, 0x63, 0xb5, 0x1e, 0x90, 0x7f, 0x1c, 0x47, 0x3a, 0x06, 0x15, 0xae, 0xd0, 0xda,
	0x14, 0x8d, 0x14, 0x97, 0x71, 0xe6, 0x55, 0x72, 0x53, 0xab, 0x54, 0x5f, 0x65, 0x2a, 0xe7, 0xcc,
	0xb3, 0xf2, 0xab, 0xd4, 0x67, 0x7c, 0x00, 0x4e, 0xca, 0x32, 0xb9, 0x48, 0x23, 0x96, 0xfb, 0x74,
	0xc2, 0xbf, 0x85, 0x9d, 0xea, 0x5f, 0x80, 0x9d, 0x4f, 0xc6, 0x0e, 0x54, 0xdb, 0xbd, 0xde, 0xe0,
	0x93, 0x5b, 0xc2, 0x7b, 0x60, 0x9d, 0x75, 0xfa, 0x9f, 0x5d, 0xd4, 0xfc, 0x81, 0xc0, 0x0a, 0xf5,
	0x84, 0x47, 0x7e, 0xc0, 0xc2, 0xab, 0xb5, 0xed, 0x75, 0x53, 0x79, 0xf5, 0x9e, 0xf2, 0xef, 0x08,
	0x6a, 0x57, 0x79, 0xf2, 0xa8, 0xea, 0x36, 0x35, 0x58, 0xf7, 0x34, 0xfc, 0x44, 0x50, 0x1b, 0xb1,
	0x34, 0xe3, 0x32, 0xc6, 0x17, 0x50, 0x15, 0xf4, 0x4e, 0xa6, 0xab, 0x87, 0x3f, 0xd9, 0x3d, 0x60,
	0x85, 0x2a, 0x62, 0x7f, 0xa1, 0x7f, 0x6c, 0x98, 0x33, 0x18, 0x2a, 0x1e, 0xcb, 0xd4, 0xa8, 0x7f,
	0x30, 0x95, 0x66, 0x68, 0x1e, 0xc2, 0xd3, 0xad, 0x3a, 0xb6, 0xa1, 0x3c, 0x7a, 0xed, 0x96, 0x4c,
	0x3c, 0x76, 0x91, 0x89, 0x81, 0x5b, 0x3e, 0x7a, 0x09, 0x96, 0x36, 0x8d, 0x5d, 0x78, 0x72, 0xfa,
	0xbe, 0xd3, 0x1d, 0x5f, 0xb6, 0xfb, 0xed, 0xf3, 0xce, 0x99, 0x5b, 0xc2, 0x00, 0xf6, 0xe9, 0xc7,
	0xe1, 0xf5, 0xe0, 0xd2, 0x45, 0x47, 0x2d, 0xb0, 0xba, 0x73, 0x3a, 0xc3, 0xcf, 0x60, 0x7f, 0xd4,
	0x09, 0x87, 0x17, 0x83, 0xfe, 0x38, 0x18, 0x6b, 0xba, 0xad, 0xc2, 0xb1, 0x8b, 0xde, 0xf5, 0xbe,
	0x7c, 0x98, 0x71, 0x75, 0xbb, 0x98, 0x90, 0x48, 0x0a, 0x5f, 0x1b, 0x58, 0x6f, 0x18, 0xff, 0xbf,
	0xb7, 0xce, 0xc4, 0x36, 0xfb, 0xe6, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x8d, 0x79,
	0xf8, 0xb1, 0x04, 0x00, 0x00,
}
