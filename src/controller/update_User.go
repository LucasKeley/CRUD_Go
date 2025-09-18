package controller

import (
	"net/http"

	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/configuration/validation"
	"github.com/LucasKeley/CRUD_Go/src/controller/model/request"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// UpdateUser valida a requisição de atualização, reforçando o formato do ID antes de acionar a camada de serviço.
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controler",
		zap.String("journey", "UpdateUser"))

	var UserRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error Trying to validate user info", err,
			zap.String("journey", "UpdateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		UserRequest.Name,
		UserRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Info("Error trying to call UpdateUser service", zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"))

	c.Status(http.StatusOK)

}
