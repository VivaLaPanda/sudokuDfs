package sudokuboard

import (
	"testing"
)

func TestBuildSudokuBoard(t *testing.T) {
	testBoard = BuildSudokuBoard()
	testBoardArray := [][]int{}
	boardStruct.testBoardArray = append(boardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})

	actual := testBoard.boardArray
	expected := testBoardArray
	if expected != actual {
		t.Errorf("Error occured while testing sayhello: '%s' != '%s'", expected, actual)
	}
}
