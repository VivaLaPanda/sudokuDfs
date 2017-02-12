package sudokuboard

import (
	"reflect"
	"testing"
)

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

func TestFillBoardCell(t *testing.T) {
	originalBoard := BuildSudokuBoard("testSudokuBoard.txt")

	testBoard := originalBoard.fillBoardCell(8, 0, 1)

	testBoardArray := [][]int{}
	testBoardArray = append(testBoardArray, []int{1, 0, 0, 0, 0, 0, 0, 0, 2})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{1, 0, 0, 0, 0, 0, 0, 0, 1})

	actual := testBoard.boardArray
	expected := testBoardArray
	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing fillBoardCell: '%s' != '%s'", expected, actual)
	}

	if originalBoard.boardArray[8][0] == 1 {
		t.Error("Error occured while testing fillBoardCell: Original board was modified")
	}
}

func TestIsValidBoard(t *testing.T) {
	// Testing valid case
	testBoard := BuildSudokuBoard("testSudokuBoard.txt")

	isValid := isValidBoard(testBoard, 0, 8)

	actual := isValid
	expected := true
	if actual != expected {
		t.Errorf("Error occured while testing isValidBoard: '%s' != '%s'", expected, actual)
	}

	// Testing invalid case
	invalidBoard_1 := testBoard.fillBoardCell(1, 7, 2)
	invalidBoard_2 := testBoard.fillBoardCell(0, 1, 1)
	invalidBoard_3 := testBoard.fillBoardCell(1, 0, 1)

	isValid_1 := isValidBoard(invalidBoard_1, 1, 7)
	isValid_2 := isValidBoard(invalidBoard_2, 0, 1)
	isValid_3 := isValidBoard(invalidBoard_3, 1, 0)

	expected = false
	if (expected != isValid_1) || (expected != isValid_2) || (expected != isValid_3) {
		t.Errorf("Error occured while testing isValidBoard: '%s' != '%s', '%s', '%s'", expected, isValid_1, isValid_2, isValid_3)
	}
}

func TestChildren(t *testing.T) {
	testBoard := BuildSudokuBoard("testSudokuBoard.txt")
	children := testBoard.Children()

	// Comparing number of children, actually comparing the children would be
	// VERY expensive
	expected := 646
	actual := children.Len()

	if actual != expected {
		t.Errorf("Error occured while testing Children: '%s' != '%s'", expected, actual)
	}
}

func TestDump(t *testing.T) {
	testBoard := BuildSudokuBoard("testSudokuBoard.txt")
	testBoard.Dump("testDumpBoard.txt")

	dumpedBoard := BuildSudokuBoard("testDumpBoard.txt")

	// Comparing number of children, actually comparing the children would be
	// VERY expensive
	expected := testBoard
	actual := dumpedBoard

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing Dump: '%s' != '%s'", expected, actual)
	}
}
