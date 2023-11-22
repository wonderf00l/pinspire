// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	user "github.com/go-park-mail-ru/2023_2_OND_team/internal/pkg/entity/user"
	user0 "github.com/go-park-mail-ru/2023_2_OND_team/internal/pkg/usecase/user"
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

// Authentication mocks base method.
func (m *MockUsecase) Authentication(ctx context.Context, credentials user0.UserCredentials) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx, credentials)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authentication indicates an expected call of Authentication.
func (mr *MockUsecaseMockRecorder) Authentication(ctx, credentials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockUsecase)(nil).Authentication), ctx, credentials)
}

// EditProfileInfo mocks base method.
func (m *MockUsecase) EditProfileInfo(ctx context.Context, userID int, updateData *user0.ProfileUpdateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfileInfo", ctx, userID, updateData)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditProfileInfo indicates an expected call of EditProfileInfo.
func (mr *MockUsecaseMockRecorder) EditProfileInfo(ctx, userID, updateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfileInfo", reflect.TypeOf((*MockUsecase)(nil).EditProfileInfo), ctx, userID, updateData)
}

// FindOutUsernameAndAvatar mocks base method.
func (m *MockUsecase) FindOutUsernameAndAvatar(ctx context.Context, userID int) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOutUsernameAndAvatar", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindOutUsernameAndAvatar indicates an expected call of FindOutUsernameAndAvatar.
func (mr *MockUsecaseMockRecorder) FindOutUsernameAndAvatar(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOutUsernameAndAvatar", reflect.TypeOf((*MockUsecase)(nil).FindOutUsernameAndAvatar), ctx, userID)
}

// GetAllProfileInfo mocks base method.
func (m *MockUsecase) GetAllProfileInfo(ctx context.Context, userID int) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProfileInfo", ctx, userID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProfileInfo indicates an expected call of GetAllProfileInfo.
func (mr *MockUsecaseMockRecorder) GetAllProfileInfo(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProfileInfo", reflect.TypeOf((*MockUsecase)(nil).GetAllProfileInfo), ctx, userID)
}

// GetProfileInfo mocks base method.
func (m *MockUsecase) GetProfileInfo(ctx context.Context) (*user.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileInfo", ctx)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProfileInfo indicates an expected call of GetProfileInfo.
func (mr *MockUsecaseMockRecorder) GetProfileInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileInfo", reflect.TypeOf((*MockUsecase)(nil).GetProfileInfo), ctx)
}

// GetUserInfo mocks base method.
func (m *MockUsecase) GetUserInfo(ctx context.Context, userID int) (*user.User, bool, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, userID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUsecaseMockRecorder) GetUserInfo(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUsecase)(nil).GetUserInfo), ctx, userID)
}

// Register mocks base method.
func (m *MockUsecase) Register(ctx context.Context, user *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUsecaseMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUsecase)(nil).Register), ctx, user)
}

// UpdateUserAvatar mocks base method.
func (m *MockUsecase) UpdateUserAvatar(ctx context.Context, userID int, mimeTypeAvatar string, sizeAvatar int64, avatar io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAvatar", ctx, userID, mimeTypeAvatar, sizeAvatar, avatar)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAvatar indicates an expected call of UpdateUserAvatar.
func (mr *MockUsecaseMockRecorder) UpdateUserAvatar(ctx, userID, mimeTypeAvatar, sizeAvatar, avatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAvatar", reflect.TypeOf((*MockUsecase)(nil).UpdateUserAvatar), ctx, userID, mimeTypeAvatar, sizeAvatar, avatar)
}
