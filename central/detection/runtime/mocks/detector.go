// Code generated by MockGen. DO NOT EDIT.
// Source: detector.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	detection "github.com/stackrox/rox/central/detection"
	reflect "reflect"
)

// MockDetector is a mock of Detector interface
type MockDetector struct {
	ctrl     *gomock.Controller
	recorder *MockDetectorMockRecorder
}

// MockDetectorMockRecorder is the mock recorder for MockDetector
type MockDetectorMockRecorder struct {
	mock *MockDetector
}

// NewMockDetector creates a new mock instance
func NewMockDetector(ctrl *gomock.Controller) *MockDetector {
	mock := &MockDetector{ctrl: ctrl}
	mock.recorder = &MockDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDetector) EXPECT() *MockDetectorMockRecorder {
	return m.recorder
}

// PolicySet mocks base method
func (m *MockDetector) PolicySet() detection.PolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicySet")
	ret0, _ := ret[0].(detection.PolicySet)
	return ret0
}

// PolicySet indicates an expected call of PolicySet
func (mr *MockDetectorMockRecorder) PolicySet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicySet", reflect.TypeOf((*MockDetector)(nil).PolicySet))
}

// DeploymentWhitelistedForPolicy mocks base method
func (m *MockDetector) DeploymentWhitelistedForPolicy(deploymentID, policyID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentWhitelistedForPolicy", deploymentID, policyID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeploymentWhitelistedForPolicy indicates an expected call of DeploymentWhitelistedForPolicy
func (mr *MockDetectorMockRecorder) DeploymentWhitelistedForPolicy(deploymentID, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentWhitelistedForPolicy", reflect.TypeOf((*MockDetector)(nil).DeploymentWhitelistedForPolicy), deploymentID, policyID)
}

// DeploymentInactive mocks base method
func (m *MockDetector) DeploymentInactive(deploymentID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentInactive", deploymentID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeploymentInactive indicates an expected call of DeploymentInactive
func (mr *MockDetectorMockRecorder) DeploymentInactive(deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentInactive", reflect.TypeOf((*MockDetector)(nil).DeploymentInactive), deploymentID)
}
