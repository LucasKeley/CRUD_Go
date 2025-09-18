package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"go.uber.org/zap"
)

// CreateUserServeces garante unicidade de e-mail, aplica hash de senha e persiste o usuário.
func (ud *userDomainService) CreateUserServeces(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	userDomanRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Info("Init createUser model", zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomanRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomanRepository, nil
}
