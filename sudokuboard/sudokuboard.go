package sudokuboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	datastruct "github.com/hishboy/gocommons/lang"
)

type boardCell struct {
	i        int
	j        int
	children []SudokuBoard
	lock     *sync.Mutex
}

type SudokuBoard struct {
	boardArray [][]int
	freeCells  *datastruct.Stack
	lock       *sync.Mutex
}

func BuildSudokuBoard(fileName string) *SudokuBoard {
	file, err := os.Open(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	boardStruct := &SudokuBoard{}
	// Convert our file to a 2D array of ints
	boardStruct.lock = &sync.Mutex{} // A mutex in case of later parallel implementations
	freeCells := datastruct.NewStack()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		// Read in one line of the file
		tempLine := scanner.Text()
		row := []int{}

		// Iterate through characters
		for column, r := range tempLine[:len(tempLine)] {
			// Convert chracters to ints
			tempInt, err := strconv.ParseInt(string(r), 10, 32)
			if err != nil {
				fmt.Print(err)
			}
			var newCell int
			newCell = int(tempInt)

			if newCell == 0 {

				cellStruct := &boardCell{
					i:        lineNumber,
					j:        column,
					children: []SudokuBoard{},
					lock:     &sync.Mutex{}}

				freeCells.Push(cellStruct)
			}

			// Put int's into an slice
			row = append(row, newCell)
		}
		// Put rows into a slice - our 2D array representing the board
		boardStruct.boardArray = append(boardStruct.boardArray, row)
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	boardStruct.freeCells = freeCells

	return boardStruct
}

// func fillBoardCell (row, column, value) *SudokuBoard {
//
// }

// Given a board determines if it is legal
// It assumes onle the cell at row,column has changed
func checkValidBoard(board *SudokuBoard, row int, column int) bool {
	rowHash := datastruct.NewHashSet()
	columnHash := datastruct.NewHashSet()
	boxHash := datastruct.NewHashSet()

	for _, value := range board.boardArray[row] {
		if value != 0 {
			if rowHash.Contains(value) {
				return false
			} else {
				rowHash.Add(value)
			}
		}
	}

	for _, row := range board.boardArray {
		if row[column] != 0 {
			if columnHash.Contains(row[column]) {
				return false
			} else {
				columnHash.Add(row[column])
			}
		}
	}

	boxMinRow := (row / 3) * 3
	boxMinCol := (column / 3) * 3

	for boxRow := boxMinRow; boxRow < (boxMinRow + 3); boxRow++ {
		for boxCol := boxMinCol; boxCol < (boxMinCol + 3); boxCol++ {
			if board.boardArray[boxRow][boxCol] != 0 {
				if boxHash.Contains(board.boardArray[boxRow][boxCol]) {
					return false
				} else {
					boxHash.Add(board.boardArray[boxRow][boxCol])
				}
			}
		}
	}

	return true
}

func (board *SudokuBoard) GenChildren() {
	// For every free cell on the sudokyu board
	for i := 0; i < board.freeCells.Len(); i++ {
		// Getting the cell from the stack and storing it
		// Weird code to assert type in multiple assignment
		// http://stackoverflow.com/questions/11403050/idiomatic-way-to-do-conversion-type-assertion-on-multiple-return-values-in-go
		temp, err := board.freeCells.Get(i)
		cell := temp.(boardCell)
		if err != nil {
			fmt.Print(err)
		}

		// Locking the cell for parallelisation
		cell.lock.Lock()

		// For every possible value in the cell, generate a board and test if it's valid
		// If it's a valid board store it in the children of that cell
		for cellValue := 1; cellValue < 10; cellValue++ {

		}
	}

	return
}
