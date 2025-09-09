package controller

import (
	"fmt"

	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrent filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

}
