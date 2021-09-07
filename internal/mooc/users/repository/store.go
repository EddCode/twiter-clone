package users

import (
	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"github.com/EddCode/twitter-clone/internal/storage"
	"github.com/EddCode/twitter-clone/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName string = "Users"

type Repository struct {
	Connection *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{Connection: db}
}

func (db *Repository) StoreUser(user *models.User) (*mongo.InsertOneResult, error) {
	collection, ctx, cancel := storage.DBCollection(db.Connection, collectionName)
	defer cancel()

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

func (db *Repository) IsUserExist(email string) (*models.User, bool, string) {
	collection, ctx, cancel := storage.DBCollection(db.Connection, collectionName)
	defer cancel()

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return &result, false, ID
	}

	return &result, true, ID

}

func (db *Repository) getUserById(id string) (*models.User, error) {
	collection, ctx, cancel := storage.DBCollection(db.Connection, collectionName)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": objectId}

	var profile models.User
	err := collection.FindOne(ctx, condition).Decode(&profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil

}

func (db *Repository) updateUserProfile(user map[string]interface{}, id primitive.ObjectID) (bool, error) {
	collection, ctx, cancel := storage.DBCollection(db.Connection, collectionName)
	defer cancel()

	updatedData := bson.M{
		"$set": user,
	}

	_, mongoError := collection.UpdateByID(ctx, id, updatedData)

	if mongoError != nil {
		return false, mongoError
	}

	return true, nil
}
