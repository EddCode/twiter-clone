package tweet

import (
	models "github.com/EddCode/twitter-clone/internal/mooc/Tweet/domain"
	tweet "github.com/EddCode/twitter-clone/internal/mooc/Tweet/domain"
	"github.com/EddCode/twitter-clone/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName string = 'Tweet'

func (repo *storage.Repository) saveTweet(tweet *models.Tweet) (*mongo.InsertOneResult, error) {
	collection, ctx, cancel := storage.DBCollection(repo.connection, collectionName)
	defer cancel()

	row := bson.M{
		"userid": tweet.UserId,
		"message": tweet.Message,
		"date": tweet.Date,
	}

	result, err := collection.InsertOne(ctx, row)

	if err != nil {
		return nil, err
	}

	return result, nil
}
