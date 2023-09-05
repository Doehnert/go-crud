package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Doehnert/crud/src/configuration/rest_err"
	"github.com/Doehnert/crud/src/model"
	"github.com/Doehnert/crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_FindUserBYIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repo)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {

		id := primitive.NewObjectID().Hex()
		user := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)
		repo.EXPECT().FindUserByID(id).Return(user, nil)

		userDomain, err := service.FindUserByIDService(id)

		assert.Nil(t, err)

		assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), user.GetName())
		assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {

		id := primitive.NewObjectID().Hex()

		repo.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomain, err := service.FindUserByIDService(id)
		assert.NotNil(t, err)

		assert.Nil(t, userDomain)
		assert.EqualValues(t, err.Message, "user not found")
	})

}
func TestUserDomainService_FindUserBYEmailService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repo)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {

		user := model.NewUserDomain(
			"test@test.com",
			"teste",
			"teste",
			33)
		repo.EXPECT().FindUserByEmail("test@test.com").Return(user, nil)

		userDomain, err := service.FindUserByEmailService("test@test.com")

		assert.Nil(t, err)

		assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), user.GetName())
		assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {

		repo.EXPECT().FindUserByEmail("test@test.com").Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomain, err := service.FindUserByEmailService("test@test.com")
		assert.NotNil(t, err)

		assert.Nil(t, userDomain)
		assert.EqualValues(t, err.Message, "user not found")
	})

}
func TestUserDomainService_FindUserBYEmailAndPasswordService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{
		repo,
	}

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		user := model.NewUserDomain(
			email,
			password,
			"teste",
			33)
		user.SetID(id)
		repo.EXPECT().FindUserByEmailAndPassword(email, password).Return(user, nil)

		userDomain, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, err)

		assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), user.GetName())
		assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		repo.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomain, err := service.findUserByEmailAndPasswordService(email, password)
		assert.NotNil(t, err)

		assert.Nil(t, userDomain)
		assert.EqualValues(t, err.Message, "user not found")
	})

}
