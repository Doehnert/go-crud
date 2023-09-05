package service

import (
	"fmt"

	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("jouney", "createUser"))

	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Info("Error trying to call CreateUser repo",
			zap.String("journey", "createUser"))
		return nil, err
	}

	fmt.Println(userDomain.GetPassword())

	return userDomainRepository, nil
}
