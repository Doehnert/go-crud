package main

import (
	"github.com/Doehnert/crud/src/controller"
	"github.com/Doehnert/crud/src/model/repository"
	"github.com/Doehnert/crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	// Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	return userController
}
