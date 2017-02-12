package sudokuboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	datastruct "github.com/hishboy/gocommons/lang"
	"github.com/oleiade/lane"
)

type boardCell struct {
	i        int
	j        int
	children []SudokuBoard
}

type SudokuBoard struct {
	boardArray [][]int
	freeCells  []boardCell
	lock       *sync.Mutex
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func BuildSudokuBoard(fileName string) *SudokuBoard {
	file, err := os.Open(fileName) // just pass the file name
	if err != nil {
		check(err)
	}
	defer file.Close()

	boardStruct := &SudokuBoard{}
	// Convert our file to a 2D array of ints
	boardStruct.lock = &sync.Mutex{} // A mutex in case of later parallel implementations
	boardStruct.freeCells = []boardCell{}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		// Read in one line of the file
		tempLine := scanner.Text()
		row := []int{}

		// Iterate through characters
		for _, r := range tempLine[:len(tempLine)] {
			// Convert chracters to ints
			tempInt, err := strconv.ParseInt(string(r), 10, 32)
			if err != nil {
				check(err)
			}
			var newCell int
			newCell = int(tempInt)

			// Put int's into an slice
			row = append(row, newCell)
		}
		// Put rows into a slice - our 2D array representing the board
		boardStruct.boardArray = append(boardStruct.boardArray, row)
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	// Get all free spaces and add them to a queue
	boardStruct.genFreeCells()

	return boardStruct
}

// When called on a board will return the board made by
// placing the value in the given row and column
// Will not modify the board it is called on
// WILL NOT preserve free cells, that should be handled before calling this function!
func (board SudokuBoard) fillBoardCell(row int, column int, value int) *SudokuBoard {
	// Deep copy board arra
	newBoardArray := [][]int{}
	for _, row := range board.boardArray {
		tempRow := []int{}
		for _, element := range row {
			tempRow = append(tempRow, element)
		}

		newBoardArray = append(newBoardArray, tempRow)
	}
	board.boardArray = newBoardArray

	if row > len(board.boardArray) || column > len(board.boardArray[0]) {
		fmt.Errorf("Row or Column out of bounds. Row: %s ; Column: %s\n", row, column)
	}

	board.boardArray[row][column] = value

	return &board
}

// Given a board determines if it is legal
// It assumes onle the cell at row,column has changed
func isValidBoard(board *SudokuBoard, row int, column int) bool {
	rowHash := datastruct.NewHashSet()
	columnHash := datastruct.NewHashSet()
	boxHash := datastruct.NewHashSet()

	for _, value := range board.boardArray[row] {
		if value != 0 {
			if rowHash.Contains(value) {
				return false
			} else {
				rowHash.Add(value)
			}
		}
	}

	for _, row := range board.boardArray {
		if row[column] != 0 {
			if columnHash.Contains(row[column]) {
				return false
			} else {
				columnHash.Add(row[column])
			}
		}
	}

	boxMinRow := (row / 3) * 3
	boxMinCol := (column / 3) * 3

	for boxRow := boxMinRow; boxRow < (boxMinRow + 3); boxRow++ {
		for boxCol := boxMinCol; boxCol < (boxMinCol + 3); boxCol++ {
			if board.boardArray[boxRow][boxCol] != 0 {
				if boxHash.Contains(board.boardArray[boxRow][boxCol]) {
					return false
				} else {
					boxHash.Add(board.boardArray[boxRow][boxCol])
				}
			}
		}
	}

	return true
}

func (board *SudokuBoard) genFreeCells() {
	for i, row := range board.boardArray {
		for j, element := range row {
			if element == 0 {
				cell := boardCell{
					i: i,
					j: j}

				board.lock.Lock()
				board.freeCells = append(board.freeCells, cell)
				board.lock.Unlock()
			}
		}
	}

	return
}

func (board *SudokuBoard) Children() *datastruct.Stack {
	// For every free cell on the sudokyu board
	childrenQueue := lane.NewPQueue(lane.MAXPQ)
	childrenStack := datastruct.NewStack()

	for i, cell := range board.freeCells {
		tempBoardSlice := []*SudokuBoard{}

		for cellValue := 1; cellValue < 10; cellValue++ {
			tempChild := board.fillBoardCell(cell.i, cell.j, cellValue)
			if isValidBoard(tempChild, cell.i, cell.j) {
				// Remove the cell from this child's freeCells

				tempChild.freeCells[len(tempChild.freeCells)-1], tempChild.freeCells[i] = tempChild.freeCells[i], tempChild.freeCells[len(tempChild.freeCells)-1]
				tempChild.freeCells = tempChild.freeCells[:len(tempChild.freeCells)-1]

				tempBoardSlice = append(tempBoardSlice, tempChild)
			}
		}

		for _, childBoard := range tempBoardSlice {
			childrenQueue.Push(childBoard, len(tempBoardSlice))
		}
	}

	for !childrenQueue.Empty() {
		temp, _ := childrenQueue.Pop()
		boardToPush := temp.(*SudokuBoard)
		childrenStack.Push(boardToPush)
	}

	return childrenStack
}

// Dumps the contents of a board to the provided file
func (board *SudokuBoard) Dump(filename string) error {
	// Open or create the file depending on whether it exists
	var file *os.File
	if _, err := os.Stat(filename); err == nil {
		temp, err := os.Open(filename) // just pass the file name
		file = temp
		check(err)
	} else {
		temp, err := os.Create(filename) // just pass the file name
		file = temp
		check(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, row := range board.boardArray {
		for _, cellInt := range row {
			fmt.Fprintf(w, "%s", strconv.Itoa(cellInt))
		}

		fmt.Fprint(w, "\n")
	}

	return w.Flush()
}
