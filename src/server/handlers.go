package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ernst01/sudoku-solver/src/solver"
)

func handleRandomGrid() http.HandlerFunc {
	type response struct {
		OriginalGrid [][]int `json:"original_grid"`
		SolvedGrid   [][]int `json:"solved_grid"`
	}
	arrayGrid := solver.ReadGrid(jsonGrid)
	return func(w http.ResponseWriter, r *http.Request) {
		origGrid, SolvGrid := startSolving(jsonGrid)
		resp := &response{OriginalGrid: origGrid, SolvedGrid: SolvGrid}
		resp_json, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(resp_json))
	}
}
