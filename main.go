package main

import (
	"bufio"
	"fmt"
	"os"

	dfs "github.com/VivaLaPanda/sudokuDfs/depthfirstsearch"
)

func main() {
	fmt.Println("Please enter the sudoku text file (e.g. file.txt): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Thinking... ")
	board := dfs.SolveBoard(scanner.Text())

	if board != nil {
		fmt.Println("Solution Found!")
		fmt.Println("Please enter the file to save your solution in (e.g. file.txt): ")
		scanner.Scan()
		board.Dump(scanner.Text())
	}
}
