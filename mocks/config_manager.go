// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v3/config (interfaces: Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/nori-io/common/v3/config"
	meta "github.com/nori-io/common/v3/meta"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// PluginVariables mocks base method
func (m *MockManager) PluginVariables(arg0 meta.ID) []config.Variable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluginVariables", arg0)
	ret0, _ := ret[0].([]config.Variable)
	return ret0
}

// PluginVariables indicates an expected call of PluginVariables
func (mr *MockManagerMockRecorder) PluginVariables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluginVariables", reflect.TypeOf((*MockManager)(nil).PluginVariables), arg0)
}

// Register mocks base method
func (m *MockManager) Register(arg0 meta.ID) config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(config.Config)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockManagerMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockManager)(nil).Register), arg0)
}
