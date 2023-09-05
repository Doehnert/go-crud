package repository

import (
	"github.com/Doehnert/crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"os"
	"testing"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_user_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90,
		)

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)

		domain.SetID(primitive.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetID(), domain)
		// test if generated id is valid

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90,
		)

		repo := NewUserRepository(databaseMock)

		err := repo.UpdateUser(domain.GetID(), domain)

		assert.NotNil(t, err)

	})
}
