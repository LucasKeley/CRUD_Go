package controller

import (
	"net/http"

	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/validation"
	"github.com/LucasKeley/CRUD_Go/src/controller/model/request"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controler",
		zap.String("journey", "CreateUser"))
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error Trying to validate user info", err,
			zap.String("journey", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.PassWord,
		UserRequest.Name,
		UserRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")

}
