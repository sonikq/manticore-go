package main

import (
	"context"
	"net/http"
	"time"

	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg configs.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
