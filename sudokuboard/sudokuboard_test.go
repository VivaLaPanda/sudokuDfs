package sudokuboard

import (
	"reflect"
	"testing"
)

func TestCheckValidBoard(t *testing.T) {
	testBoard := BuildSudokuBoard("testSudokuBoard.txt")

	isValid := checkValidBoard(testBoard, 0, 8)

	actual := isValid
	expected := true
	if actual != expected {
		t.Errorf("Error occured while testing TestcheckValidBoard: '%s' != '%s'", expected, actual)
	}
}

func TestBuildSudokuBoard(t *testing.T) {
	testBoard := BuildSudokuBoard("testSudokuBoard.txt")

	// Making example board

	testBoardArray := [][]int{}
	// Board amde inline
	testBoardArray = append(testBoardArray, []int{1, 0, 0, 0, 0, 0, 0, 0, 2})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})

	actual := testBoard.boardArray
	expected := testBoardArray
	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing BuildSudokuBoard: '%s' != '%s'", expected, actual)
	}
}
