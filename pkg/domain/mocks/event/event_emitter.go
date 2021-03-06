// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v5/pkg/domain/event (interfaces: EventEmitter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	event "github.com/nori-io/common/v5/pkg/domain/event"
)

// MockEventEmitter is a mock of EventEmitter interface.
type MockEventEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockEventEmitterMockRecorder
}

// MockEventEmitterMockRecorder is the mock recorder for MockEventEmitter.
type MockEventEmitterMockRecorder struct {
	mock *MockEventEmitter
}

// NewMockEventEmitter creates a new mock instance.
func NewMockEventEmitter(ctrl *gomock.Controller) *MockEventEmitter {
	mock := &MockEventEmitter{ctrl: ctrl}
	mock.recorder = &MockEventEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventEmitter) EXPECT() *MockEventEmitterMockRecorder {
	return m.recorder
}

// Emit mocks base method.
func (m *MockEventEmitter) Emit(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Emit", arg0, arg1)
}

// Emit indicates an expected call of Emit.
func (mr *MockEventEmitterMockRecorder) Emit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockEventEmitter)(nil).Emit), arg0, arg1)
}

// On mocks base method.
func (m *MockEventEmitter) On(arg0 string, arg1 ...event.Middleware) (<-chan event.Event, func()) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "On", varargs...)
	ret0, _ := ret[0].(<-chan event.Event)
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// On indicates an expected call of On.
func (mr *MockEventEmitterMockRecorder) On(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "On", reflect.TypeOf((*MockEventEmitter)(nil).On), varargs...)
}

// Use mocks base method.
func (m *MockEventEmitter) Use(arg0 string, arg1 ...event.Middleware) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Use", varargs...)
}

// Use indicates an expected call of Use.
func (mr *MockEventEmitterMockRecorder) Use(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockEventEmitter)(nil).Use), varargs...)
}
