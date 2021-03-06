// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v5/pkg/domain/meta (interfaces: Description)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDescription is a mock of Description interface.
type MockDescription struct {
	ctrl     *gomock.Controller
	recorder *MockDescriptionMockRecorder
}

// MockDescriptionMockRecorder is the mock recorder for MockDescription.
type MockDescriptionMockRecorder struct {
	mock *MockDescription
}

// NewMockDescription creates a new mock instance.
func NewMockDescription(ctrl *gomock.Controller) *MockDescription {
	mock := &MockDescription{ctrl: ctrl}
	mock.recorder = &MockDescriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDescription) EXPECT() *MockDescriptionMockRecorder {
	return m.recorder
}

// GetDescription mocks base method.
func (m *MockDescription) GetDescription() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDescription indicates an expected call of GetDescription.
func (mr *MockDescriptionMockRecorder) GetDescription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescription", reflect.TypeOf((*MockDescription)(nil).GetDescription))
}

// GetTitle mocks base method.
func (m *MockDescription) GetTitle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTitle indicates an expected call of GetTitle.
func (mr *MockDescriptionMockRecorder) GetTitle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockDescription)(nil).GetTitle))
}
