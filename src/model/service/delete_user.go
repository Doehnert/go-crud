package service

import (
	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	logger.Info("delete user", zap.String("jouney", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Info("Error trying to call deleteUser repo",
			zap.String("journey", "createUser"))
		return err
	}

	logger.Info(
		"deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}
