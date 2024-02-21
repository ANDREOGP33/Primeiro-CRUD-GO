package controller

import (
	"net/http"

	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/logger"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/validation"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/request"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		zap.Field{
			Key:    "journey",
			String: "createUser",
		},
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.Field{
			Key:    "journey",
			String: "createUser",
		})
		errReset := validation.ValidateUserError(err)

		c.JSON(errReset.Code, errReset)
		return
	}
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully", zap.Field{
		Key:    "journey",
		String: "createUser",
	})
	c.JSON(http.StatusOK, response)
}
