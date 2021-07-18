package api

import "github.com/rs/cors"

func Start(port string)  {
    router := routes()
    handler := cors.AllowAll().Handler(router)

    server := newServer(port, handler)

    server.Run()
}
