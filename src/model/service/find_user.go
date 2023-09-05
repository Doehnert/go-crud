package service

import (
	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDService(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID service", zap.String("jouney", "findUserByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("find user", zap.String("jouney", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("find user", zap.String("jouney", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
