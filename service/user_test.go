package service

import (
	"context"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/stretchr/testify/assert"
	"github/malekradhouane/test-cdi/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

var userCollection *mongo.Collection

func (us *UserService) TestInsertOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	var friend store.Friend
	var friends []store.Friend
	var tags []string
	friend.Id = 1
	friend.Name = "Mike"
	friends = append(friends, friend)
	tags = append(tags, "DataImpact")
	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		id, err := gonanoid.ID(100)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		insertedUser, err := us.userStore.CreateUser(context.Background(),&store.User{
			ID:         id,
			Email:      "malek.radhouen@gmail.com",
			Password:   "0000",
			IsActive:   false,
			Balance:    "$3,897.00",
			Age:        26,
			Name:       "Malek Radhouen",
			Gender:     "male",
			Company:    "Data impact",
			Phone:      "0627867645",
			Address:    "35 rue gallieni",
			About:      "",
			Registered: "2014-07-24T05:08:53 -02:00",
			Latitude:   66.016803,
			Longitude:  -176.499556,
			Friends:    friends,
			Tags:       tags,
			Data:       "test",
		})
		assert.Nil(t, err)
		assert.Equal(t, &store.User{
			ID:         id,
			Email:      "malek.radhouen@gmail.com",
			Password:   "0000",
			IsActive:   false,
			Balance:    "$3,897.00",
			Age:        26,
			Name:       "Malek Radhouen",
			Gender:     "male",
			Company:    "Data impact",
			Phone:      "0627867645",
			Address:    "35 rue gallieni",
			About:      "",
			Registered: "2014-07-24T05:08:53 -02:00",
			Latitude:   66.016803,
			Longitude:  -176.499556,
			Friends:    friends,
			Tags:       tags,
			Data:       "test",
		}, insertedUser)
	})

	mt.Run("custom error duplicate", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		insertedUser, err := mt.Client.Database("user").Collection("users").InsertOne(context.Background(), &store.User{})
		assert.Nil(t, insertedUser)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})

	mt.Run("simple error", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 0}})

		insertedUser, err := mt.Client.Database("user").Collection("users").InsertOne(context.Background(), &store.User{})

		assert.Nil(t, insertedUser)
		assert.NotNil(t, err)
	})
}

func (us *UserService) TestFindOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		expectedUser := store.User{
			ID:         "IrunD0fGdbkRpyLUZcc8aEUKACM8015OhVkSYmVcVKEDVDnjbpbnYJlfu6LlaaYCnvnCWiQsrO2KjhGYCRT1jvOutUhdkG4ufkrx",
			Password:   "3hZRODrkOg0Z2zJQyY2K",
			IsActive:   false,
			Balance:    "$3,897.00",
			Age:        39,
			Name:       "Lola Wolfe",
			Gender:     "female",
			Company:    "ORBIN",
			Email:      "lolawolfe@orbin.com",
			Phone:      "+1 (908) 509-2921",
			Address:    "305 Allen Avenue, Blanford, Virginia, 8797",
			About:      "Aliqua enim elit amet duis ex pariatur commodo fugiat aute. Mollit eu reprehenderit exercitation sunt culpa esse qui et tempor. Cupidatat aute sit Lorem et reprehenderit tempor ut qui aliquip voluptate excepteur sint excepteur sit. Esse laboris eu ad labore ullamco nulla nulla do velit Lorem cillum.\r\n",
			Registered: "2014-07-24T05:08:53 -02:00",
			Latitude:   66.016803,
			Longitude:  -176.499556,
			Tags:       nil,
			Friends:    nil,
			Data:       "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",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "", mtest.FirstBatch, bson.D{
			{"_id", expectedUser.ID},
			{"name", expectedUser.Name},
			{"email", expectedUser.Email},
			{"password", expectedUser.Password},
			{"address", expectedUser.Address},
			{"phone", expectedUser.Phone},
			{"isActive", expectedUser.IsActive},
			{"balance", expectedUser.Balance},
			{"age", expectedUser.Age},
			{"gender", expectedUser.Gender},
			{"company", expectedUser.Company},
			{"about", expectedUser.About},
			{"latitude", expectedUser.Latitude},
			{"longitude", expectedUser.Longitude},
			{"data", expectedUser.Data},
			{"registered", expectedUser.Registered},
		}))
		userResponse, err := us.userStore.Get(context.Background(),expectedUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, userResponse)
	})
}

func (us *UserService) TestDeleteOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := us.userStore.DeleteUser(context.Background(),"1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := us.userStore.DeleteUser(context.Background(),"1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN")
		assert.NotNil(t, err)
	})
}
