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

// IsDraftPresent mocks base method
func (m *MockDraftRepository) IsDraftPresent(ctx context.Context, draftId, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDraftPresent", ctx, draftId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsDraftPresent indicates an expected call of IsDraftPresent
func (mr *MockDraftRepositoryMockRecorder) IsDraftPresent(ctx, draftId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDraftPresent", reflect.TypeOf((*MockDraftRepository)(nil).IsDraftPresent), ctx, draftId, userId)
}

// CreateNewPostWithData mocks base method
func (m *MockDraftRepository) CreateNewPostWithData(draft models.UpsertDraft, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewPostWithData", draft, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewPostWithData indicates an expected call of CreateNewPostWithData
func (mr *MockDraftRepositoryMockRecorder) CreateNewPostWithData(draft, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewPostWithData", reflect.TypeOf((*MockDraftRepository)(nil).CreateNewPostWithData), draft, ctx)
}

// UpdateDraft mocks base method
func (m *MockDraftRepository) UpdateDraft(draft models.UpsertDraft, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDraft", draft, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDraft indicates an expected call of UpdateDraft
func (mr *MockDraftRepositoryMockRecorder) UpdateDraft(draft, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDraft", reflect.TypeOf((*MockDraftRepository)(nil).UpdateDraft), draft, ctx)
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

// SaveInterestsToDraft mocks base method
func (m *MockDraftRepository) SaveInterestsToDraft(interestsSaveRequest request.InterestsSaveRequest, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInterestsToDraft", interestsSaveRequest, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveInterestsToDraft indicates an expected call of SaveInterestsToDraft
func (mr *MockDraftRepositoryMockRecorder) SaveInterestsToDraft(interestsSaveRequest, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInterestsToDraft", reflect.TypeOf((*MockDraftRepository)(nil).SaveInterestsToDraft), interestsSaveRequest, ctx)
}

// GetDraft mocks base method
func (m *MockDraftRepository) GetDraft(ctx context.Context, draftUID, userId string) (db.DraftDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDraft", ctx, draftUID, userId)
	ret0, _ := ret[0].(db.DraftDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDraft indicates an expected call of GetDraft
func (mr *MockDraftRepositoryMockRecorder) GetDraft(ctx, draftUID, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraft", reflect.TypeOf((*MockDraftRepository)(nil).GetDraft), ctx, draftUID, userId)
}

// UpsertPreviewImage mocks base method
func (m *MockDraftRepository) UpsertPreviewImage(ctx context.Context, saveRequest request.PreviewImageSaveRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPreviewImage", ctx, saveRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPreviewImage indicates an expected call of UpsertPreviewImage
func (mr *MockDraftRepositoryMockRecorder) UpsertPreviewImage(ctx, saveRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPreviewImage", reflect.TypeOf((*MockDraftRepository)(nil).UpsertPreviewImage), ctx, saveRequest)
}

// DeleteDraft mocks base method
func (m *MockDraftRepository) DeleteDraft(ctx context.Context, draftUID, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDraft", ctx, draftUID, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDraft indicates an expected call of DeleteDraft
func (mr *MockDraftRepositoryMockRecorder) DeleteDraft(ctx, draftUID, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDraft", reflect.TypeOf((*MockDraftRepository)(nil).DeleteDraft), ctx, draftUID, userId)
}

// GetAllDraft mocks base method
func (m *MockDraftRepository) GetAllDraft(ctx context.Context, allDraftReq models.GetAllDraftRequest) ([]db.DraftDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDraft", ctx, allDraftReq)
	ret0, _ := ret[0].([]db.DraftDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDraft indicates an expected call of GetAllDraft
func (mr *MockDraftRepositoryMockRecorder) GetAllDraft(ctx, allDraftReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDraft", reflect.TypeOf((*MockDraftRepository)(nil).GetAllDraft), ctx, allDraftReq)
}
