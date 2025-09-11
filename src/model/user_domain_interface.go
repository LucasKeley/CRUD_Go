package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int
	GetName() string

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
