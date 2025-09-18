package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"go.uber.org/zap"
)

// LoginUserServices autentica o usuário e retorna um token JWT pronto para uso em rotas protegidas.
func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {

	user, err := ud.FindUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, "", err
	}

	logger.Info("Init loginUser model", zap.String("journey", "loginUser"))
	userDomain.EncryptPassword()

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))

	return user, token, nil
}
