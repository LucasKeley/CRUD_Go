package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int
	GetName() string
	GetID() string
	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
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
