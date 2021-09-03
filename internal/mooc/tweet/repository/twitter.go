package tweet

import (
	"github.com/EddCode/twitter-clone/internal/application/customError"
	models "github.com/EddCode/twitter-clone/internal/mooc/Tweet/domain"
	"github.com/EddCode/twitter-clone/internal/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TweetInterface interface {
	Save(tweet *models.Tweet) (string, *customError.CustomError)
}

type TweetRepository struct {
	Tweet *storage.Repository
}

func NewTweetRepo(db *mongo.Client) TweetInterface {
	return &TweetRepo{Tweet: storage.NewRepository(db)}
}

func (repo *TweetRepository) Save(tweet *models.Tweet) (string, *customError.CustomError) {

	res, err := repo.Save(tweet)

	if err != nil {
		return "", err
	}

	objectId, _ := res.Inserted.(primitive.ObjectID)

	return objectId, nil
}
