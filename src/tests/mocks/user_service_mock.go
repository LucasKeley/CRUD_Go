package mocks

import (
	reflect "reflect"

	rest_err "github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	model "github.com/LucasKeley/CRUD_Go/src/model"
	gomock "github.com/golang/mock/gomock"
)

type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
}

type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

func (m *MockUserDomainService) CreateUserServices(arg0 model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserServices", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserDomainServiceMockRecorder) CreateUserServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserServices", reflect.TypeOf((*MockUserDomainService)(nil).CreateUserServices), arg0)
}

func (m *MockUserDomainService) DeleteUser(arg0 string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

func (mr *MockUserDomainServiceMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserDomainService)(nil).DeleteUser), arg0)
}

func (m *MockUserDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailServices", email)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserDomainServiceMockRecorder) FindUserByEmailServices(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailServices", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByEmailServices), email)
}

func (m *MockUserDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByIDServices", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserDomainServiceMockRecorder) FindUserByIDServices(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByIDServices", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByIDServices), id)
}

func (m *MockUserDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUserServices", userDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rest_err.RestErr)
	return ret0, ret1, ret2
}

func (mr *MockUserDomainServiceMockRecorder) LoginUserServices(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUserServices", reflect.TypeOf((*MockUserDomainService)(nil).LoginUserServices), userDomain)
}

func (m *MockUserDomainService) UpdateUser(arg0 string, arg1 model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

func (mr *MockUserDomainServiceMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserDomainService)(nil).UpdateUser), arg0, arg1)
}
