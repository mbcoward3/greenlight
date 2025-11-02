package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes creates a new [http.ServeMux] and registers routes.
func (s *Server) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", s.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", s.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", s.showMovieHandler)

	return router
}
