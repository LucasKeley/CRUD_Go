package model

import "github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int
	GetName() string
	GetID() string
	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}

}

func NewUserUpdateDomain(
	name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}

}

func NewUserLoginDomain(
	email, password string,

) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}

}
