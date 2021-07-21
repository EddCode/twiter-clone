package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EddCode/twitter-clone/internal/httpresponse"
	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	users "github.com/EddCode/twitter-clone/internal/mooc/users/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
    userRepository users.UserRepository
}

func NewUserService(db *mongo.Client) *Service {
    return &Service{
        userRepository: users.NewUserRepository(db),
    }
}

func (service *Service) SingupHandler(w http.ResponseWriter, r *http.Request) {
    body := r.Body
    defer body.Close()

    var user models.SingupUser
    err := json.NewDecoder(body).Decode(&user)

    if err != nil {
        httpresponse.BadRequest("invalid json").Send(w)
        return
    }

    newUser, err := service.userRepository.Singup(&user)

    if err != nil {
        httpresponse.BadRequest(err.Error()).Send(w)
        return
    }

    log.Print(user)
    httpresponse.Success(&newUser, 200).Send(w)
}

func (service *Service) FindHandler(w http.ResponseWriter, r *http.Request) {
    service.userRepository.Find()
}

func (service *Service) StoreHandler(w http.ResponseWriter, r *http.Request) {
    service.userRepository.Store()
}
