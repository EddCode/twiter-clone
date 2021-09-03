package tweet

import (
	"encoding/json"
	"net/http"

	"github.com/EddCode/twitter-clone/internal/httpresponse"
	models "github.com/EddCode/twitter-clone/internal/mooc/tweet/domain"
	tweet "github.com/EddCode/twitter-clone/internal/mooc/tweet/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type TweetService struct {
	tweet tweet.TweetRepository
}

func NewTweetService(db *mongo.Client) TweetService {
	return &TweetService{
		tweet: tweet.NewTweetRepo(db)
	}
}

func (service *TweetService) SaveHandler(w http.ResponseWriter, r *http.Request){
	body := r.Body
	defer body.Close()

	var tweet models.Tweet
	err := json.NewDecoder(body).Decode(&tweet)

	if err != nil {
		httpresponse.Error("BadRequest", "Wrong parameters").Send()
		return
	}

	_, errSave := service.tweet.Save(&tweet)

	if errSave != nil {
		httpresponse.Error(errSave.ErrorType(), errSave.Error()).Send(w)
		return
	}

	httpresponse.Success('success').Send(w)
}
