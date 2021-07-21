package api

import (
	users "github.com/EddCode/twitter-clone/internal/mooc/users/application"
	"github.com/EddCode/twitter-clone/internal/storage"
	"github.com/rs/cors"
)

func Start(port string)  {
    db, _ := storage.NewMongoClient()

    router := routes(users.NewUserService(db))
    handler := cors.AllowAll().Handler(router)

    server := newServer(port, handler)

    server.Run()
}
