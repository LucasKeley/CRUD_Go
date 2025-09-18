package mocks

import (
	reflect "reflect"

	rest_err "github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	gomock "go.uber.org/mock/gomock"
)

type MockUserDomainInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainInterfaceMockRecorder
}

type MockUserDomainInterfaceMockRecorder struct {
	mock *MockUserDomainInterface
}

func (mr *MockUserDomainInterfaceMockRecorder) DeleteUser(id string) {
	panic("unimplemented")
}

func NewMockUserDomainInterface(ctrl *gomock.Controller) *MockUserDomainInterface {
	mock := &MockUserDomainInterface{ctrl: ctrl}
	mock.recorder = &MockUserDomainInterfaceMockRecorder{mock}
	return mock
}

func (m *MockUserDomainInterface) EXPECT() *MockUserDomainInterfaceMockRecorder {
	return m.recorder
}

func (m *MockUserDomainInterface) EncryptPassword() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EncryptPassword")
}

func (mr *MockUserDomainInterfaceMockRecorder) EncryptPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).EncryptPassword))
}

func (m *MockUserDomainInterface) GenerateToken() (string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

func (mr *MockUserDomainInterfaceMockRecorder) GenerateToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUserDomainInterface)(nil).GenerateToken))
}

func (m *MockUserDomainInterface) GetAge() int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAge")
	ret0, _ := ret[0].(int8)
	return ret0
}

func (mr *MockUserDomainInterfaceMockRecorder) GetAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAge", reflect.TypeOf((*MockUserDomainInterface)(nil).GetAge))
}

func (m *MockUserDomainInterface) GetEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *MockUserDomainInterfaceMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockUserDomainInterface)(nil).GetEmail))
}

func (m *MockUserDomainInterface) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *MockUserDomainInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUserDomainInterface)(nil).GetID))
}

func (m *MockUserDomainInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *MockUserDomainInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockUserDomainInterface)(nil).GetName))
}

func (m *MockUserDomainInterface) GetPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *MockUserDomainInterfaceMockRecorder) GetPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).GetPassword))
}

func (m *MockUserDomainInterface) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

func (mr *MockUserDomainInterfaceMockRecorder) SetID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockUserDomainInterface)(nil).SetID), arg0)
}
