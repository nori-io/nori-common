// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/plugin (interfaces: Installable)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	registry "github.com/nori-io/common/v4/pkg/domain/registry"
	reflect "reflect"
)

// MockInstallable is a mock of Installable interface
type MockInstallable struct {
	ctrl     *gomock.Controller
	recorder *MockInstallableMockRecorder
}

// MockInstallableMockRecorder is the mock recorder for MockInstallable
type MockInstallableMockRecorder struct {
	mock *MockInstallable
}

// NewMockInstallable creates a new mock instance
func NewMockInstallable(ctrl *gomock.Controller) *MockInstallable {
	mock := &MockInstallable{ctrl: ctrl}
	mock.recorder = &MockInstallableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstallable) EXPECT() *MockInstallableMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockInstallable) Install(arg0 context.Context, arg1 registry.Registry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockInstallableMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockInstallable)(nil).Install), arg0, arg1)
}

// UnInstall mocks base method
func (m *MockInstallable) UnInstall(arg0 context.Context, arg1 registry.Registry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnInstall", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnInstall indicates an expected call of UnInstall
func (mr *MockInstallableMockRecorder) UnInstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnInstall", reflect.TypeOf((*MockInstallable)(nil).UnInstall), arg0, arg1)
}
