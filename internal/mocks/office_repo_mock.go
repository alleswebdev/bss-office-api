// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonmp/bss-office-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ozonmp/bss-office-api/internal/model"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// CreateOffice mocks base method.
func (m *MockRepo) CreateOffice(arg0 context.Context, arg1 model.Office) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOffice", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOffice indicates an expected call of CreateOffice.
func (mr *MockRepoMockRecorder) CreateOffice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOffice", reflect.TypeOf((*MockRepo)(nil).CreateOffice), arg0, arg1)
}

// DescribeOffice mocks base method.
func (m *MockRepo) DescribeOffice(arg0 context.Context, arg1 uint64) (*model.Office, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeOffice", arg0, arg1)
	ret0, _ := ret[0].(*model.Office)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOffice indicates an expected call of DescribeOffice.
func (mr *MockRepoMockRecorder) DescribeOffice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOffice", reflect.TypeOf((*MockRepo)(nil).DescribeOffice), arg0, arg1)
}

// ListOffices mocks base method.
func (m *MockRepo) ListOffices(arg0 context.Context) ([]model.Office, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOffices", arg0)
	ret0, _ := ret[0].([]model.Office)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOffices indicates an expected call of ListOffices.
func (mr *MockRepoMockRecorder) ListOffices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOffices", reflect.TypeOf((*MockRepo)(nil).ListOffices), arg0)
}

// RemoveOffice mocks base method.
func (m *MockRepo) RemoveOffice(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOffice", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveOffice indicates an expected call of RemoveOffice.
func (mr *MockRepoMockRecorder) RemoveOffice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOffice", reflect.TypeOf((*MockRepo)(nil).RemoveOffice), arg0, arg1)
}
