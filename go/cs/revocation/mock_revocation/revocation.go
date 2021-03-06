// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/cs/revocation (interfaces: Store)

// Package mock_revocation is a generated GoMock package.
package mock_revocation

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	path_mgmt "github.com/scionproto/scion/go/lib/ctrl/path_mgmt"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// InsertRevocations mocks base method
func (m *MockStore) InsertRevocations(arg0 context.Context, arg1 ...*path_mgmt.SignedRevInfo) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertRevocations", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertRevocations indicates an expected call of InsertRevocations
func (mr *MockStoreMockRecorder) InsertRevocations(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRevocations", reflect.TypeOf((*MockStore)(nil).InsertRevocations), varargs...)
}
