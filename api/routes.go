package api

import (
	users "github.com/EddCode/twitter-clone/internal/mooc/users/application"
	"github.com/gorilla/mux"
)

func routes(userService *users.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/singup", userService.SingupHandler).Methods("Post")
	router.HandleFunc("/login", userService.LoginHandler).Methods("POST")

	return router
}
