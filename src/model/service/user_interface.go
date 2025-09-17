package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/LucasKeley/CRUD_Go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email string, password string) (any, any) {
	panic("unimplemented")
}

func (ud *userDomainService) CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	panic("unimplemented")
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
