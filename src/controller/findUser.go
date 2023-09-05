package controller

import (
	"net/http"
	"net/mail"

	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("jouney", "findUserByID"))

	// token := c.Request.Header.Get("Authorization")
	// user, err := model.VerifyToken(token)
	// if err != nil {

	// 	c.JSON(err.Code, err)
	// 	return
	// }

	// logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("jouney", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to call vindUserBYID services",
			err,
			zap.String("jouney", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller called succesfully",
		zap.String("jouney", "findUserByID"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller",
		zap.String("jouney", "FindUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("jouney", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error("Error trying to call vindUserBYID services",
			err,
			zap.String("jouney", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller called succesfully",
		zap.String("jouney", "FindUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
