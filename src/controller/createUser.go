package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/validation"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/request"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errReset := validation.ValidateUserError(err)

		c.JSON(errReset.Code, errReset)
		return
	}
	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}
