package server

import (
	"net/http"
	//"github.com/ernst01/sudoku-solver/src/server"
)

// Routes defines all our routes
func (s *Server) Routes() {
	http.HandleFunc("/sudoku", s.handleRandomGrid())
}
