package sudokuboard

import (
	"fmt"
	//"github.com/hishboy/gocommons/lang"
	"io/ioutil"
	"sync"
)

type SudokuBoard struct {
	boardArray   [][]int
	freeChildren []int
	lock         *sync.Mutex
}

func BuildSudokuBoard(fileName string) *SudokuBoard {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	charArray := []rune(string(b)) // convert file to array of characters
	for i, r := range charArray {
		fmt.Printf("i%d r %c\n", i, r)
	}

	boardStruct := &SudokuBoard{}
	boardStruct.lock = &sync.Mutex{}

	boardArray := [][]int{}
	boardStruct.boardArray = append(boardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	return boardStruct
}
