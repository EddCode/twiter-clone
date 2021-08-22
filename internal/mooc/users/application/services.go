package users

import (
	"encoding/json"
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

	httpresponse.Success(&newUser, http.StatusCreated).Send(w)
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
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		httpresponse.BadRequest("missing parameters").Send(w)
		return
	}

	userProfile, err := service.userRepository.GetUserProfile(id)

	if err != nil {
		httpresponse.BadRequest(err.Error()).Send(w)
		return
	}

	httpresponse.Success(&userProfile, http.StatusOK).Send(w)
}

func (service *Service) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	var userProfile models.User

	err := json.NewDecoder(body).Decode(&userProfile)

	if err != nil {
		httpresponse.BadRequest("Incorrect data").Send(w)
	}

	service.userRepository.UpdateUserProfile(userProfile)

	httpresponse.Success()

}
