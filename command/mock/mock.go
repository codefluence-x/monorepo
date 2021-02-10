// Code generated by MockGen. DO NOT EDIT.
// Source: ./command.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockScaffold is a mock of Scaffold interface
type MockScaffold struct {
	ctrl     *gomock.Controller
	recorder *MockScaffoldMockRecorder
}

// MockScaffoldMockRecorder is the mock recorder for MockScaffold
type MockScaffoldMockRecorder struct {
	mock *MockScaffold
}

// NewMockScaffold creates a new mock instance
func NewMockScaffold(ctrl *gomock.Controller) *MockScaffold {
	mock := &MockScaffold{ctrl: ctrl}
	mock.recorder = &MockScaffoldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScaffold) EXPECT() *MockScaffoldMockRecorder {
	return m.recorder
}

// Use mocks base method
func (m *MockScaffold) Use() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Use")
	ret0, _ := ret[0].(string)
	return ret0
}

// Use indicates an expected call of Use
func (mr *MockScaffoldMockRecorder) Use() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockScaffold)(nil).Use))
}

// Example mocks base method
func (m *MockScaffold) Example() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Example")
	ret0, _ := ret[0].(string)
	return ret0
}

// Example indicates an expected call of Example
func (mr *MockScaffoldMockRecorder) Example() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Example", reflect.TypeOf((*MockScaffold)(nil).Example))
}

// Short mocks base method
func (m *MockScaffold) Short() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Short")
	ret0, _ := ret[0].(string)
	return ret0
}

// Short indicates an expected call of Short
func (mr *MockScaffoldMockRecorder) Short() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Short", reflect.TypeOf((*MockScaffold)(nil).Short))
}

// Run mocks base method
func (m *MockScaffold) Run(args []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", args)
}

// Run indicates an expected call of Run
func (mr *MockScaffoldMockRecorder) Run(args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockScaffold)(nil).Run), args)
}
