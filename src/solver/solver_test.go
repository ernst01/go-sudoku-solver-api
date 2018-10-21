package solver_test

import (
	"reflect"
	"testing"

	"github.com/ernst01/sudoku-solver/src/solver"
)

var jsonGrid = `[
   [0, 4, 8, 0, 1, 2, 6, 0, 5],
   [6, 3, 0, 0, 4, 0, 0, 0, 0],
   [7, 0, 0, 8, 0, 6, 3, 0, 0],
   [4, 0, 3, 5, 0, 0, 0, 0, 0],
   [0, 0, 6, 4, 0, 1, 5, 0, 0],
   [0, 0, 0, 0, 0, 9, 1, 0, 4],
   [0, 0, 9, 7, 0, 3, 0, 0, 1],
   [0, 0, 0, 0, 9, 0, 0, 5, 2],
   [1, 0, 4, 2, 6, 0, 7, 9, 0]
]`

func TestSolver(t *testing.T) {
	sudokuGrid := new(solver.SudokuGrid)
	reflectInstance := reflect.TypeOf(sudokuGrid)
	if reflectInstance.String() != "*solver.SudokuGrid" {
		t.Errorf("Wrong instance type")
	}
	if err := sudokuGrid.Init(jsonGrid); err != nil {
		t.Errorf("Unable to init a test json grid")
	}
	if err := sudokuGrid.Solve(); err != nil {
		t.Errorf("Unable to solve")
	}
	if solver.VerifySolvedGrid(sudokuGrid.SolvedGrid) == false {
		t.Errorf("Unable to verify it was solved")
	}
	sudokuGrid.Display(sudokuGrid.SolvedGrid)
}

var jsonInvalidGrid = `[
   ????0, 4, 8, 0, 1, 2, 6, 0, 5],
   [6, 3, 0, 0, 4, 0, 0, 0, 0],
   [7, 0, 0, 8, 0, 6, 3, 0, 0],
   [4, 0, 3, 5, 0, 0, 0, 0, 0],
   [0, 0, 6, 4, 0, 1, 5, 0, 0],
   [0, 0, 0, 0, 0, 9, 1, 0, 4],
   [0, 0, 9, 7, 0, 3, 0, 0, 1],
   [0, 0, 0, 0, 9, 0, 0, 5, 2],
   [1, 0, 4, 2, 6, 0, 7, 9, 0]
]`

func TestInvalidGrid(t *testing.T) {
	sudokuGrid := new(solver.SudokuGrid)
	if err := sudokuGrid.Init(jsonInvalidGrid); err == nil {
		t.Errorf("Shouldn't have initialized")
	}
}
