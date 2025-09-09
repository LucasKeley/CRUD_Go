package controller

import (
	"github.com/LucasKeley/CRUD_Go/src/configuration/validation"
	"github.com/LucasKeley/CRUD_Go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

}
