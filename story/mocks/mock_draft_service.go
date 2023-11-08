// Code generated by MockGen. DO NOT EDIT.
// Source: draft_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "post-api/story/models"
	db "post-api/story/models/db"
	request "post-api/story/models/request"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	golaerror "github.com/inclusi-blog/gola-utils/golaerror"
)

// MockDraftService is a mock of DraftService interface.
type MockDraftService struct {
	ctrl     *gomock.Controller
	recorder *MockDraftServiceMockRecorder
}

// MockDraftServiceMockRecorder is the mock recorder for MockDraftService.
type MockDraftServiceMockRecorder struct {
	mock *MockDraftService
}

// NewMockDraftService creates a new mock instance.
func NewMockDraftService(ctrl *gomock.Controller) *MockDraftService {
	mock := &MockDraftService{ctrl: ctrl}
	mock.recorder = &MockDraftServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDraftService) EXPECT() *MockDraftServiceMockRecorder {
	return m.recorder
}

// CreateDraft mocks base method.
func (m *MockDraftService) CreateDraft(ctx context.Context, draft models.CreateDraft) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDraft", ctx, draft)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDraft indicates an expected call of CreateDraft.
func (mr *MockDraftServiceMockRecorder) CreateDraft(ctx, draft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDraft", reflect.TypeOf((*MockDraftService)(nil).CreateDraft), ctx, draft)
}

// DeleteDraft mocks base method.
func (m *MockDraftService) DeleteDraft(ctx context.Context, draftID, userUUID uuid.UUID) *golaerror.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDraft", ctx, draftID, userUUID)
	ret0, _ := ret[0].(*golaerror.Error)
	return ret0
}

// DeleteDraft indicates an expected call of DeleteDraft.
func (mr *MockDraftServiceMockRecorder) DeleteDraft(ctx, draftID, userUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDraft", reflect.TypeOf((*MockDraftService)(nil).DeleteDraft), ctx, draftID, userUUID)
}

// GetAllDraft mocks base method.
func (m *MockDraftService) GetAllDraft(ctx context.Context, allDraftReq models.GetAllDraftRequest) ([]db.DraftPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDraft", ctx, allDraftReq)
	ret0, _ := ret[0].([]db.DraftPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDraft indicates an expected call of GetAllDraft.
func (mr *MockDraftServiceMockRecorder) GetAllDraft(ctx, allDraftReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDraft", reflect.TypeOf((*MockDraftService)(nil).GetAllDraft), ctx, allDraftReq)
}

// GetDraft mocks base method.
func (m *MockDraftService) GetDraft(ctx context.Context, draftUID, userUUID uuid.UUID) (db.Draft, *golaerror.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDraft", ctx, draftUID, userUUID)
	ret0, _ := ret[0].(db.Draft)
	ret1, _ := ret[1].(*golaerror.Error)
	return ret0, ret1
}

// GetDraft indicates an expected call of GetDraft.
func (mr *MockDraftServiceMockRecorder) GetDraft(ctx, draftUID, userUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraft", reflect.TypeOf((*MockDraftService)(nil).GetDraft), ctx, draftUID, userUUID)
}

// SavePreviewImage mocks base method.
func (m *MockDraftService) SavePreviewImage(ctx context.Context, imageSaveRequest request.PreviewImageSaveRequest) *golaerror.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreviewImage", ctx, imageSaveRequest)
	ret0, _ := ret[0].(*golaerror.Error)
	return ret0
}

// SavePreviewImage indicates an expected call of SavePreviewImage.
func (mr *MockDraftServiceMockRecorder) SavePreviewImage(ctx, imageSaveRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreviewImage", reflect.TypeOf((*MockDraftService)(nil).SavePreviewImage), ctx, imageSaveRequest)
}

// UpdateDraft mocks base method.
func (m *MockDraftService) UpdateDraft(postData models.UpsertDraft, ctx context.Context) *golaerror.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDraft", postData, ctx)
	ret0, _ := ret[0].(*golaerror.Error)
	return ret0
}

// UpdateDraft indicates an expected call of UpdateDraft.
func (mr *MockDraftServiceMockRecorder) UpdateDraft(postData, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDraft", reflect.TypeOf((*MockDraftService)(nil).UpdateDraft), postData, ctx)
}

// UpsertInterests mocks base method.
func (m *MockDraftService) UpsertInterests(interestRequest request.InterestsSaveRequest, ctx context.Context) *golaerror.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertInterests", interestRequest, ctx)
	ret0, _ := ret[0].(*golaerror.Error)
	return ret0
}

// UpsertInterests indicates an expected call of UpsertInterests.
func (mr *MockDraftServiceMockRecorder) UpsertInterests(interestRequest, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertInterests", reflect.TypeOf((*MockDraftService)(nil).UpsertInterests), interestRequest, ctx)
}

// UpsertTagline mocks base method.
func (m *MockDraftService) UpsertTagline(taglineRequest request.TaglineSaveRequest, ctx context.Context) *golaerror.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTagline", taglineRequest, ctx)
	ret0, _ := ret[0].(*golaerror.Error)
	return ret0
}

// UpsertTagline indicates an expected call of UpsertTagline.
func (mr *MockDraftServiceMockRecorder) UpsertTagline(taglineRequest, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTagline", reflect.TypeOf((*MockDraftService)(nil).UpsertTagline), taglineRequest, ctx)
}
