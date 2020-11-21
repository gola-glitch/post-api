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
func (m *MockPostsRepository) CreatePost(ctx context.Context, post db.PublishPost) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, post)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockPostsRepositoryMockRecorder) CreatePost(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostsRepository)(nil).CreatePost), ctx, post)
}

// GetLikeCountByPost mocks base method
func (m *MockPostsRepository) GetLikeCountByPost(ctx context.Context, postID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikeCountByPost", ctx, postID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeCountByPost indicates an expected call of GetLikeCountByPost
func (mr *MockPostsRepositoryMockRecorder) GetLikeCountByPost(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeCountByPost", reflect.TypeOf((*MockPostsRepository)(nil).GetLikeCountByPost), ctx, postID)
}

// AppendOrRemoveUserFromLikedBy mocks base method
func (m *MockPostsRepository) AppendOrRemoveUserFromLikedBy(postID, userID string, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendOrRemoveUserFromLikedBy", postID, userID, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendOrRemoveUserFromLikedBy indicates an expected call of AppendOrRemoveUserFromLikedBy
func (mr *MockPostsRepositoryMockRecorder) AppendOrRemoveUserFromLikedBy(postID, userID, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendOrRemoveUserFromLikedBy", reflect.TypeOf((*MockPostsRepository)(nil).AppendOrRemoveUserFromLikedBy), postID, userID, ctx)
}

// SaveInitialLike mocks base method
func (m *MockPostsRepository) SaveInitialLike(ctx context.Context, postID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInitialLike", ctx, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveInitialLike indicates an expected call of SaveInitialLike
func (mr *MockPostsRepositoryMockRecorder) SaveInitialLike(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInitialLike", reflect.TypeOf((*MockPostsRepository)(nil).SaveInitialLike), ctx, postID)
}

// GetPostID mocks base method
func (m *MockPostsRepository) GetPostID(ctx context.Context, postUUID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostID", ctx, postUUID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostID indicates an expected call of GetPostID
func (mr *MockPostsRepositoryMockRecorder) GetPostID(ctx, postUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostID", reflect.TypeOf((*MockPostsRepository)(nil).GetPostID), ctx, postUUID)
}
