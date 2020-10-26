// Code generated by MockGen. DO NOT EDIT.
// Source: draft_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "post-api/models"
	db "post-api/models/db"
	request "post-api/models/request"
	reflect "reflect"
)

// MockDraftRepository is a mock of DraftRepository interface
type MockDraftRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDraftRepositoryMockRecorder
}

// MockDraftRepositoryMockRecorder is the mock recorder for MockDraftRepository
type MockDraftRepositoryMockRecorder struct {
	mock *MockDraftRepository
}

// NewMockDraftRepository creates a new mock instance
func NewMockDraftRepository(ctrl *gomock.Controller) *MockDraftRepository {
	mock := &MockDraftRepository{ctrl: ctrl}
	mock.recorder = &MockDraftRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDraftRepository) EXPECT() *MockDraftRepositoryMockRecorder {
	return m.recorder
}

// SavePostDraft mocks base method
func (m *MockDraftRepository) SavePostDraft(draft models.UpsertDraft, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePostDraft", draft, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePostDraft indicates an expected call of SavePostDraft
func (mr *MockDraftRepositoryMockRecorder) SavePostDraft(draft, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePostDraft", reflect.TypeOf((*MockDraftRepository)(nil).SavePostDraft), draft, ctx)
}

// SaveTitleDraft mocks base method
func (m *MockDraftRepository) SaveTitleDraft(draft models.UpsertDraft, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTitleDraft", draft, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTitleDraft indicates an expected call of SaveTitleDraft
func (mr *MockDraftRepositoryMockRecorder) SaveTitleDraft(draft, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTitleDraft", reflect.TypeOf((*MockDraftRepository)(nil).SaveTitleDraft), draft, ctx)
}

// SaveTaglineToDraft mocks base method
func (m *MockDraftRepository) SaveTaglineToDraft(taglineSaveRequest request.TaglineSaveRequest, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTaglineToDraft", taglineSaveRequest, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTaglineToDraft indicates an expected call of SaveTaglineToDraft
func (mr *MockDraftRepositoryMockRecorder) SaveTaglineToDraft(taglineSaveRequest, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTaglineToDraft", reflect.TypeOf((*MockDraftRepository)(nil).SaveTaglineToDraft), taglineSaveRequest, ctx)
}

// GetDraft mocks base method
func (m *MockDraftRepository) GetDraft(ctx context.Context, draftUID string) (db.Draft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDraft", ctx, draftUID)
	ret0, _ := ret[0].(db.Draft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDraft indicates an expected call of GetDraft
func (mr *MockDraftRepositoryMockRecorder) GetDraft(ctx, draftUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraft", reflect.TypeOf((*MockDraftRepository)(nil).GetDraft), ctx, draftUID)
}
