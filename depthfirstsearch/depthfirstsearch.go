package depthfirstsearch

import (
	ds "github.com/hishboy/gocommons/lang"
	sudobo "github.com/vivalapanda/sudokuDfs/sudokuboard"
)

func SolveBoard(filename string) *sudobo.SudokuBoard {
	board := sudobo.BuildSudokuBoard(filename)

	return depthFirstSearch(board)
}

// Function which searches the given board for a winning nodes
// If it succeeds it returns the board
// Otherwise it returns nill
func depthFirstSearch(board *sudobo.SudokuBoard) *sudobo.SudokuBoard {

	children := board.Children()
	fringe := ds.NewStack()

	// Move elemnts from children into stack
	for !children.Empty() {
		temp, _ := children.Pop()
		boardToPush := temp.(*sudobo.SudokuBoard)
		fringe.Push(boardToPush)
	}

	// DFS Implementation
	// We keep searching until we hit the goal or we run out of nodes
	temp := fringe.Pop()
	currentNode := temp.(*sudobo.SudokuBoard)
	for ; (fringe.Len() == 0) || currentNode.IsGoal(); currentNode = fringe.Pop().(*sudobo.SudokuBoard) {
		nodesToPushQueue := currentNode.Children()
		for !nodesToPushQueue.Empty() {
			temp, _ := nodesToPushQueue.Pop()
			boardToPush := temp.(*sudobo.SudokuBoard)
			fringe.Push(boardToPush)
		}
	}

	// If we've exited the loop, we can assume that if we aren't at the goalBoard
	// there is none
	currentNode.Dump("test.txt")
	if currentNode.IsGoal() {
		return currentNode
	} else {
		return nil
	}
}
