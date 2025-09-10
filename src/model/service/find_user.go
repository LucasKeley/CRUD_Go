package service

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
