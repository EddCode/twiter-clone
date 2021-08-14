package users

import (
	"context"
	"time"

	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"github.com/EddCode/twitter-clone/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	connection *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{connection: db}
}

func (repo *Repository) StoreUser(user *models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := repo.connection.Database("twitterclone")
	collection := db.Collection("Users")

	password, errToHash := utils.HashPassword(user.Password, 6)

	if errToHash != nil {
		return nil, errToHash
	}

	user.Password = password

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Repository) IsUserExist(email string) (*models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := repo.connection.Database("twitterclone")
	collection := db.Collection("Users")

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return &result, false, ID
	}

	return &result, true, ID

}

func (repo *Repository) getUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := repo.connection.Database("twitterclone")
	collection := db.Collection("Users")
	objectId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": objectId}

	var profile models.User
	err := collection.FindOne(ctx, condition).Decode(&profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil

}
