package mocks

import (
	reflect "reflect"

	rest_err "github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	model "github.com/LucasKeley/CRUD_Go/src/model"
	gomock "go.uber.org/mock/gomock"
)

type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

func (m *MockUserRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) CreateUser(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), userDomain)
}

func (m *MockUserRepository) DeleteUser(userId string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

func (mr *MockUserRepositoryMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), userId)
}

func (m *MockUserRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), email)
}

func (m *MockUserRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailAndPassword", email, password)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) FindUserByEmailAndPassword(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailAndPassword", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmailAndPassword), email, password)
}

func (m *MockUserRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) FindUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByID), id)
}

func (m *MockUserRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, userDomain)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

func (mr *MockUserRepositoryMockRecorder) UpdateUser(userId, userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), userId, userDomain)
}
