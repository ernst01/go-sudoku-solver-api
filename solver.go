package main

import (
    "os"
    "fmt"
    "encoding/json"
)

var arrayGrid [][]int

func startSolving(jsonGrid string) ([][]int, [][]int) {
    arrayGrid = readGrid(jsonGrid)

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
                tmp_number := arrayGrid[pos_y][pos_x]
                arrayGrid[pos_y][pos_x] = val
                solve(pos_y, pos_x + 1)
                if isDone() {
                    //fmt.Println("Welcome to my Donezo List")
                    return 
                }
                arrayGrid[pos_y][pos_x] = tmp_number
            }
        }
        return 
    } else {
        solve(pos_y, pos_x + 1)
    }
    return 
}

func isAllowed(val int, arrayGrid [][]int, pos_y int, pos_x int) bool {
    if true == integerInYSlice(val, arrayGrid[pos_y]) {
        return false
    }
    if true == integerInXSlice(val, arrayGrid, pos_x) {
        return false
    }
    if true == integerInSquareSlice(val, arrayGrid, pos_y, pos_x) {
        return false
    }
    return true;
}

//fmt.Println(fmt.Sprintf("%d : %d is a number (%v)", pos_y, pos_x, arrayGrid[pos_y][pos_x]))

func validateGrid(arrayGrid [][]int) bool {
    if len(arrayGrid) != 9 {
        return false;
    }
    for _, line := range arrayGrid {
        if (len(line) != 9) {
            return false;
        }
    }
    return true;
}

func readGrid(jsonGrid string) [][]int {
    var grid [][]int
    if err := json.Unmarshal([]byte(jsonGrid), &grid); err != nil {
        panic(err)
    }

    return grid;
}

func integerInYSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func integerInXSlice(a int, list [][]int, pos_x int) bool {
    for _, b := range list {
        if b[pos_x] == a {
            return true
        }
    }
    return false
}

func integerInSquareSlice(a int, list [][]int, pos_y int, pos_x int) bool {
    start_y := pos_y - (pos_y % 3);
    start_x := pos_x - (pos_x % 3);
    for y := start_y; y < (start_y + 3); y++ {
        for x := start_x; x < (start_x + 3); x++ {
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