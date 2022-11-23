// Code generated by MockGen. DO NOT EDIT.
// Source: services.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/zhuravlev-pe/course-watch/internal/core"
	service "github.com/zhuravlev-pe/course-watch/internal/service"
)

// MockCourses is a mock of Courses interface.
type MockCourses struct {
	ctrl     *gomock.Controller
	recorder *MockCoursesMockRecorder
}

// MockCoursesMockRecorder is the mock recorder for MockCourses.
type MockCoursesMockRecorder struct {
	mock *MockCourses
}

// NewMockCourses creates a new mock instance.
func NewMockCourses(ctrl *gomock.Controller) *MockCourses {
	mock := &MockCourses{ctrl: ctrl}
	mock.recorder = &MockCoursesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourses) EXPECT() *MockCoursesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCourses) Create(ctx context.Context, input service.CreateCourseInput) (*core.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*core.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCoursesMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCourses)(nil).Create), ctx, input)
}

// GetById mocks base method.
func (m *MockCourses) GetById(ctx context.Context, id string) (*core.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*core.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCoursesMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCourses)(nil).GetById), ctx, id)
}

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// GetUserInfo mocks base method.
func (m *MockUsers) GetUserInfo(ctx context.Context, id string) (*service.GetUserInfoOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, id)
	ret0, _ := ret[0].(*service.GetUserInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUsersMockRecorder) GetUserInfo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUsers)(nil).GetUserInfo), ctx, id)
}

// UpdateUserInfo mocks base method.
func (m *MockUsers) UpdateUserInfo(ctx context.Context, id string, input *service.UpdateUserInfoInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", ctx, id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockUsersMockRecorder) UpdateUserInfo(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUsers)(nil).UpdateUserInfo), ctx, id, input)
}