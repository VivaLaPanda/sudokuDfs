package depthfirstsearch

import "testing"

func TestSolveBoard(t *testing.T) {
	board := SolveBoard("../testInputs/testSudokuBoard.txt")

	if board == nil {
		t.Error("Error occured while testing SolveBoard: No solution returned for solvable board")
	}
}
