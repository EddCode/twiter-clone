package api

import (
	tweet "github.com/EddCode/twitter-clone/internal/mooc/tweet/application"
	users "github.com/EddCode/twitter-clone/internal/mooc/users/application"
	"github.com/EddCode/twitter-clone/internal/storage"
	"github.com/rs/cors"
)

func Start(port string) {
	db, _ := storage.NewMongoClient()

	router := routes(users.NewUserService(db.Client), tweet.NewTweetService(db.Client))
	handler := cors.AllowAll().Handler(router)

	server := newServer(port, handler)

	server.Run()
}
