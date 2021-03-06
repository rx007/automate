// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/nodes/nodes.proto

package nodes

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the NodesServiceServer interface (at compile time)
var _ NodesServiceServer = &NodesServiceServerMock{}

// NewNodesServiceServerMock gives you a fresh instance of NodesServiceServerMock.
func NewNodesServiceServerMock() *NodesServiceServerMock {
	return &NodesServiceServerMock{validateRequests: true}
}

// NewNodesServiceServerMockWithoutValidation gives you a fresh instance of
// NodesServiceServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewNodesServiceServerMockWithoutValidation() *NodesServiceServerMock {
	return &NodesServiceServerMock{}
}

// NodesServiceServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type NodesServiceServerMock struct {
	validateRequests   bool
	CreateFunc         func(context.Context, *Node) (*Id, error)
	ReadFunc           func(context.Context, *Id) (*Node, error)
	UpdateFunc         func(context.Context, *Node) (*empty.Empty, error)
	DeleteFunc         func(context.Context, *Id) (*empty.Empty, error)
	BulkDeleteByIdFunc func(context.Context, *Ids) (*BulkDeleteResponse, error)
	ListFunc           func(context.Context, *Query) (*Nodes, error)
	RerunFunc          func(context.Context, *Id) (*RerunResponse, error)
	BulkDeleteFunc     func(context.Context, *Query) (*BulkDeleteResponse, error)
	BulkCreateFunc     func(context.Context, *Nodes) (*Ids, error)
}

func (m *NodesServiceServerMock) Create(ctx context.Context, req *Node) (*Id, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'Create' not implemented")
}

func (m *NodesServiceServerMock) Read(ctx context.Context, req *Id) (*Node, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ReadFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'Read' not implemented")
}

func (m *NodesServiceServerMock) Update(ctx context.Context, req *Node) (*empty.Empty, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'Update' not implemented")
}

func (m *NodesServiceServerMock) Delete(ctx context.Context, req *Id) (*empty.Empty, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'Delete' not implemented")
}

func (m *NodesServiceServerMock) BulkDeleteById(ctx context.Context, req *Ids) (*BulkDeleteResponse, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.BulkDeleteByIdFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'BulkDeleteById' not implemented")
}

func (m *NodesServiceServerMock) List(ctx context.Context, req *Query) (*Nodes, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'List' not implemented")
}

func (m *NodesServiceServerMock) Rerun(ctx context.Context, req *Id) (*RerunResponse, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.RerunFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'Rerun' not implemented")
}

func (m *NodesServiceServerMock) BulkDelete(ctx context.Context, req *Query) (*BulkDeleteResponse, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.BulkDeleteFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'BulkDelete' not implemented")
}

func (m *NodesServiceServerMock) BulkCreate(ctx context.Context, req *Nodes) (*Ids, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.BulkCreateFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'BulkCreate' not implemented")
}

// Reset resets all overridden functions
func (m *NodesServiceServerMock) Reset() {
	m.CreateFunc = nil
	m.ReadFunc = nil
	m.UpdateFunc = nil
	m.DeleteFunc = nil
	m.BulkDeleteByIdFunc = nil
	m.ListFunc = nil
	m.RerunFunc = nil
	m.BulkDeleteFunc = nil
	m.BulkCreateFunc = nil
}
