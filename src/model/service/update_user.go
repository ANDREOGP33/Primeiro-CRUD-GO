package service

import (
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/rest_err"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
