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

func (srv *server) Run()  {

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal("Could not listen on #{srv.Addr} due to #{err.Error()}")
        }
    }()

    log.Printf("Api is ready to handle requedt #{srv.Addr}")
}
