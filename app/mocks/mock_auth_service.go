// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shasw94/projX/app/interfaces (interfaces: IAuthService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	schema "github.com/shasw94/projX/app/schema"
)

// MockIAuthService is a mock of IAuthService interface.
type MockIAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthServiceMockRecorder
}

// MockIAuthServiceMockRecorder is the mock recorder for MockIAuthService.
type MockIAuthServiceMockRecorder struct {
	mock *MockIAuthService
}

// NewMockIAuthService creates a new mock instance.
func NewMockIAuthService(ctrl *gomock.Controller) *MockIAuthService {
	mock := &MockIAuthService{ctrl: ctrl}
	mock.recorder = &MockIAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthService) EXPECT() *MockIAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockIAuthService) Login(arg0 context.Context, arg1 *schema.LoginBodyParams) (*schema.UserTokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*schema.UserTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIAuthServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIAuthService)(nil).Login), arg0, arg1)
}

// Logout mocks base method.
func (m *MockIAuthService) Logout(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockIAuthServiceMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockIAuthService)(nil).Logout), arg0)
}

// Refresh mocks base method.
func (m *MockIAuthService) Refresh(arg0 context.Context, arg1 *schema.RefreshBodyParams) (*schema.UserTokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0, arg1)
	ret0, _ := ret[0].(*schema.UserTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockIAuthServiceMockRecorder) Refresh(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockIAuthService)(nil).Refresh), arg0, arg1)
}

// Register mocks base method.
func (m *MockIAuthService) Register(arg0 context.Context, arg1 *schema.RegisterBodyParams) (*schema.UserTokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*schema.UserTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockIAuthServiceMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAuthService)(nil).Register), arg0, arg1)
}
