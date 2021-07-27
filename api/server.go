package api

import (
	"log"
	"net/http"
	"time"
)

type server struct {
	*http.Server
}

func newServer(port string, handler http.Handler) *server {
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{Server: srv}
}

func (srv *server) Run() {

	log.Printf("Api is ready to handle request %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(`Could not listen on %s due to %s`, srv.Addr, err.Error())
	}
}
