package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "UpdateUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))
		return err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
