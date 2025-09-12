package controller

import (
	"net/http"
	"net/mail"

	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "FindUserByID"))

	userId := c.Param(("userId"))

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userID controller", err,
			zap.String("journey", "FindUserByID"))
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to validate userID services", err,
			zap.String("journey", "FindUserByID"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConverteDomaintoResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "FindUserByEmail"))

	userEmail := c.Param(("userEmail"))

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail services", err,
			zap.String("journey", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("findUserByEmail controler execudet successfully",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConverteDomaintoResponse(userDomain))

}
