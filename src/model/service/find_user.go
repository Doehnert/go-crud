package service

import (
	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(userId string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("find user", zap.String("jouney", "findUser"))

	return nil, nil
}
