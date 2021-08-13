package api

import (
	"github.com/EddCode/twitter-clone/internal/middlewares"
	users "github.com/EddCode/twitter-clone/internal/mooc/users/application"
	"github.com/gorilla/mux"
)

func routes(userService *users.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/singup", userService.SingupHandler).Methods("Post")
	router.HandleFunc("/login", userService.LoginHandler).Methods("POST")
	router.HandleFunc("/profile", middlewares.AttachMiddelwares(userService.ProfileHandler, middlewares.ValidJWT())).Methods("GET")

	return router
}
