// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/pkg/sac/client (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	payload "github.com/stackrox/default-authz-plugin/pkg/payload"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ForUser mocks base method
func (m *MockClient) ForUser(arg0 context.Context, arg1 payload.Principal, arg2 ...payload.AccessScope) ([]payload.AccessScope, []payload.AccessScope, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ForUser", varargs...)
	ret0, _ := ret[0].([]payload.AccessScope)
	ret1, _ := ret[1].([]payload.AccessScope)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ForUser indicates an expected call of ForUser
func (mr *MockClientMockRecorder) ForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForUser", reflect.TypeOf((*MockClient)(nil).ForUser), varargs...)
}
