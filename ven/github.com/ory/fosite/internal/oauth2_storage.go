// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: CoreStorage)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of CoreStorage interface
type MockCoreStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockCoreStorageRecorder
}

// Recorder for MockCoreStorage (not exported)
type _MockCoreStorageRecorder struct {
	mock *MockCoreStorage
}

func NewMockCoreStorage(ctrl *gomock.Controller) *MockCoreStorage {
	mock := &MockCoreStorage{ctrl: ctrl}
	mock.recorder = &_MockCoreStorageRecorder{mock}
	return mock
}

func (_m *MockCoreStorage) EXPECT() *_MockCoreStorageRecorder {
	return _m.recorder
}

func (_m *MockCoreStorage) CreateAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStorageRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockCoreStorage) CreateAuthorizeCodeSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStorageRecorder) CreateAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockCoreStorage) CreateRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStorageRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockCoreStorage) DeleteAccessTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAccessTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStorageRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAccessTokenSession", arg0, arg1)
}

func (_m *MockCoreStorage) DeleteAuthorizeCodeSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAuthorizeCodeSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStorageRecorder) DeleteAuthorizeCodeSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAuthorizeCodeSession", arg0, arg1)
}

func (_m *MockCoreStorage) DeleteRefreshTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRefreshTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStorageRecorder) DeleteRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRefreshTokenSession", arg0, arg1)
}

func (_m *MockCoreStorage) GetAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCoreStorageRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockCoreStorage) GetAuthorizeCodeSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCoreStorageRecorder) GetAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockCoreStorage) GetRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCoreStorageRecorder) GetRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRefreshTokenSession", arg0, arg1, arg2)
}
