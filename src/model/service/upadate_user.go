package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))
		return err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"))

	return nil
}
