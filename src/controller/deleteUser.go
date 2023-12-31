package controller

import (
	"net/http"

	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "update"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalida userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error(
			"Error trying to call DeleteUser service",
			err,
			zap.String("journey", "DeleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("jouney", "createUser"))

	c.Status(http.StatusOK)

}
