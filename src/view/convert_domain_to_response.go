package view

import (
	"github.com/LucasKeley/CRUD_Go/src/controller/model/response"
	"github.com/LucasKeley/CRUD_Go/src/model"
)

func ConverteDomaintoResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
