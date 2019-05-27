// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/rules.proto

package v2beta

import (
	"context"

	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the RulesServer interface (at compile time)
var _ RulesServer = &RulesServerMock{}

// NewRulesServerMock gives you a fresh instance of RulesServerMock.
func NewRulesServerMock() *RulesServerMock {
	return &RulesServerMock{validateRequests: true}
}

// NewRulesServerMockWithoutValidation gives you a fresh instance of
// RulesServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewRulesServerMockWithoutValidation() *RulesServerMock {
	return &RulesServerMock{}
}

// RulesServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type RulesServerMock struct {
	validateRequests bool
	CreateRuleFunc   func(context.Context, *request.CreateRuleReq) (*response.CreateRuleResp, error)
}

func (m *RulesServerMock) CreateRule(ctx context.Context, req *request.CreateRuleReq) (*response.CreateRuleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateRuleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateRule' not implemented")
}

// Reset resets all overridden functions
func (m *RulesServerMock) Reset() {
	m.CreateRuleFunc = nil
}
