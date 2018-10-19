package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ernst01/sudoku-solver/src/solver"
)

func (s *Server) handleRandomGrid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sudokuGrid := new(solver.SudokuGrid)
		if err := sudokuGrid.Init(jsonGrid); err != nil {
			sendError(w, http.StatusUnprocessableEntity, err.Error())
		}
		if err := sudokuGrid.Solve(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
		}

		respJSON, _ := json.Marshal(sudokuGrid)
		fmt.Fprintf(w, string(respJSON))
	}
}

//origGrid, SolvGrid := solver.StartSolving(jsonGrid)
//resp := &response{OriginalGrid: origGrid, SolvedGrid: SolvGrid}
