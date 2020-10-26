// Code generated by MockGen. DO NOT EDIT.
// Source: posts_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	db "post-api/models/db"
	reflect "reflect"
)

// MockPostsRepository is a mock of PostsRepository interface
type MockPostsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostsRepositoryMockRecorder
}

// MockPostsRepositoryMockRecorder is the mock recorder for MockPostsRepository
type MockPostsRepositoryMockRecorder struct {
	mock *MockPostsRepository
}

// NewMockPostsRepository creates a new mock instance
func NewMockPostsRepository(ctrl *gomock.Controller) *MockPostsRepository {
	mock := &MockPostsRepository{ctrl: ctrl}
	mock.recorder = &MockPostsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostsRepository) EXPECT() *MockPostsRepositoryMockRecorder {
	return m.recorder
}

// CreatePost mocks base method
func (m *MockPostsRepository) CreatePost(ctx context.Context, post db.PublishPost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockPostsRepositoryMockRecorder) CreatePost(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostsRepository)(nil).CreatePost), ctx, post)
}
