package service

import (
	"testing"

	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"github.com/Doehnert/crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_UpdateUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repo)

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)
		userDomain.SetID(id)

		repo.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUserService(id, userDomain)

		assert.Nil(t, err)
	})
	t.Run("when_sending_a_invalid_user_and_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)
		userDomain.SetID(id)

		repo.EXPECT().UpdateUser(id, userDomain).Return(rest_err.NewInternalServerError(
			"error trying to update user",
		))

		err := service.UpdateUserService(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
