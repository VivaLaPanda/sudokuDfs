package sudokuboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	datastruct "github.com/hishboy/gocommons/lang"
)

type SudokuBoard struct {
	boardArray   [][]int
	freeChildren []int
	lock         *sync.Mutex
}

func BuildSudokuBoard(fileName string) *SudokuBoard {
	file, err := os.Open(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	boardStruct := &SudokuBoard{}
	// Convert our file to a 2D array of ints
	boardArray := [][]int{}          // The 2D slice represngint our board
	boardStruct.lock = &sync.Mutex{} // A mutex in case of later parallel implementations
	freeNodes := datastruct.NewStack()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read in one line of the file
		tempLine := scanner.Text()
		row := []int{}

		// Iterate through characters
		for i, r := range tempLine[:len(tempLine)] {
			// Trashing index
			_ = i

			// Convert chracters to ints
			tempInt, err := strconv.ParseInt(string(r), 10, 32)
			if err != nil {
				fmt.Print(err)
			}
			var newCell int
			newCell = int(tempInt)

			if newCell == 0 {
				freeNodes.Push(newCell)
			}

			// Put int's into an slice
			row = append(row, newCell)
		}

		// Put rows into a slice - our 2D array representing the board
		boardStruct.boardArray = append(boardArray, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return boardStruct
}
