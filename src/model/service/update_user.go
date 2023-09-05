package service

import (
	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("update user", zap.String("jouney", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Info("Error trying to call UpdateUser repo",
			zap.String("journey", "createUser"))
		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
