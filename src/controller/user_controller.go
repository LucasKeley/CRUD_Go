package controller

import (
	"github.com/LucasKeley/CRUD_Go/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewUserControllerInterface injeta a camada de serviço e entrega uma implementação pronta do controlador HTTP.
func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

// UserControllerInterface define os handlers expostos à camada de rotas do Gin.
type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
