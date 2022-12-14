// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shasw94/projX/app/interfaces (interfaces: IRoleService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/shasw94/projX/app/models"
	schema "github.com/shasw94/projX/app/schema"
)

// MockIRoleService is a mock of IRoleService interface.
type MockIRoleService struct {
	ctrl     *gomock.Controller
	recorder *MockIRoleServiceMockRecorder
}

// MockIRoleServiceMockRecorder is the mock recorder for MockIRoleService.
type MockIRoleServiceMockRecorder struct {
	mock *MockIRoleService
}

// NewMockIRoleService creates a new mock instance.
func NewMockIRoleService(ctrl *gomock.Controller) *MockIRoleService {
	mock := &MockIRoleService{ctrl: ctrl}
	mock.recorder = &MockIRoleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoleService) EXPECT() *MockIRoleServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRoleService) Create(arg0 context.Context, arg1 *schema.RoleBodyParams) (*models.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*models.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIRoleServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRoleService)(nil).Create), arg0, arg1)
}
