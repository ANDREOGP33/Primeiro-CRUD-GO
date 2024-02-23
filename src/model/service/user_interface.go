package service

import (
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/rest_err"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/model"
)

func NewUserDomainService() UserDomaService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomaService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeletUser(string) *rest_err.RestErr
}
