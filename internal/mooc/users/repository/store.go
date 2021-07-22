package users

import (
	"context"
	"log"
	"time"

	users "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"github.com/EddCode/twitter-clone/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type Repository struct {
    connection *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository{
    return &Repository{ connection: db }
}

func (repo *Repository)  StoreUser(user *users.User) (*mongo.InsertOneResult, error)  {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    db := repo.connection.Database("twitterclone")
    collection := db.Collection("Users")

    password, errToHash := utils.HashPassword(user.Password, 6)

    log.Print("im here yea!!!")
    if errToHash != nil {
        return nil, errToHash
    }

    user.Password = password

    result, err := collection.InsertOne(ctx, user)

    log.Print("im here yea 2!!!")
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (repo *Repository) IsUserExist(email string) (*users.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    db := repo.connection.Database("twitterclone")
    collection := db.Collection("Users")

    condition := bson.M{"email":email}

    var result users.User

    err := collection.FindOne(ctx, condition).Decode(&result)
    log.Print("try to find something", err)

    if err != nil {
        return &result, nil
    }

    return &result, nil

}
