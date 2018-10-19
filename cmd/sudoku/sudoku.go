package main

import (
	"log"
	"net/http"

	"github.com/ernst01/sudoku-solver/src/server"
	"github.com/gorilla/mux"
)

func main() {
	srv := server.Server{
		Router: mux.NewRouter(),
	}

	srv.Routes()

	http.Handle("/", srv.Router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
