package controller

import (
	"fmt"
	"net/http"

	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/validation"
	"github.com/LucasKeley/CRUD_Go/src/controller/model/request"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/LucasKeley/CRUD_Go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controler",
		zap.String("journey", "loginUser"))

	var UserRequest request.UserLogin

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenricated: %#v", user))

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error Trying to validate user info", err,
			zap.String("journey", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		UserRequest.Email,
		UserRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Info("Error trying to call CreateUser service", zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "CreateUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConverteDomaintoResponse(
		domainResult,
	))

}
