// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/dependency/post.go

// Package post is a generated GoMock package.
package post

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/olongfen/go-ddd-hex/internal/domain/entity"
	query "github.com/olongfen/go-ddd-hex/lib/query"
)

// MockPostRepo is a mock of PostRepo interface.
type MockPostRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepoMockRecorder
}

// MockPostRepoMockRecorder is the mock recorder for MockPostRepo.
type MockPostRepoMockRecorder struct {
	mock *MockPostRepo
}

// NewMockPostRepo creates a new mock instance.
func NewMockPostRepo(ctrl *gomock.Controller) *MockPostRepo {
	mock := &MockPostRepo{ctrl: ctrl}
	mock.recorder = &MockPostRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepo) EXPECT() *MockPostRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostRepo) Create(post *entity.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPostRepoMockRecorder) Create(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostRepo)(nil).Create), post)
}

// Delete mocks base method.
func (m *MockPostRepo) Delete(cond map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", cond)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostRepoMockRecorder) Delete(cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostRepo)(nil).Delete), cond)
}

// Find mocks base method.
func (m *MockPostRepo) Find(cond map[string]interface{}, meta *query.Meta) ([]*entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", cond, meta)
	ret0, _ := ret[0].([]*entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPostRepoMockRecorder) Find(cond, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPostRepo)(nil).Find), cond, meta)
}

// Get mocks base method.
func (m *MockPostRepo) Get(id string) (*entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostRepoMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostRepo)(nil).Get), id)
}

// Update mocks base method.
func (m *MockPostRepo) Update(cond map[string]interface{}, change interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", cond, change)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostRepoMockRecorder) Update(cond, change interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostRepo)(nil).Update), cond, change)
}
