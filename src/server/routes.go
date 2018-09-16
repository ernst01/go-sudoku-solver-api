package server

import "net/http"

// Routes defines all our routes
func (s *Server) Routes() {
	http.HandleFunc("/sudoku", handleRandomGrid())
}
