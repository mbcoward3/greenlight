package server

import (
	"fmt"
	"net/http"
)

// healthCheckHandler writes basic diagnostic information about.
func (s *Server) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": s.config.Env,
		"version":     s.config.Version,
	}

	if err := s.writeJSON(w, http.StatusOK, data, nil); err != nil {
		s.logger.Error(err.Error())
		http.Error(w, "server error while processing the request", http.StatusInternalServerError)
	}
}

func (s *Server) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (s *Server) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := s.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)

		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
