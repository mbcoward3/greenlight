package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Config contins the variables needed to initialize a new [Server].
type Config struct {
	Env     string
	Port    int
	Version string
}

// Server contains functionality to run the greenlight api.
type Server struct {
	config     *Config
	logger     *slog.Logger
	httpServer *http.Server
}

// New creates a new [Server] using the provided [Config] and registers api routes.
func New(cfg *Config) *Server {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	s := &Server{
		config: cfg,
		logger: logger,
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		IdleTimeout:  time.Minute,
		Handler:      s.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	s.httpServer = httpServer

	return s
}

// Run runs [http.Server.ListenAndServe] on the underlying http server.
func (s *Server) Run() error {
	s.logger.Info("starting server", "addr", s.httpServer.Addr, "env", s.config.Env)

	err := s.httpServer.ListenAndServe()
	s.logger.Error(err.Error())

	return err
}
