package controller

import (
	"net/http"

	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/logger"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/validation"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/request"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errReset := validation.ValidateUserError(err)

		c.JSON(errReset.Code, errReset)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
