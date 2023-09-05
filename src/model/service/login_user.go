package service

import (
	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init login model", zap.String("jouney", "loginUser"))

	userDomain.EncryptPassword()
	user, err := ud.findUserByEmailAndPasswordService(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)

	return user, token, nil
}
