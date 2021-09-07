package tweet

import (
	"github.com/EddCode/twitter-clone/internal/application/customError"
	models "github.com/EddCode/twitter-clone/internal/mooc/Tweet/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type TweetInterface interface {
	Save(tweet *models.Tweet) (string, *customError.CustomError)
}

type Tweet struct {
	Repository *TweetRepository
}

func NewTweetRepo(db *mongo.Client) TweetInterface {
	return &Tweet{Repository: NewRepository(db)}
}

func (repo *Tweet) Save(tweet *models.Tweet) (string, *customError.CustomError) {

	res, err := repo.Save(tweet)

	if err != nil {
		return "", err
	}

	objectId, _ := res.Inserted.(primitive.ObjectID)

	return objectId, nil
}
