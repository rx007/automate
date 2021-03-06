// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/workflow_server/config_request.proto

package workflow_server // import "github.com/chef/automate/api/config/workflow_server"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shared "github.com/chef/automate/api/config/shared"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(dst, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0}
}
func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(dst, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                    *shared.Mlsa                 `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Log                     *ConfigRequest_V1_System_Log `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	DevEnv                  bool                         `protobuf:"varint,3,opt,name=dev_env,json=devEnv,proto3" json:"dev_env,omitempty" toml:"dev_env,omitempty" mapstructure:"dev_env,omitempty"`
	Telemetry               *ConfigRequest_V1_Telemetry  `protobuf:"bytes,4,opt,name=telemetry,proto3" json:"telemetry,omitempty" toml:"telemetry,omitempty" mapstructure:"telemetry,omitempty"`
	ChefServer              *ConfigRequest_V1_ChefServer `protobuf:"bytes,5,opt,name=chef_server,json=chefServer,proto3" json:"chef_server,omitempty" toml:"chef_server,omitempty" mapstructure:"chef_server,omitempty"`
	Postgresql              *ConfigRequest_V1_Postgresql `protobuf:"bytes,6,opt,name=postgresql,proto3" json:"postgresql,omitempty" toml:"postgresql,omitempty" mapstructure:"postgresql,omitempty"`
	SshGit                  *ConfigRequest_V1_SshGit     `protobuf:"bytes,7,opt,name=ssh_git,json=sshGit,proto3" json:"ssh_git,omitempty" toml:"ssh_git,omitempty" mapstructure:"ssh_git,omitempty"`
	Fqdn                    *wrappers.StringValue        `protobuf:"bytes,9,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	ApiPort                 *wrappers.Int32Value         `protobuf:"bytes,11,opt,name=api_port,json=apiPort,proto3" json:"api_port,omitempty" toml:"api_port,omitempty" mapstructure:"api_port,omitempty"`
	ApiProto                *wrappers.StringValue        `protobuf:"bytes,12,opt,name=api_proto,json=apiProto,proto3" json:"api_proto,omitempty" toml:"api_proto,omitempty" mapstructure:"api_proto,omitempty"`
	GitExecutable           *wrappers.StringValue        `protobuf:"bytes,13,opt,name=git_executable,json=gitExecutable,proto3" json:"git_executable,omitempty" toml:"git_executable,omitempty" mapstructure:"git_executable,omitempty"`
	TrustedCertificatesFile *wrappers.StringValue        `protobuf:"bytes,14,opt,name=trusted_certificates_file,json=trustedCertificatesFile,proto3" json:"trusted_certificates_file,omitempty" toml:"trusted_certificates_file,omitempty" mapstructure:"trusted_certificates_file,omitempty"`
	Tls                     *shared.TLSCredentials       `protobuf:"bytes,15,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized        []byte                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache           int32                        `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 0}
}
func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(dst, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetDevEnv() bool {
	if m != nil {
		return m.DevEnv
	}
	return false
}

func (m *ConfigRequest_V1_System) GetTelemetry() *ConfigRequest_V1_Telemetry {
	if m != nil {
		return m.Telemetry
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetChefServer() *ConfigRequest_V1_ChefServer {
	if m != nil {
		return m.ChefServer
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetPostgresql() *ConfigRequest_V1_Postgresql {
	if m != nil {
		return m.Postgresql
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetSshGit() *ConfigRequest_V1_SshGit {
	if m != nil {
		return m.SshGit
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetFqdn() *wrappers.StringValue {
	if m != nil {
		return m.Fqdn
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetApiPort() *wrappers.Int32Value {
	if m != nil {
		return m.ApiPort
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetApiProto() *wrappers.StringValue {
	if m != nil {
		return m.ApiProto
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetGitExecutable() *wrappers.StringValue {
	if m != nil {
		return m.GitExecutable
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTrustedCertificatesFile() *wrappers.StringValue {
	if m != nil {
		return m.TrustedCertificatesFile
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 0, 0}
}
func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 1}
}
func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(dst, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

type ConfigRequest_V1_ChefServer struct {
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty" toml:"url,omitempty" mapstructure:"url,omitempty"`
	WebUiUrl             *wrappers.StringValue `protobuf:"bytes,2,opt,name=web_ui_url,json=webUiUrl,proto3" json:"web_ui_url,omitempty" toml:"web_ui_url,omitempty" mapstructure:"web_ui_url,omitempty"`
	Vip                  *wrappers.StringValue `protobuf:"bytes,3,opt,name=vip,proto3" json:"vip,omitempty" toml:"vip,omitempty" mapstructure:"vip,omitempty"`
	ChefUser             *wrappers.StringValue `protobuf:"bytes,5,opt,name=chef_user,json=chefUser,proto3" json:"chef_user,omitempty" toml:"chef_user,omitempty" mapstructure:"chef_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_ChefServer) Reset()         { *m = ConfigRequest_V1_ChefServer{} }
func (m *ConfigRequest_V1_ChefServer) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_ChefServer) ProtoMessage()    {}
func (*ConfigRequest_V1_ChefServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 2}
}
func (m *ConfigRequest_V1_ChefServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_ChefServer.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_ChefServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_ChefServer.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_ChefServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_ChefServer.Merge(dst, src)
}
func (m *ConfigRequest_V1_ChefServer) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_ChefServer.Size(m)
}
func (m *ConfigRequest_V1_ChefServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_ChefServer.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_ChefServer proto.InternalMessageInfo

func (m *ConfigRequest_V1_ChefServer) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *ConfigRequest_V1_ChefServer) GetWebUiUrl() *wrappers.StringValue {
	if m != nil {
		return m.WebUiUrl
	}
	return nil
}

func (m *ConfigRequest_V1_ChefServer) GetVip() *wrappers.StringValue {
	if m != nil {
		return m.Vip
	}
	return nil
}

func (m *ConfigRequest_V1_ChefServer) GetChefUser() *wrappers.StringValue {
	if m != nil {
		return m.ChefUser
	}
	return nil
}

type ConfigRequest_V1_Postgresql struct {
	DbName               *wrappers.StringValue `protobuf:"bytes,5,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty" toml:"db_name,omitempty" mapstructure:"db_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Postgresql) Reset()         { *m = ConfigRequest_V1_Postgresql{} }
func (m *ConfigRequest_V1_Postgresql) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Postgresql) ProtoMessage()    {}
func (*ConfigRequest_V1_Postgresql) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 3}
}
func (m *ConfigRequest_V1_Postgresql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Postgresql.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Postgresql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Postgresql.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_Postgresql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Postgresql.Merge(dst, src)
}
func (m *ConfigRequest_V1_Postgresql) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Postgresql.Size(m)
}
func (m *ConfigRequest_V1_Postgresql) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Postgresql.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Postgresql proto.InternalMessageInfo

func (m *ConfigRequest_V1_Postgresql) GetDbName() *wrappers.StringValue {
	if m != nil {
		return m.DbName
	}
	return nil
}

type ConfigRequest_V1_SshGit struct {
	Hostname             *wrappers.StringValue `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty" toml:"hostname,omitempty" mapstructure:"hostname,omitempty"`
	HostAddress          *wrappers.StringValue `protobuf:"bytes,2,opt,name=host_address,json=hostAddress,proto3" json:"host_address,omitempty" toml:"host_address,omitempty" mapstructure:"host_address,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	GitRepoTemplate      *wrappers.StringValue `protobuf:"bytes,4,opt,name=git_repo_template,json=gitRepoTemplate,proto3" json:"git_repo_template,omitempty" toml:"git_repo_template,omitempty" mapstructure:"git_repo_template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_SshGit) Reset()         { *m = ConfigRequest_V1_SshGit{} }
func (m *ConfigRequest_V1_SshGit) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_SshGit) ProtoMessage()    {}
func (*ConfigRequest_V1_SshGit) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 4}
}
func (m *ConfigRequest_V1_SshGit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_SshGit.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_SshGit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_SshGit.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_SshGit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_SshGit.Merge(dst, src)
}
func (m *ConfigRequest_V1_SshGit) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_SshGit.Size(m)
}
func (m *ConfigRequest_V1_SshGit) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_SshGit.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_SshGit proto.InternalMessageInfo

func (m *ConfigRequest_V1_SshGit) GetHostname() *wrappers.StringValue {
	if m != nil {
		return m.Hostname
	}
	return nil
}

func (m *ConfigRequest_V1_SshGit) GetHostAddress() *wrappers.StringValue {
	if m != nil {
		return m.HostAddress
	}
	return nil
}

func (m *ConfigRequest_V1_SshGit) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_SshGit) GetGitRepoTemplate() *wrappers.StringValue {
	if m != nil {
		return m.GitRepoTemplate
	}
	return nil
}

type ConfigRequest_V1_Telemetry struct {
	Enabled              *wrappers.BoolValue   `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty" toml:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	Api                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty" toml:"api,omitempty" mapstructure:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Telemetry) Reset()         { *m = ConfigRequest_V1_Telemetry{} }
func (m *ConfigRequest_V1_Telemetry) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Telemetry) ProtoMessage()    {}
func (*ConfigRequest_V1_Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_5a380dee6c9cddc5, []int{0, 0, 5}
}
func (m *ConfigRequest_V1_Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Telemetry.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Telemetry.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Telemetry.Merge(dst, src)
}
func (m *ConfigRequest_V1_Telemetry) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Telemetry.Size(m)
}
func (m *ConfigRequest_V1_Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Telemetry proto.InternalMessageInfo

func (m *ConfigRequest_V1_Telemetry) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *ConfigRequest_V1_Telemetry) GetApi() *wrappers.StringValue {
	if m != nil {
		return m.Api
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.workflow_server.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.Service")
	proto.RegisterType((*ConfigRequest_V1_ChefServer)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.ChefServer")
	proto.RegisterType((*ConfigRequest_V1_Postgresql)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.Postgresql")
	proto.RegisterType((*ConfigRequest_V1_SshGit)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.SshGit")
	proto.RegisterType((*ConfigRequest_V1_Telemetry)(nil), "chef.automate.domain.workflow_server.ConfigRequest.V1.Telemetry")
}

func init() {
	proto.RegisterFile("api/config/workflow_server/config_request.proto", fileDescriptor_config_request_5a380dee6c9cddc5)
}

var fileDescriptor_config_request_5a380dee6c9cddc5 = []byte{
	// 909 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xcb, 0x4e, 0xe4, 0x46,
	0x14, 0x86, 0x45, 0xdf, 0x39, 0x0c, 0x97, 0xa9, 0x2c, 0x70, 0x3c, 0x11, 0x42, 0x51, 0x16, 0xa3,
	0x48, 0xb8, 0x03, 0xcc, 0x44, 0x99, 0x51, 0x6e, 0x03, 0x1a, 0x92, 0x21, 0x24, 0x83, 0xcc, 0x45,
	0x51, 0x16, 0xb1, 0xca, 0xf6, 0x69, 0x77, 0x29, 0x65, 0x97, 0xa9, 0x2a, 0xbb, 0xc3, 0x32, 0xeb,
	0x3c, 0x40, 0x56, 0x79, 0x19, 0xde, 0x26, 0x6b, 0xb2, 0xca, 0x2a, 0xaa, 0xb2, 0xbb, 0x61, 0x06,
	0x09, 0x2c, 0x58, 0xb5, 0xbb, 0xeb, 0xfc, 0xdf, 0x5f, 0x55, 0x7d, 0xfe, 0x23, 0xc3, 0x90, 0xe6,
	0x6c, 0x18, 0x89, 0x6c, 0xc4, 0x92, 0xe1, 0x44, 0xc8, 0xdf, 0x46, 0x5c, 0x4c, 0x02, 0x85, 0xb2,
	0x44, 0x59, 0xff, 0x1c, 0x48, 0x3c, 0x2b, 0x50, 0x69, 0x2f, 0x97, 0x42, 0x0b, 0xf2, 0x49, 0x34,
	0xc6, 0x91, 0x47, 0x0b, 0x2d, 0x52, 0xaa, 0xd1, 0x8b, 0x45, 0x4a, 0x59, 0xe6, 0xbd, 0x27, 0x75,
	0xd7, 0xae, 0x61, 0xd5, 0x98, 0x4a, 0x8c, 0x87, 0x09, 0x17, 0x21, 0xe5, 0x15, 0xc5, 0x7d, 0x72,
	0x73, 0x5d, 0x73, 0x55, 0x2f, 0xee, 0x47, 0x22, 0xcd, 0x45, 0x86, 0x99, 0x56, 0xc3, 0xa9, 0xd1,
	0x46, 0x22, 0xf3, 0x68, 0x68, 0xd7, 0xa3, 0x8d, 0x04, 0xb3, 0x0d, 0xba, 0xb5, 0x51, 0xeb, 0x0d,
	0x8a, 0x6e, 0x99, 0x2f, 0x43, 0x9a, 0x65, 0x42, 0x53, 0xcd, 0x44, 0x36, 0x65, 0xad, 0x25, 0x42,
	0x24, 0x1c, 0x2b, 0x65, 0x58, 0x8c, 0x86, 0x13, 0x49, 0xf3, 0x1c, 0x65, 0xbd, 0xfe, 0xf1, 0xbf,
	0xcb, 0xb0, 0xb8, 0x6b, 0x39, 0x7e, 0x75, 0x4c, 0xb2, 0x07, 0xad, 0x72, 0xd3, 0x99, 0x5b, 0x9f,
	0x7b, 0xba, 0xb0, 0xf5, 0xb9, 0xd7, 0xe4, 0xb4, 0xde, 0x3b, 0x00, 0xef, 0x74, 0xd3, 0x6f, 0x95,
	0x9b, 0xee, 0x7f, 0x4b, 0xd0, 0x3a, 0xdd, 0x24, 0x6f, 0xa1, 0xad, 0xce, 0x55, 0xcd, 0xfb, 0xea,
	0x7e, 0x3c, 0xef, 0xe8, 0x5c, 0x69, 0x4c, 0x7d, 0x43, 0x22, 0x87, 0xd0, 0x56, 0x65, 0xe4, 0xb4,
	0x2c, 0xf0, 0xeb, 0xfb, 0x02, 0x51, 0x96, 0x2c, 0x42, 0xdf, 0xa0, 0xdc, 0x3f, 0x06, 0xd0, 0xab,
	0x1c, 0xc8, 0x33, 0xe8, 0xa4, 0x5c, 0xd1, 0x7a, 0xbb, 0xeb, 0xef, 0xd1, 0x59, 0x36, 0x92, 0xd4,
	0xab, 0xae, 0xdd, 0xfb, 0x91, 0x2b, 0xea, 0xdb, 0x6a, 0x72, 0x04, 0x6d, 0x2e, 0x92, 0x7a, 0x4b,
	0xaf, 0x1e, 0x74, 0x46, 0xef, 0x40, 0x24, 0xbe, 0xa1, 0x91, 0x55, 0xe8, 0xc7, 0x58, 0x06, 0x98,
	0x95, 0x4e, 0x7b, 0x7d, 0xee, 0xe9, 0xc0, 0xef, 0xc5, 0x58, 0xbe, 0xce, 0x4a, 0xf2, 0x2b, 0xcc,
	0x6b, 0xe4, 0x98, 0xa2, 0x96, 0xe7, 0x4e, 0xc7, 0x7a, 0x7e, 0x7b, 0x4f, 0xcf, 0xe3, 0x29, 0xc7,
	0xbf, 0x42, 0x92, 0x10, 0x16, 0x0c, 0xad, 0x16, 0x39, 0xdd, 0x07, 0x9d, 0x6a, 0x77, 0x8c, 0xa3,
	0x23, 0xbb, 0xe6, 0x43, 0x34, 0x7b, 0x26, 0x14, 0x20, 0x17, 0x4a, 0x27, 0x12, 0xd5, 0x19, 0x77,
	0x7a, 0x0f, 0xb2, 0x38, 0x9c, 0x81, 0xfc, 0x6b, 0x50, 0x72, 0x0a, 0x7d, 0xa5, 0xc6, 0x41, 0xc2,
	0xb4, 0xd3, 0x7f, 0x58, 0xf3, 0xa9, 0xf1, 0x77, 0x4c, 0xfb, 0x3d, 0x65, 0x3f, 0xc9, 0x67, 0xd0,
	0x19, 0x9d, 0xc5, 0x99, 0x33, 0x6f, 0xa1, 0x1f, 0x79, 0x55, 0xc0, 0xbc, 0x69, 0xc0, 0xbc, 0x23,
	0x2d, 0x59, 0x96, 0x9c, 0x52, 0x5e, 0xa0, 0x6f, 0x2b, 0xc9, 0x5b, 0x18, 0xd0, 0x9c, 0x05, 0xb9,
	0x90, 0xda, 0x59, 0xb0, 0xaa, 0x27, 0x37, 0x54, 0x6f, 0x32, 0xbd, 0xbd, 0x65, 0x45, 0x3b, 0xab,
	0x17, 0x97, 0xce, 0x07, 0xd0, 0x57, 0x55, 0x8b, 0xae, 0xfc, 0xf9, 0x83, 0xdb, 0x1d, 0x6b, 0x9d,
	0x2b, 0xbf, 0x4f, 0x73, 0x76, 0x28, 0xa4, 0x26, 0x2f, 0x60, 0xde, 0x02, 0x8d, 0xd8, 0x79, 0xd4,
	0x60, 0x1f, 0xc6, 0xff, 0xd0, 0x8e, 0xaf, 0x5d, 0x58, 0x4a, 0x98, 0x0e, 0xf0, 0x77, 0x8c, 0x0a,
	0x4d, 0x43, 0x8e, 0xce, 0x62, 0x03, 0xfd, 0x62, 0xc2, 0xf4, 0xeb, 0x99, 0x84, 0xfc, 0x0c, 0x1f,
	0x6a, 0x59, 0x28, 0x8d, 0x71, 0x10, 0xa1, 0xd4, 0x6c, 0xc4, 0x22, 0xaa, 0x51, 0x05, 0x23, 0xc6,
	0xd1, 0x59, 0x6a, 0xc0, 0x5b, 0xad, 0xe5, 0xbb, 0xd7, 0xd4, 0x7b, 0x8c, 0x23, 0xf9, 0x12, 0xda,
	0x9a, 0x2b, 0x67, 0xd9, 0x32, 0x3e, 0xbd, 0x2d, 0x7e, 0xc7, 0x07, 0x47, 0xbb, 0x12, 0x63, 0xcc,
	0x34, 0xa3, 0x5c, 0xf9, 0x46, 0xe6, 0xbe, 0x80, 0xf6, 0x81, 0x48, 0xc8, 0x16, 0x74, 0x39, 0x96,
	0xc8, 0xeb, 0x14, 0xdf, 0xbe, 0x95, 0xaa, 0xd4, 0x9d, 0x87, 0x7e, 0x3d, 0x13, 0xdc, 0x7f, 0xe6,
	0x00, 0xae, 0xda, 0x96, 0x78, 0xd0, 0x2e, 0x64, 0x33, 0x96, 0x29, 0x24, 0x2f, 0x01, 0x26, 0x18,
	0x06, 0x05, 0x0b, 0x8c, 0xac, 0xd5, 0xe4, 0xdf, 0x99, 0x60, 0x78, 0xc2, 0x4e, 0x24, 0x37, 0x5e,
	0x25, 0xcb, 0x6d, 0xde, 0xef, 0xf4, 0x2a, 0x59, 0x6e, 0x1a, 0xc1, 0x46, 0xb5, 0x50, 0xb3, 0xa0,
	0xde, 0x61, 0x65, 0xca, 0x4f, 0x14, 0xca, 0xfd, 0xce, 0xa0, 0xb3, 0xd2, 0x75, 0xdf, 0x00, 0x5c,
	0xc5, 0x87, 0x3c, 0x87, 0x7e, 0x1c, 0x06, 0x19, 0x4d, 0xb1, 0x11, 0xac, 0x17, 0x87, 0x3f, 0xd1,
	0x14, 0xf7, 0x3b, 0x83, 0xb9, 0x95, 0xae, 0xfb, 0x57, 0x0b, 0x7a, 0x55, 0x54, 0xc8, 0x17, 0x30,
	0x18, 0x0b, 0xa5, 0x2d, 0xa8, 0xc9, 0xbd, 0xcd, 0xaa, 0xc9, 0x37, 0xf0, 0xc8, 0x3c, 0x07, 0x34,
	0x8e, 0x25, 0x2a, 0xd5, 0xe8, 0xfa, 0x16, 0x8c, 0xe2, 0x55, 0x25, 0x20, 0x3b, 0xd0, 0xb1, 0x39,
	0x6b, 0xdf, 0x9d, 0xb3, 0xc7, 0x17, 0x97, 0xce, 0x22, 0xb4, 0x13, 0xa6, 0x57, 0xfe, 0xde, 0x73,
	0xdb, 0x4a, 0x8d, 0x7d, 0xab, 0x25, 0xdf, 0xc3, 0x63, 0x93, 0x11, 0x89, 0xb9, 0x08, 0x34, 0xa6,
	0x39, 0xa7, 0x1a, 0xeb, 0x41, 0x7b, 0xfb, 0x4e, 0x96, 0x13, 0xa6, 0x7d, 0xcc, 0xc5, 0x71, 0x2d,
	0xda, 0xef, 0x0c, 0xba, 0x2b, 0x03, 0xf7, 0x0c, 0xe6, 0x67, 0x83, 0x96, 0x3c, 0x83, 0x3e, 0x66,
	0x26, 0x45, 0x71, 0x7d, 0x35, 0xee, 0x0d, 0xe4, 0x8e, 0x10, 0xbc, 0x02, 0x4e, 0x4b, 0x4d, 0x63,
	0xd0, 0x9c, 0x35, 0xba, 0x0e, 0x53, 0xf8, 0x72, 0xed, 0xe2, 0xd2, 0x71, 0xc1, 0x99, 0xbd, 0x3d,
	0x4c, 0x87, 0xdc, 0x46, 0x35, 0xe4, 0x76, 0x9e, 0xff, 0xb2, 0x9d, 0x30, 0x3d, 0x2e, 0x42, 0x2f,
	0x12, 0xe9, 0xd0, 0x34, 0xc5, 0xec, 0x4d, 0xe3, 0x96, 0x37, 0xa2, 0xb0, 0x67, 0x1d, 0xb7, 0xff,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x40, 0x70, 0x2f, 0x4f, 0x36, 0x09, 0x00, 0x00,
}
