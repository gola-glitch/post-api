// Code generated by MockGen. DO NOT EDIT.
// Source: posts_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	neo4j "github.com/neo4j/neo4j-go-driver/neo4j"
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
func (m *MockPostsRepository) CreatePost(ctx context.Context, post db.PublishPost, transaction neo4j.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, post, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockPostsRepositoryMockRecorder) CreatePost(ctx, post, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostsRepository)(nil).CreatePost), ctx, post, transaction)
}

// LikePost mocks base method
func (m *MockPostsRepository) LikePost(postID, userID string, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikePost", postID, userID, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// LikePost indicates an expected call of LikePost
func (mr *MockPostsRepositoryMockRecorder) LikePost(postID, userID, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikePost", reflect.TypeOf((*MockPostsRepository)(nil).LikePost), postID, userID, ctx)
}

// UnlikePost mocks base method
func (m *MockPostsRepository) UnlikePost(ctx context.Context, userId, postId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlikePost", ctx, userId, postId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlikePost indicates an expected call of UnlikePost
func (mr *MockPostsRepositoryMockRecorder) UnlikePost(ctx, userId, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlikePost", reflect.TypeOf((*MockPostsRepository)(nil).UnlikePost), ctx, userId, postId)
}

// IsPostLikedByPerson mocks base method
func (m *MockPostsRepository) IsPostLikedByPerson(ctx context.Context, userId, postId string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPostLikedByPerson", ctx, userId, postId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPostLikedByPerson indicates an expected call of IsPostLikedByPerson
func (mr *MockPostsRepositoryMockRecorder) IsPostLikedByPerson(ctx, userId, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPostLikedByPerson", reflect.TypeOf((*MockPostsRepository)(nil).IsPostLikedByPerson), ctx, userId, postId)
}

// CommentPost mocks base method
func (m *MockPostsRepository) CommentPost(ctx context.Context, userId, comment, postId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommentPost", ctx, userId, comment, postId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommentPost indicates an expected call of CommentPost
func (mr *MockPostsRepositoryMockRecorder) CommentPost(ctx, userId, comment, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentPost", reflect.TypeOf((*MockPostsRepository)(nil).CommentPost), ctx, userId, comment, postId)
}

// GetLikesCountByPostID mocks base method
func (m *MockPostsRepository) GetLikesCountByPostID(ctx context.Context, postId string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikesCountByPostID", ctx, postId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikesCountByPostID indicates an expected call of GetLikesCountByPostID
func (mr *MockPostsRepositoryMockRecorder) GetLikesCountByPostID(ctx, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikesCountByPostID", reflect.TypeOf((*MockPostsRepository)(nil).GetLikesCountByPostID), ctx, postId)
}
