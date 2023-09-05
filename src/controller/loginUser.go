package controller

import (
	"net/http"

	"github.com/Doehnert/crud/src/configuration/logger"
	"github.com/Doehnert/crud/src/configuration/rest_err/validation"
	"github.com/Doehnert/crud/src/controller/model/request"
	"github.com/Doehnert/crud/src/model"
	"github.com/Doehnert/crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller")
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error(
			"Error trying to call LoginUser service",
			err,
			zap.String("journey", "LoginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("jouney", "LoginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
