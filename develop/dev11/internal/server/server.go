package server

import (
	"context"
	"dev11/internal/config"
	"fmt"
	"net/http"
)

type HttpServer interface {
	RunServer() error
	Shutdown()
}

type server struct {
	server *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) HttpServer {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	server := server{server: &http.Server{
		Addr:    addr,
		Handler: handler,
	}}
	return &server
}

func (s *server) RunServer() error {
	return s.server.ListenAndServe()
}

func (s *server) Shutdown() {
	s.server.Shutdown(context.Background())
}
