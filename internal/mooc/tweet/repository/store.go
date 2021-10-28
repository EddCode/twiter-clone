package tweet

import (
	tweet "github.com/EddCode/twitter-clone/internal/mooc/tweet/domain"
	"github.com/EddCode/twitter-clone/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName string = "Tweet"

type TweetRepository struct {
	Connection *mongo.Client
}

func NewRepository(db *mongo.Client) *TweetRepository {
	return &TweetRepository{Connection: db}
}

func (db *TweetRepository) SaveTweet(tweet *tweet.Tweet) (*mongo.InsertOneResult, error) {
	collection, ctx, cancel := storage.DBCollection(db.Connection, collectionName)
	defer cancel()

	row := bson.M{
		"userid":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(ctx, row)

	if err != nil {
		return nil, err
	}

	return result, nil
}
