// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/ingest/job_scheduler.proto

package ingest // import "github.com/chef/automate/api/external/ingest"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import request "github.com/chef/automate/api/external/ingest/request"
import response "github.com/chef/automate/api/external/ingest/response"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobSchedulerClient is the client API for JobScheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobSchedulerClient interface {
	GetStatusJobScheduler(ctx context.Context, in *request.GetStatusJobScheduler, opts ...grpc.CallOption) (*response.JobSchedulerStatus, error)
	ConfigureNodesMissingScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureNodesMissingScheduler, error)
	ConfigureDeleteNodesScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureDeleteNodesScheduler, error)
	ConfigureMissingNodesForDeletionScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureMissingNodesForDeletionScheduler, error)
}

type jobSchedulerClient struct {
	cc *grpc.ClientConn
}

func NewJobSchedulerClient(cc *grpc.ClientConn) JobSchedulerClient {
	return &jobSchedulerClient{cc}
}

func (c *jobSchedulerClient) GetStatusJobScheduler(ctx context.Context, in *request.GetStatusJobScheduler, opts ...grpc.CallOption) (*response.JobSchedulerStatus, error) {
	out := new(response.JobSchedulerStatus)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/GetStatusJobScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSchedulerClient) ConfigureNodesMissingScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureNodesMissingScheduler, error) {
	out := new(response.ConfigureNodesMissingScheduler)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/ConfigureNodesMissingScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSchedulerClient) ConfigureDeleteNodesScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureDeleteNodesScheduler, error) {
	out := new(response.ConfigureDeleteNodesScheduler)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/ConfigureDeleteNodesScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSchedulerClient) ConfigureMissingNodesForDeletionScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureMissingNodesForDeletionScheduler, error) {
	out := new(response.ConfigureMissingNodesForDeletionScheduler)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/ConfigureMissingNodesForDeletionScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobSchedulerServer is the server API for JobScheduler service.
type JobSchedulerServer interface {
	GetStatusJobScheduler(context.Context, *request.GetStatusJobScheduler) (*response.JobSchedulerStatus, error)
	ConfigureNodesMissingScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureNodesMissingScheduler, error)
	ConfigureDeleteNodesScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureDeleteNodesScheduler, error)
	ConfigureMissingNodesForDeletionScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureMissingNodesForDeletionScheduler, error)
}

func RegisterJobSchedulerServer(s *grpc.Server, srv JobSchedulerServer) {
	s.RegisterService(&_JobScheduler_serviceDesc, srv)
}

func _JobScheduler_GetStatusJobScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetStatusJobScheduler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).GetStatusJobScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/GetStatusJobScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).GetStatusJobScheduler(ctx, req.(*request.GetStatusJobScheduler))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobScheduler_ConfigureNodesMissingScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SchedulerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).ConfigureNodesMissingScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/ConfigureNodesMissingScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).ConfigureNodesMissingScheduler(ctx, req.(*request.SchedulerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobScheduler_ConfigureDeleteNodesScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SchedulerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).ConfigureDeleteNodesScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/ConfigureDeleteNodesScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).ConfigureDeleteNodesScheduler(ctx, req.(*request.SchedulerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobScheduler_ConfigureMissingNodesForDeletionScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SchedulerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).ConfigureMissingNodesForDeletionScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/ConfigureMissingNodesForDeletionScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).ConfigureMissingNodesForDeletionScheduler(ctx, req.(*request.SchedulerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobScheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.ingest.JobScheduler",
	HandlerType: (*JobSchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatusJobScheduler",
			Handler:    _JobScheduler_GetStatusJobScheduler_Handler,
		},
		{
			MethodName: "ConfigureNodesMissingScheduler",
			Handler:    _JobScheduler_ConfigureNodesMissingScheduler_Handler,
		},
		{
			MethodName: "ConfigureDeleteNodesScheduler",
			Handler:    _JobScheduler_ConfigureDeleteNodesScheduler_Handler,
		},
		{
			MethodName: "ConfigureMissingNodesForDeletionScheduler",
			Handler:    _JobScheduler_ConfigureMissingNodesForDeletionScheduler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/external/ingest/job_scheduler.proto",
}

func init() {
	proto.RegisterFile("api/external/ingest/job_scheduler.proto", fileDescriptor_job_scheduler_93f2a7630365589d)
}

var fileDescriptor_job_scheduler_93f2a7630365589d = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x69, 0x94, 0x42, 0x47, 0x41, 0x8c, 0x54, 0x63, 0xa8, 0xab, 0x06, 0xb5, 0x5a, 0xcc,
	0x4c, 0x55, 0x44, 0xc8, 0xa9, 0xfe, 0x41, 0x41, 0xd4, 0x4b, 0x6f, 0x5e, 0x64, 0x92, 0xbc, 0x3b,
	0x3b, 0x25, 0x99, 0x89, 0x99, 0x09, 0xd8, 0x6b, 0x8e, 0xbd, 0x7a, 0x12, 0x3f, 0x83, 0xe0, 0xc1,
	0x7c, 0x12, 0xaf, 0x3d, 0xfa, 0x41, 0x24, 0x93, 0xdd, 0xd8, 0x6e, 0xb2, 0xbb, 0x96, 0xdd, 0x6b,
	0xde, 0xf7, 0x79, 0x9e, 0x79, 0x7e, 0x19, 0x06, 0x6d, 0xd3, 0x8c, 0x13, 0xf8, 0xa2, 0x21, 0x17,
	0x34, 0x21, 0x5c, 0x30, 0x50, 0x9a, 0x1c, 0xc8, 0xf0, 0x93, 0x8a, 0x46, 0x10, 0x17, 0x09, 0xe4,
	0x38, 0xcb, 0xa5, 0x96, 0xb6, 0x13, 0x8d, 0x60, 0x88, 0x69, 0xa1, 0x65, 0x4a, 0x35, 0x60, 0x9a,
	0x71, 0xdc, 0x6c, 0xbb, 0x5b, 0x4c, 0x4a, 0x96, 0x00, 0xa9, 0x9d, 0xa8, 0x10, 0x52, 0x53, 0xcd,
	0xa5, 0x50, 0x8d, 0xce, 0xdd, 0x8b, 0x64, 0x9a, 0x49, 0x01, 0x42, 0x2b, 0x32, 0x51, 0xfb, 0x2c,
	0xcf, 0x22, 0x62, 0xe6, 0x91, 0xcf, 0x40, 0xf8, 0x99, 0x4c, 0x78, 0x74, 0xb8, 0x02, 0x07, 0x4e,
	0xd3, 0x1e, 0x87, 0xdd, 0xbe, 0x92, 0x39, 0xa8, 0x4c, 0x0a, 0x05, 0x7d, 0x6d, 0x5d, 0xd2, 0xaf,
	0xf8, 0x5c, 0xcc, 0xc0, 0xf3, 0xf8, 0x78, 0x03, 0x5d, 0x7c, 0x2b, 0xc3, 0xfd, 0xc9, 0x67, 0xbb,
	0xb4, 0xd0, 0xe6, 0x1b, 0xd0, 0xfb, 0x9a, 0xea, 0x42, 0x9d, 0x9a, 0x3c, 0xc3, 0xb3, 0x50, 0xe2,
	0x71, 0x02, 0xee, 0x15, 0xba, 0x4f, 0xe7, 0x09, 0x9b, 0x32, 0xf8, 0xa4, 0xa0, 0xb1, 0xf0, 0x0e,
	0xcb, 0xca, 0x19, 0xa0, 0xad, 0x68, 0xc8, 0x52, 0x96, 0xea, 0xa0, 0x3d, 0x79, 0x70, 0x20, 0xc3,
	0x40, 0x99, 0x9d, 0xb2, 0x72, 0xd6, 0xed, 0xf3, 0x39, 0xd0, 0xf8, 0xa8, 0x72, 0x2e, 0xa3, 0x4b,
	0x39, 0x68, 0x10, 0x35, 0xc1, 0x40, 0xc8, 0x18, 0xd4, 0x51, 0xe5, 0x6c, 0xda, 0x57, 0xa6, 0x3e,
	0x06, 0x0c, 0x74, 0xf9, 0xfb, 0xcf, 0x57, 0xeb, 0xba, 0x7d, 0x8d, 0xb4, 0x33, 0x62, 0x66, 0xa4,
	0xb1, 0xb5, 0x7f, 0x59, 0x68, 0xf0, 0x52, 0x8a, 0x21, 0x67, 0x45, 0x0e, 0x1f, 0xea, 0xc9, 0x7b,
	0xae, 0x14, 0x17, 0xec, 0x1f, 0x8d, 0x47, 0x8b, 0x69, 0xb4, 0xcb, 0x8d, 0x95, 0xfb, 0xfc, 0x3f,
	0x38, 0xcc, 0x4f, 0xf5, 0xbe, 0xaf, 0x95, 0x95, 0x73, 0x1b, 0xdd, 0xec, 0x42, 0x49, 0x9b, 0x55,
	0xdf, 0xf4, 0x28, 0x2b, 0xe7, 0x82, 0xbd, 0x11, 0x4d, 0xbc, 0x66, 0xc1, 0x71, 0xec, 0xab, 0xd3,
	0x70, 0x8a, 0x2c, 0xa6, 0x1a, 0x0c, 0x9f, 0x6d, 0xef, 0x6e, 0x87, 0xcf, 0xa9, 0x14, 0xd2, 0x04,
	0x98, 0xe5, 0x73, 0xc1, 0xda, 0x8e, 0xfd, 0xd3, 0x42, 0x37, 0xda, 0x02, 0xaf, 0x20, 0x01, 0xdd,
	0xd4, 0x58, 0x8a, 0xda, 0xde, 0x59, 0xa8, 0xf5, 0x85, 0x7a, 0xdf, 0x6a, 0x68, 0xb7, 0xd0, 0xa0,
	0x0b, 0x2d, 0x36, 0xeb, 0x2b, 0x64, 0x76, 0xcf, 0xbb, 0xd3, 0x61, 0x76, 0x32, 0xa4, 0x83, 0xec,
	0xd8, 0x42, 0x0f, 0xda, 0xd3, 0x8f, 0x7f, 0xb7, 0x39, 0xfe, 0x6b, 0x99, 0x9b, 0x32, 0x5c, 0x8a,
	0xa5, 0xf0, 0xbd, 0x3b, 0x0b, 0xbe, 0x45, 0x07, 0xf0, 0x7e, 0xd4, 0x28, 0x77, 0xd0, 0xfd, 0x05,
	0xf7, 0xcf, 0x8f, 0xc7, 0xca, 0x95, 0x40, 0xdd, 0xf5, 0xf0, 0xfc, 0x8b, 0xd8, 0xc6, 0x4d, 0xe3,
	0x7d, 0x81, 0x3f, 0x3e, 0x64, 0x5c, 0x8f, 0x8a, 0x10, 0x47, 0x32, 0x25, 0x35, 0x89, 0xf6, 0x29,
	0xee, 0x7b, 0x2a, 0xc3, 0x75, 0xf3, 0x2a, 0x3e, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x70,
	0xf7, 0x63, 0x5f, 0x06, 0x00, 0x00,
}
