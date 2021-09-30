package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

var BOARD [][]byte

// All possible Tetris figures variations
var FIGURES [][][]byte = [][][]byte{
	{{35, 35, 35}, {46, 35, 46}},
	{{35, 46}, {35, 35}, {35, 46}},
	{{46, 35, 46}, {35, 35, 35}},
	{{46, 35}, {35, 35}, {46, 35}},
	{{35, 35}, {35, 35}},
	{{35}, {35}, {35}, {35}},
	{{35, 35, 35, 35}},
	{{35, 35, 46}, {46, 35, 35}},
	{{46, 35}, {35, 35}, {35, 46}},
	{{46, 35, 35}, {35, 35, 46}},
	{{35, 46}, {35, 35}, {46, 35}},
	{{46, 35}, {46, 35}, {35, 35}},
	{{35, 35, 35}, {46, 46, 35}},
	{{35, 35}, {35, 46}, {35, 46}},
	{{35, 46, 46}, {35, 35, 35}},
	{{35, 46}, {35, 46}, {35, 35}},
	{{46, 46, 35}, {35, 35, 35}},
	{{35, 35}, {46, 35}, {46, 35}},
	{{35, 35, 35}, {35, 46, 46}},
}

type Tetrominoes struct {
	form [][]byte
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	CheckError(err)

	sep := []byte{13, 10}
	transformedContent := bytes.Split(content, sep)

	CheckFormat(transformedContent)

	tetrominoesList := FindTetrominoes(transformedContent)
	minSize := FindBoartMinSize(tetrominoesList)

	CreateBoard(minSize)
	TryPosition(0, tetrominoesList, minSize)
}

// Checks if file exists
func CheckError(err error) {
	if err != nil {
		fmt.Println("Missing file")
		os.Exit(3)
	}
}

// Checks file format
func CheckFormat(transformedContent [][]byte) {
	index := 0

	for i := 0; i < len(transformedContent); i++ {
		if len(transformedContent[i]) == 0 {
			if index != 4 {
				fmt.Println("Error. Bad format or figure doesn't consist of 4 cubes.")
				os.Exit(1)
			} else {
				index = 0
			}
		}
		for k := 0; k < len(transformedContent[i]); k++ {
			if transformedContent[i][k] == 35 {
				index++
			}
		}
	}
}

// Finds all tetrominoes from the file
func FindTetrominoes(transformedContent [][]byte) []Tetrominoes {
	var tetrominoToAppend Tetrominoes
	var tetrominoesList []Tetrominoes
	tetrisCounter := 1

	for a := 0; a < len(transformedContent); a++ {
		for b := 0; b < len(transformedContent[a]); b++ {
			var found bool = false
			for i := 0; i < len(FIGURES); i++ {
				var skip bool = false
				for k := 0; k < len(FIGURES[i]); k++ {
					for m := 0; m < len(FIGURES[i][k]); m++ {
						if a+k == tetrisCounter*5-1 || b+m >= 4 {
							skip = true
							break
						}
						if transformedContent[a+k][b+m] == FIGURES[i][k][m] {
							continue
						}
						skip = true
						break
					}
					if skip {
						break
					}
				}
				if !skip {
					found = true
					tetrominoToAppend.form = FIGURES[i]
					tetrominoesList = append(tetrominoesList, tetrominoToAppend)
					break
				}
			}
			if found {
				if a+(tetrisCounter*5-a) >= len(transformedContent) {
					a = len(transformedContent) - 1
				} else {
					a = a + (tetrisCounter*5 - a) - 1
				}
				tetrisCounter++
				break
			}

			if a == (tetrisCounter*5)-2 && b == 3 {
				fmt.Println("Error. Unknown tetromino.")
				os.Exit(2)
			}
		}
	}

	return tetrominoesList
}

// Calculates minimum possible board size
func FindBoartMinSize(tetrominoesList []Tetrominoes) int {
	minSideSize := 0
	for i := 0; i < len(tetrominoesList); i++ {
		for k := 0; k < len(tetrominoesList[i].form); k++ {
			if minSideSize < len(tetrominoesList[i].form[k]) {
				minSideSize = len(tetrominoesList[i].form[k])
			}

			if minSideSize < len(tetrominoesList[i].form) {
				minSideSize = len(tetrominoesList[i].form)
			}

			if minSideSize == 4 {
				break
			}
		}
	}

	blockCounter := math.Sqrt(float64(len(tetrominoesList) * 4))
	_, frac := math.Modf(blockCounter)
	if frac != 0 {
		blockCounter = math.Floor(blockCounter) + 1
	}

	return int(math.Max(blockCounter, float64(minSideSize)))
}

// Finds all tetrominoes from the file
func CreateBoard(size int) {
	BOARD = nil

	for i := 0; i < size; i++ {
		BOARD = append(BOARD, nil)
		for k := 0; k < size; k++ {
			BOARD[i] = append(BOARD[i], 0)
		}
	}
}

// Position selection start
func TryPosition(piece int, tetrominoesList []Tetrominoes, size int) {
	for y := 0; y < len(BOARD); y++ {
		for x := 0; x < len(BOARD); x++ {
			if CheckPosition(y, x, piece, tetrominoesList) {
				if y == len(BOARD)-1 || piece == len(tetrominoesList)-1 {
					PrintBoard()
					os.Exit(0)
				} else {
					TryPosition(piece+1, tetrominoesList, size)
				}
				ClearPosition(y, x, piece, tetrominoesList)
			}
		}
	}

	if piece == 0 {
		increaseSize := size + 1
		CreateBoard(increaseSize)
		TryPosition(0, tetrominoesList, increaseSize)
	}
}

// Checks if applicable position is free
func CheckPosition(y int, x int, piece int, tetrominoesList []Tetrominoes) bool {
	for i := 0; i < len(tetrominoesList[piece].form); i++ {
		if len(tetrominoesList[piece].form)+y > len(BOARD) || len(tetrominoesList[piece].form[i])+x > len(BOARD) {
			return false
		}
	}

	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if tetrominoesList[piece].form[a-y][b-x] == 46 {
				continue
			}

			if BOARD[a][b] == 0 {
				BOARD[a][b] = byte(65 + piece)
			} else {
				ClearPosition(y, x, piece, tetrominoesList)
				return false
			}
		}
	}

	return true
}

// Clears certain position after apply
func ClearPosition(y int, x int, piece int, tetrominoesList []Tetrominoes) {
	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if (tetrominoesList[piece].form[a-y][b-x]) == 46 {
				continue
			}
			if BOARD[a][b] == byte(65+piece) {
				BOARD[a][b] = 0
			}
		}
	}
}

// Prints final board
func PrintBoard() {
	for i := 0; i < len(BOARD); i++ {
		for k := 0; k < len(BOARD); k++ {
			if BOARD[i][k] == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(string(BOARD[i][k]))
		}
		fmt.Println()
	}
}
