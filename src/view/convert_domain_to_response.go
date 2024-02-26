package view

import (
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/model/response"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
