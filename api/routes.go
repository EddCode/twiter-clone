package api

import (
	"github.com/EddCode/twitter-clone/internal/middlewares"
	tweet "github.com/EddCode/twitter-clone/internal/mooc/tweet/application"
	users "github.com/EddCode/twitter-clone/internal/mooc/users/application"
	"github.com/gorilla/mux"
)

func routes(userService *users.Service, tweetService *tweet.TweetService) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/singup", userService.SingupHandler).Methods("Post")
	router.HandleFunc("/login", userService.LoginHandler).Methods("POST")
	router.HandleFunc("/profile", middlewares.AttachMiddelwares(userService.ProfileHandler, middlewares.ValidJWT())).Methods("GET")
	router.HandleFunc("/update-profile", middlewares.AttachMiddelwares(userService.UpdateProfileHandler, middlewares.ValidJWT())).Methods("PUT")

	router.HandleFunc("/post-tweet", middlewares.AttachMiddelwares(tweetService.SaveHandler, middlewares.ValidJWT())).Methods("POST")
	return router
}
