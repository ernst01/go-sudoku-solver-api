package solver

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SudokuGrid represents a sudoku grid
type SudokuGrid struct {
	OriginalGrid [][]int `json:"original_grid"`
	SolvedGrid   [][]int `json:"solved_grid"`
}

// Init inits the grid
func (sg *SudokuGrid) Init(jsonGrid string) error {
	if err := json.Unmarshal([]byte(jsonGrid), &sg.OriginalGrid); err != nil {
		return errors.New("Invalid grid")
	}

	sg.SolvedGrid = make([][]int, 9)
	for i := range sg.OriginalGrid {
		sg.SolvedGrid[i] = make([]int, len(sg.OriginalGrid[i]))
		copy(sg.SolvedGrid[i], sg.OriginalGrid[i])
	}
	return nil
}

// Solve starts solving the sudoku
func (sg *SudokuGrid) Solve() error {
	if !sg.validate(sg.OriginalGrid) {
		return errors.New("Not a valid grid")
	}

	sg.solve(0, 0)

	return nil
}

func (sg *SudokuGrid) solve(posY int, posX int) {
	if posX == 9 {
		posY = posY + 1
		posX = 0
	}
	if posX == 0 && posY == 9 {
		return
	}
	if sg.SolvedGrid[posY][posX] == 0 {
		for val := 1; val <= 9; val++ {
			if true == sg.isAllowed(val, sg.SolvedGrid, posY, posX) {
				//sg.display(sg.SolvedGrid)
				tmpNumber := sg.SolvedGrid[posY][posX]
				sg.SolvedGrid[posY][posX] = val
				sg.solve(posY, posX+1)
				if sg.isDone() {
					return
				}
				sg.SolvedGrid[posY][posX] = tmpNumber
			}
		}
		return
	}
	sg.solve(posY, posX+1)

	return
}

func (sg *SudokuGrid) isAllowed(val int, arrayGrid [][]int, posY int, posX int) bool {
	if true == sg.integerInYSlice(val, arrayGrid[posY]) {
		return false
	}
	if true == sg.integerInXSlice(val, arrayGrid, posX) {
		return false
	}
	if true == sg.integerInSquareSlice(val, arrayGrid, posY, posX) {
		return false
	}
	return true
}

func (sg *SudokuGrid) validate(arrayGrid [][]int) bool {
	if len(arrayGrid) != 9 {
		return false
	}
	for _, line := range arrayGrid {
		if len(line) != 9 {
			return false
		}
	}
	return true
}

func (sg *SudokuGrid) integerInYSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (sg *SudokuGrid) integerInXSlice(a int, list [][]int, posX int) bool {
	for _, b := range list {
		if b[posX] == a {
			return true
		}
	}
	return false
}

func (sg *SudokuGrid) integerInSquareSlice(a int, list [][]int, posY int, posX int) bool {
	startY := posY - (posY % 3)
	startX := posX - (posX % 3)
	for y := startY; y < (startY + 3); y++ {
		for x := startX; x < (startX + 3); x++ {
			if list[y][x] == a {
				return true
			}
		}
	}
	return false
}

func (sg *SudokuGrid) display(arrayGrid [][]int) {
	fmt.Println("------------------")
	for _, line := range arrayGrid {
		fmt.Println(line)
	}
}

func (sg *SudokuGrid) isDone() bool {
	for _, line := range sg.SolvedGrid {
		for _, a := range line {
			if a == 0 {
				return false
			}
		}
	}
	return true
}
