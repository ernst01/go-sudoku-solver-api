package sudoku

import (
	"net/http"
)

// Routes defines all our routes
func (s *Server) Routes() {
	s.Router.Path("/sudoku").Methods("GET").
		HandlerFunc(corsHandler(s.handleRandomGrid()))
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
