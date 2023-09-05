package service

import (
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"github.com/Doehnert/crud/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUserService(string) *rest_err.RestErr
	FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	findUserByEmailAndPasswordService(email string, password string) (model.UserDomainInterface, *rest_err.RestErr)
	LoginUserService(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_err.RestErr)
}
