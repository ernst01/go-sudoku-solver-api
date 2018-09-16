package solver

import (
	"encoding/json"
	"fmt"
	"os"
)

var arrayGrid [][]int

// StartSolving starts solving the sudoku
func StartSolving(jsonGrid string) ([][]int, [][]int) {
	arrayGrid = ReadGrid(jsonGrid)

	if !validateGrid(arrayGrid) {
		fmt.Println("It is NOT a valid sudoku grid")
		os.Exit(0)
	}

	originalArrayGrid := make([][]int, 9)
	for i := range arrayGrid {
		originalArrayGrid[i] = make([]int, len(arrayGrid[i]))
		copy(originalArrayGrid[i], arrayGrid[i])
	}

	solve(0, 0)

	return originalArrayGrid, arrayGrid
}

func solve(pos_y int, pos_x int) {
	if pos_x == 9 {
		pos_y = pos_y + 1
		pos_x = 0
	}
	if pos_x == 0 && pos_y == 9 {
		return
	}
	if arrayGrid[pos_y][pos_x] == 0 {
		for val := 1; val <= 9; val++ {
			if true == isAllowed(val, arrayGrid, pos_y, pos_x) {
				//displayGrid(arrayGrid)
				//fmt.Println(fmt.Sprintf("%d : %d = %v", pos_y , pos_x, val))
				tmpNumber := arrayGrid[pos_y][pos_x]
				arrayGrid[pos_y][pos_x] = val
				solve(pos_y, pos_x+1)
				if isDone() {
					//fmt.Println("Welcome to my Donezo List")
					return
				}
				arrayGrid[pos_y][pos_x] = tmpNumber
			}
		}
		return
	} else {
		solve(pos_y, pos_x+1)
	}
	return
}

func isAllowed(val int, arrayGrid [][]int, posY int, posX int) bool {
	if true == integerInYSlice(val, arrayGrid[posY]) {
		return false
	}
	if true == integerInXSlice(val, arrayGrid, posX) {
		return false
	}
	if true == integerInSquareSlice(val, arrayGrid, posY, posX) {
		return false
	}
	return true
}

//fmt.Println(fmt.Sprintf("%d : %d is a number (%v)", pos_y, pos_x, arrayGrid[pos_y][pos_x]))

func validateGrid(arrayGrid [][]int) bool {
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

// ReadGrid is our starting point
func ReadGrid(jsonGrid string) [][]int {
	var grid [][]int
	if err := json.Unmarshal([]byte(jsonGrid), &grid); err != nil {
		panic(err)
	}

	return grid
}

func integerInYSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func integerInXSlice(a int, list [][]int, posX int) bool {
	for _, b := range list {
		if b[posX] == a {
			return true
		}
	}
	return false
}

func integerInSquareSlice(a int, list [][]int, posY int, posX int) bool {
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

func displayGrid(arrayGrid [][]int) {
	fmt.Println("------------------")
	for _, line := range arrayGrid {
		fmt.Println(line)
	}
}

func isDone() bool {
	for _, line := range arrayGrid {
		for _, a := range line {
			if a == 0 {
				return false
			}
		}
	}
	return true
}
