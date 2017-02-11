package sudokuboard

import (
	"reflect"
	"testing"
)

func TestBuildSudokuBoard(t *testing.T) {
	testBoard := BuildSudokuBoard("testSudokuBoard.txt")

	// Making example board

	testBoardArray := [][]int{}
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})

	actual := testBoard.boardArray
	expected := testBoardArray
	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing BuildSudokuBoard: '%s' != '%s'", expected, actual)
	}
}
