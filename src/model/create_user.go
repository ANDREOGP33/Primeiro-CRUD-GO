package model

import (
	"fmt"

	"crypto/md5"
	"encoding/hex"

	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/logger"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	func() {
		hash := md5.New()
		defer hash.Reset()
		hash.Write([]byte(ud.Password))
		ud.Password = hex.EncodeToString(hash.Sum(nil))
	}()
	fmt.Println(ud)
	return nil
}
