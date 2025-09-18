package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"go.uber.org/zap"
)

// FindUserByIDServices orquestra a busca por ID garantindo logs consistentes.
func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID services", zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByID(id)
}

// FindUserByEmailServices centraliza a leitura por e-mail dentro da camada de serviços.
func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail services", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

// FindUserByEmailAndPasswordServices suporta a autenticação reutilizando o repositório.
func (ud *userDomainService) FindUserByEmailAndPasswordServices(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail services", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
