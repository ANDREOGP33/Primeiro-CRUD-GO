package service

import (
	"fmt"

	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/logger"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/rest_err"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(
		userDomain.GetPassword(),
		userDomain.GetName(),
		userDomain.GetEmail(),
		userDomain.GetAge(),
	)

	return nil
}
