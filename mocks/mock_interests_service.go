// Code generated by MockGen. DO NOT EDIT.
// Source: interests_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	golaerror "github.com/gola-glitch/gola-utils/golaerror"
	gomock "github.com/golang/mock/gomock"
	db "post-api/models/db"
	reflect "reflect"
)

// MockInterestsService is a mock of InterestsService interface
type MockInterestsService struct {
	ctrl     *gomock.Controller
	recorder *MockInterestsServiceMockRecorder
}

// MockInterestsServiceMockRecorder is the mock recorder for MockInterestsService
type MockInterestsServiceMockRecorder struct {
	mock *MockInterestsService
}

// NewMockInterestsService creates a new mock instance
func NewMockInterestsService(ctrl *gomock.Controller) *MockInterestsService {
	mock := &MockInterestsService{ctrl: ctrl}
	mock.recorder = &MockInterestsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterestsService) EXPECT() *MockInterestsServiceMockRecorder {
	return m.recorder
}

// GetInterests mocks base method
func (m *MockInterestsService) GetInterests(ctx context.Context) ([]db.Interest, *golaerror.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterests", ctx)
	ret0, _ := ret[0].([]db.Interest)
	ret1, _ := ret[1].(*golaerror.Error)
	return ret0, ret1
}

// GetInterests indicates an expected call of GetInterests
func (mr *MockInterestsServiceMockRecorder) GetInterests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterests", reflect.TypeOf((*MockInterestsService)(nil).GetInterests), ctx)
}
