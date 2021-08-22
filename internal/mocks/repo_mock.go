// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-certificate-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ozoncp/ocp-certificate-api/internal/model"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddCertificates mocks base method.
func (m *MockRepo) AddCertificates(arg0 context.Context, arg1 []model.Certificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCertificates", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCertificates indicates an expected call of AddCertificates.
func (mr *MockRepoMockRecorder) AddCertificates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCertificates", reflect.TypeOf((*MockRepo)(nil).AddCertificates), arg0, arg1)
}

// CreateCertificate mocks base method.
func (m *MockRepo) CreateCertificate(arg0 context.Context, arg1 *model.Certificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCertificate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCertificate indicates an expected call of CreateCertificate.
func (mr *MockRepoMockRecorder) CreateCertificate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificate", reflect.TypeOf((*MockRepo)(nil).CreateCertificate), arg0, arg1)
}

// GetCertificate mocks base method.
func (m *MockRepo) GetCertificate(arg0 context.Context, arg1 uint64) (*model.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate", arg0, arg1)
	ret0, _ := ret[0].(*model.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate.
func (mr *MockRepoMockRecorder) GetCertificate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockRepo)(nil).GetCertificate), arg0, arg1)
}

// ListCertificates mocks base method.
func (m *MockRepo) ListCertificates(arg0 context.Context, arg1, arg2 uint64) ([]model.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificates", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificates indicates an expected call of ListCertificates.
func (mr *MockRepoMockRecorder) ListCertificates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockRepo)(nil).ListCertificates), arg0, arg1, arg2)
}

// RemoveCertificate mocks base method.
func (m *MockRepo) RemoveCertificate(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCertificate", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveCertificate indicates an expected call of RemoveCertificate.
func (mr *MockRepoMockRecorder) RemoveCertificate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCertificate", reflect.TypeOf((*MockRepo)(nil).RemoveCertificate), arg0, arg1)
}

// UpdateCertificate mocks base method.
func (m *MockRepo) UpdateCertificate(arg0 context.Context, arg1 model.Certificate) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCertificate", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCertificate indicates an expected call of UpdateCertificate.
func (mr *MockRepoMockRecorder) UpdateCertificate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificate", reflect.TypeOf((*MockRepo)(nil).UpdateCertificate), arg0, arg1)
}
