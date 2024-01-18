// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	comment "github.com/go-park-mail-ru/2023_2_OND_team/internal/pkg/entity/comment"
	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// DeleteComment mocks base method.
func (m *MockUsecase) DeleteComment(ctx context.Context, userID, commentID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, userID, commentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockUsecaseMockRecorder) DeleteComment(ctx, userID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockUsecase)(nil).DeleteComment), ctx, userID, commentID)
}

// GetFeedCommentOnPin mocks base method.
func (m *MockUsecase) GetFeedCommentOnPin(ctx context.Context, userID, pinID, count, lastID int) ([]comment.Comment, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedCommentOnPin", ctx, userID, pinID, count, lastID)
	ret0, _ := ret[0].([]comment.Comment)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFeedCommentOnPin indicates an expected call of GetFeedCommentOnPin.
func (mr *MockUsecaseMockRecorder) GetFeedCommentOnPin(ctx, userID, pinID, count, lastID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedCommentOnPin", reflect.TypeOf((*MockUsecase)(nil).GetFeedCommentOnPin), ctx, userID, pinID, count, lastID)
}

// PutCommentOnPin mocks base method.
func (m *MockUsecase) PutCommentOnPin(ctx context.Context, userID int, comment *comment.Comment) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCommentOnPin", ctx, userID, comment)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutCommentOnPin indicates an expected call of PutCommentOnPin.
func (mr *MockUsecaseMockRecorder) PutCommentOnPin(ctx, userID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCommentOnPin", reflect.TypeOf((*MockUsecase)(nil).PutCommentOnPin), ctx, userID, comment)
}

// MockavailablePinChecker is a mock of availablePinChecker interface.
type MockavailablePinChecker struct {
	ctrl     *gomock.Controller
	recorder *MockavailablePinCheckerMockRecorder
}

// MockavailablePinCheckerMockRecorder is the mock recorder for MockavailablePinChecker.
type MockavailablePinCheckerMockRecorder struct {
	mock *MockavailablePinChecker
}

// NewMockavailablePinChecker creates a new mock instance.
func NewMockavailablePinChecker(ctrl *gomock.Controller) *MockavailablePinChecker {
	mock := &MockavailablePinChecker{ctrl: ctrl}
	mock.recorder = &MockavailablePinCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockavailablePinChecker) EXPECT() *MockavailablePinCheckerMockRecorder {
	return m.recorder
}

// GetAuthorIdOfThePin mocks base method.
func (m *MockavailablePinChecker) GetAuthorIdOfThePin(ctx context.Context, pinID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorIdOfThePin", ctx, pinID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorIdOfThePin indicates an expected call of GetAuthorIdOfThePin.
func (mr *MockavailablePinCheckerMockRecorder) GetAuthorIdOfThePin(ctx, pinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorIdOfThePin", reflect.TypeOf((*MockavailablePinChecker)(nil).GetAuthorIdOfThePin), ctx, pinID)
}

// IsAvailablePinForViewingUser mocks base method.
func (m *MockavailablePinChecker) IsAvailablePinForViewingUser(ctx context.Context, userID, pinID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAvailablePinForViewingUser", ctx, userID, pinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsAvailablePinForViewingUser indicates an expected call of IsAvailablePinForViewingUser.
func (mr *MockavailablePinCheckerMockRecorder) IsAvailablePinForViewingUser(ctx, userID, pinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailablePinForViewingUser", reflect.TypeOf((*MockavailablePinChecker)(nil).IsAvailablePinForViewingUser), ctx, userID, pinID)
}
