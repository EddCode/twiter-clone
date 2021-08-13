package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
		httpresponse.BadRequest("Missing email/password").Send(w)
		return
	}

	newUser, err := service.userRepository.Singup(&user)

	if err != nil {
		httpresponse.BadRequest(err.Error()).Send(w)
		return
	}

	httpresponse.Success(&newUser, 201).Send(w)
}

func (service *Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	var user models.UserLogin

	err := json.NewDecoder(body).Decode(&user)

	if err != nil {
		httpresponse.BadRequest("Missing email/password").Send(w)
		return
	}

	token, errLogin := service.userRepository.Login(&user)

	if errLogin != nil {
		httpresponse.UnauthoriedRequest(errLogin.Error()).Send(w)
		return
	}

	cookie := &http.Cookie{
		Name:     "Token",
		Value:    token.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	httpresponse.Success(token, http.StatusOK).Send(w)
}

func (service *Service) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("execiting Profile")
	service.userRepository.Find()
}

func (service *Service) StoreHandler(w http.ResponseWriter, r *http.Request) {
	service.userRepository.Store()
}
