package converter

import (
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/LucasKeley/CRUD_Go/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID)
	return domain
}
