package sudoku

import (
	"net/http"

	"github.com/ernst01/sudoku-solver/pkg/solver"
)

func (s *Server) handleRandomGrid() http.HandlerFunc {
	type SudokuGridResponse struct {
		Message      string  `json:"message"`
		TimeTaken    string  `json:"time_taken"`
		OriginalGrid [][]int `json:"original_grid"`
		SolvedGrid   [][]int `json:"solved_grid"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sudokuGrid := new(solver.SudokuGrid)
		if err := sudokuGrid.Init(jsonGrid); err != nil {
			sendError(w, http.StatusUnprocessableEntity, err.Error())
		}
		if err := sudokuGrid.Solve(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		resp := SudokuGridResponse{
			OriginalGrid: sudokuGrid.OriginalGrid,
			SolvedGrid:   sudokuGrid.SolvedGrid,
			TimeTaken:    sudokuGrid.TimeTaken,
		}
		if solver.VerifySolvedGrid(sudokuGrid.SolvedGrid) == true {
			resp.Message = "Good solving sparky!"
		} else {
			resp.Message = "It's ok... at least we gave it our best!"
		}

		sendSuccess(w, http.StatusOK, resp)
	}
}
