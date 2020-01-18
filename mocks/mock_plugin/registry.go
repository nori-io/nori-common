// Code generated by MockGen. DO NOT EDIT.
// Source: plugin/registry.go

// Package mock_plugin is a generated GoMock package.
package mock_plugin

import (
	gomock "github.com/golang/mock/gomock"
	logger "github.com/nori-io/nori-common/logger"
	meta "github.com/nori-io/nori-common/meta"
	reflect "reflect"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockRegistry) ID(id meta.ID) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", id)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *MockRegistryMockRecorder) ID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRegistry)(nil).ID), id)
}

// Interface mocks base method
func (m *MockRegistry) Interface(i meta.Interface) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Interface", i)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Interface indicates an expected call of Interface
func (mr *MockRegistryMockRecorder) Interface(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Interface", reflect.TypeOf((*MockRegistry)(nil).Interface), i)
}

// Resolve mocks base method
func (m *MockRegistry) Resolve(dep meta.Dependency) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", dep)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockRegistryMockRecorder) Resolve(dep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockRegistry)(nil).Resolve), dep)
}

// Logger mocks base method
func (m *MockRegistry) Logger() logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockRegistryMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockRegistry)(nil).Logger))
}
