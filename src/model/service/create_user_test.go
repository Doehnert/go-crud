package service

import (
	"testing"

	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"github.com/Doehnert/crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repo)

	t.Run("when_user_exists_in_database", func(t *testing.T) {

		user := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)

		repo.EXPECT().FindUserByEmail(user.GetEmail()).Return(user, nil)

		user, err := service.CreateUserService(user)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")

	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {

		user := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)

		repo.EXPECT().FindUserByEmail(user.GetEmail()).Return(nil, nil)

		repo.EXPECT().CreateUser(user).Return(nil, rest_err.NewInternalServerError(
			"error trying to create user",
		))

		user, err := service.CreateUserService(user)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create user")

	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {

		userDomain := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repo.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetName(), user.GetName())
		assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
		assert.EqualValues(t, userDomain.GetID(), user.GetID())
		assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())
	})
}
