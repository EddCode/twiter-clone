package tweet

import (
	"github.com/EddCode/twitter-clone/internal/application/customError"
	models "github.com/EddCode/twitter-clone/internal/mooc/tweet/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type TweetInterface interface {
	Save(tweet *models.Tweet) (string, *customError.CustomError)
}

type Tweet struct {
	*TweetRepository
}

func NewTweetRepo(db *mongo.Client) TweetInterface {
	return &Tweet{TweetRepository: NewRepository(db)}
}

func (repo *Tweet) Save(tweet *models.Tweet) (string, *customError.CustomError) {

	_, err := repo.SaveTweet(tweet)

	if err != nil {
		return "", customError.ThrowError("InternalServerError", err)
	}

	return "inserted", nil
}
