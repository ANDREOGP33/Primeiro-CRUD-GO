package controller

import (
	"fmt"

	rest_err "github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect filds, error=%s\n", err.Error))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
