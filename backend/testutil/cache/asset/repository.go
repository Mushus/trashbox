// Code generated by MockGen. DO NOT EDIT.
// Source: ../server/app/cache/asset/repository.go

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	asset "github.com/Mushus/trashbox/backend/server/app/asset"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetCache mocks base method
func (m *MockRepository) GetCache(id, format string) (asset.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", id, format)
	ret0, _ := ret[0].(asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache
func (mr *MockRepositoryMockRecorder) GetCache(id, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockRepository)(nil).GetCache), id, format)
}

// PutCache mocks base method
func (m *MockRepository) PutCache(asset asset.Asset, format string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCache", asset, format)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCache indicates an expected call of PutCache
func (mr *MockRepositoryMockRecorder) PutCache(asset, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCache", reflect.TypeOf((*MockRepository)(nil).PutCache), asset, format)
}

// PurgeAll mocks base method
func (m *MockRepository) PurgeAll(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeAll", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeAll indicates an expected call of PurgeAll
func (mr *MockRepositoryMockRecorder) PurgeAll(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeAll", reflect.TypeOf((*MockRepository)(nil).PurgeAll), id)
}