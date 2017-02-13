package sudokuboard

import (
	"reflect"
	"testing"
)

func TestBuildSudokuBoard(t *testing.T) {
	testBoard := BuildSudokuBoard("../testInputs/testSudokuBoard.txt")

	// Making example board

	testBoardArray := [][]int{}
	// Board amde inline
	testBoardArray = append(testBoardArray, []int{7, 0, 3, 0, 8, 0, 0, 4, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 9, 3, 5, 8})
	testBoardArray = append(testBoardArray, []int{0, 0, 4, 3, 0, 0, 0, 6, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 3, 6, 0, 1})
	testBoardArray = append(testBoardArray, []int{0, 3, 0, 0, 0, 0, 0, 8, 0})
	testBoardArray = append(testBoardArray, []int{1, 0, 6, 5, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 2, 7, 0, 0})
	testBoardArray = append(testBoardArray, []int{5, 4, 9, 6, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 7, 0, 0, 9, 0, 0, 0, 6})

	actual := testBoard.boardArray
	expected := testBoardArray
	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing BuildSudokuBoard: '%s' != '%s'", expected, actual)
	}
}

func TestFillBoardCell(t *testing.T) {
	originalBoard := BuildSudokuBoard("../testInputs/testSudokuBoard.txt")

	testBoard := originalBoard.fillBoardCell(8, 0, 2)

	testBoardArray := [][]int{}
	testBoardArray = append(testBoardArray, []int{7, 0, 3, 0, 8, 0, 0, 4, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 9, 3, 5, 8})
	testBoardArray = append(testBoardArray, []int{0, 0, 4, 3, 0, 0, 0, 6, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 3, 6, 0, 1})
	testBoardArray = append(testBoardArray, []int{0, 3, 0, 0, 0, 0, 0, 8, 0})
	testBoardArray = append(testBoardArray, []int{1, 0, 6, 5, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{0, 0, 0, 0, 0, 2, 7, 0, 0})
	testBoardArray = append(testBoardArray, []int{5, 4, 9, 6, 0, 0, 0, 0, 0})
	testBoardArray = append(testBoardArray, []int{2, 7, 0, 0, 9, 0, 0, 0, 6})

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
	testBoard := BuildSudokuBoard("../testInputs/testSudokuBoard.txt")

	isValid := isValidBoard(testBoard, 0, 8)

	actual := isValid
	expected := true
	if actual != expected {
		t.Errorf("Error occured while testing isValidBoard: '%s' != '%s'", expected, actual)
	}

	// Testing invalid case
	invalidBoard_1 := testBoard.fillBoardCell(1, 7, 3)
	invalidBoard_2 := testBoard.fillBoardCell(0, 1, 3)
	invalidBoard_3 := testBoard.fillBoardCell(1, 0, 5)

	isValid_1 := isValidBoard(invalidBoard_1, 1, 7)
	isValid_2 := isValidBoard(invalidBoard_2, 0, 1)
	isValid_3 := isValidBoard(invalidBoard_3, 1, 0)

	expected = false
	if (expected != isValid_1) || (expected != isValid_2) || (expected != isValid_3) {
		t.Errorf("Error occured while testing isValidBoard: '%s' != '%s', '%s', '%s'", expected, isValid_1, isValid_2, isValid_3)
	}
}

func TestChildren(t *testing.T) {
	testBoard := BuildSudokuBoard("../testInputs/testSudokuBoard.txt")
	children := testBoard.Children()
	freeCells := len(testBoard.freeCells)

	// Comparing number of children, actually comparing the children would be
	// VERY expensive
	expectedChildren := 182
	actualChildren := children.Size()

	expectedFree := 53
	actualFree := freeCells

	if expectedChildren != actualChildren {
		t.Errorf("Error occured while testing Children: '%s' != '%s'", expectedChildren, actualChildren)
	}

	if expectedFree != actualFree {
		t.Errorf("Error occured while testing Free Cells: '%s' != '%s'", expectedFree, actualFree)
	}
}

func TestDump(t *testing.T) {
	testBoard := BuildSudokuBoard("../testInputs/testSudokuBoard.txt")
	testBoard.Dump("testDumpBoard.txt")

	dumpedBoard := BuildSudokuBoard("../testInputs/testDumpBoard.txt")

	// Comparing number of children, actually comparing the children would be
	// VERY expensive
	expected := testBoard
	actual := dumpedBoard

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing Dump: '%s' != '%s'", expected, actual)
	}
}

func TestIsGoal(t *testing.T) {
	testBoard := BuildSudokuBoard("../testInputs/testSudokuBoard.txt")
	actual := testBoard.IsGoal()

	expected := false

	if actual != expected {
		t.Errorf("Error occured while testing IsGoal: '%s' != '%s'", expected, actual)
	}

	goalBoard := BuildSudokuBoard("../testInputs/testSudokuBoard_completed.txt")
	actual = goalBoard.IsGoal()

	expected = true

	if actual != expected {
		t.Errorf("Error occured while testing IsGoal: '%s' != '%s'", expected, actual)
	}
}
