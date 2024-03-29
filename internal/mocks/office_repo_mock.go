// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonmp/bss-office-api/internal/repo (interfaces: OfficeRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	model "github.com/ozonmp/bss-office-api/internal/model"
)

// MockOfficeRepo is a mock of OfficeRepo interface.
type MockOfficeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOfficeRepoMockRecorder
}

// MockOfficeRepoMockRecorder is the mock recorder for MockOfficeRepo.
type MockOfficeRepoMockRecorder struct {
	mock *MockOfficeRepo
}

// NewMockOfficeRepo creates a new mock instance.
func NewMockOfficeRepo(ctrl *gomock.Controller) *MockOfficeRepo {
	mock := &MockOfficeRepo{ctrl: ctrl}
	mock.recorder = &MockOfficeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfficeRepo) EXPECT() *MockOfficeRepoMockRecorder {
	return m.recorder
}

// CreateOffice mocks base method.
func (m *MockOfficeRepo) CreateOffice(arg0 context.Context, arg1 model.Office, arg2 *sqlx.Tx) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOffice", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOffice indicates an expected call of CreateOffice.
func (mr *MockOfficeRepoMockRecorder) CreateOffice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOffice", reflect.TypeOf((*MockOfficeRepo)(nil).CreateOffice), arg0, arg1, arg2)
}

// DescribeOffice mocks base method.
func (m *MockOfficeRepo) DescribeOffice(arg0 context.Context, arg1 uint64) (*model.Office, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeOffice", arg0, arg1)
	ret0, _ := ret[0].(*model.Office)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOffice indicates an expected call of DescribeOffice.
func (mr *MockOfficeRepoMockRecorder) DescribeOffice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOffice", reflect.TypeOf((*MockOfficeRepo)(nil).DescribeOffice), arg0, arg1)
}

// ListOffices mocks base method.
func (m *MockOfficeRepo) ListOffices(arg0 context.Context, arg1, arg2 uint64) ([]*model.Office, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOffices", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Office)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOffices indicates an expected call of ListOffices.
func (mr *MockOfficeRepoMockRecorder) ListOffices(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOffices", reflect.TypeOf((*MockOfficeRepo)(nil).ListOffices), arg0, arg1, arg2)
}

// RemoveOffice mocks base method.
func (m *MockOfficeRepo) RemoveOffice(arg0 context.Context, arg1 uint64, arg2 *sqlx.Tx) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOffice", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveOffice indicates an expected call of RemoveOffice.
func (mr *MockOfficeRepoMockRecorder) RemoveOffice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOffice", reflect.TypeOf((*MockOfficeRepo)(nil).RemoveOffice), arg0, arg1, arg2)
}

// UpdateOffice mocks base method.
func (m *MockOfficeRepo) UpdateOffice(arg0 context.Context, arg1 uint64, arg2 model.Office, arg3 *sqlx.Tx) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOffice", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOffice indicates an expected call of UpdateOffice.
func (mr *MockOfficeRepoMockRecorder) UpdateOffice(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOffice", reflect.TypeOf((*MockOfficeRepo)(nil).UpdateOffice), arg0, arg1, arg2, arg3)
}

// UpdateOfficeDescription mocks base method.
func (m *MockOfficeRepo) UpdateOfficeDescription(arg0 context.Context, arg1 uint64, arg2 string, arg3 *sqlx.Tx) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfficeDescription", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfficeDescription indicates an expected call of UpdateOfficeDescription.
func (mr *MockOfficeRepoMockRecorder) UpdateOfficeDescription(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfficeDescription", reflect.TypeOf((*MockOfficeRepo)(nil).UpdateOfficeDescription), arg0, arg1, arg2, arg3)
}

// UpdateOfficeName mocks base method.
func (m *MockOfficeRepo) UpdateOfficeName(arg0 context.Context, arg1 uint64, arg2 string, arg3 *sqlx.Tx) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfficeName", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfficeName indicates an expected call of UpdateOfficeName.
func (mr *MockOfficeRepoMockRecorder) UpdateOfficeName(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfficeName", reflect.TypeOf((*MockOfficeRepo)(nil).UpdateOfficeName), arg0, arg1, arg2, arg3)
}
