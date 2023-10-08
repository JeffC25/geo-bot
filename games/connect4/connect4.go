// Got this from old project
// Need to refactor

package connect4

import (
	"fmt"
	"math"
	"math/rand"
)

type state struct {
	turn  int8
	board [][]int8
}

// intantiate new game state
func newGame() state {
	newState := state{}
	newState.turn = 1
	newState.board = [][]int8{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	return newState
}

// display board tile
func (s state) displayTile(row int8, col int8) {
	switch s.board[row][col] {
	case 0:
		fmt.Printf(" - ")
		return
	case 1:
		fmt.Printf(" X ")
	case 2:
		fmt.Printf(" O ")
	}
	return
}

// print board to stdout
func (s state) displayBoard(showCol bool, showBar bool) {
	if showCol {
		fmt.Println(" 1  2  3  4  5  6  7 ")
	}

	var row int8
	var col int8
	for row = 0; row < 6; row++ {
		for col = 0; col < 7; col++ {
			s.displayTile(row, col)
		}
		fmt.Println()
	}

	if showBar {
		fmt.Println("--------------------")
	}
}

// get move (col) from stdin -> input is 1-7, output is 0-6
func (s state) getMove() int8 {
	// get input
	var col int8
	fmt.Printf("Player %d's move: ", s.turn)
	fmt.Scanln(&col)

	for col > 7 || col < 1 || s.board[0][col-1] != 0 {
		// flush input buffer
		var discard string
		fmt.Scanln(&discard)

		// get input again
		fmt.Printf("Invalid input. Player %d's move: ", s.turn)
		fmt.Scanln(&col)
	}
	// offset col by 1 for array
	return col - 1
}

// get random move (col)
func (s state) randMove() int8 {
	col := int8(rand.Intn(7))
	for s.board[0][col] != 0 {
		col = int8(rand.Intn(7))
	}
	return col
}

// make player move
func (s state) makeMove(col int8) (int8, int8) {
	// look for first available tile in column
	var row int8
	for row = 5; row >= 0; row-- {
		if s.board[row][col] == 0 {
			// update tile
			s.board[row][col] = s.turn

			// return tile position
			return row, col
		}
	}
	return -1, -1
}

// update player turn
func (s *state) updateTurn() {
	s.turn = 1 + s.turn%2
}

// check for horizontal win
func (s state) checkHorizontal(row int8) bool {
	for col, counter := 0, 0; col < 7; col++ {
		// increment or reset counter
		if s.board[row][col] == (s.turn) {
			counter++
		} else {
			counter = 0
		}

		// check counter
		if counter == 4 {
			return true
		}
	}
	return false
}

// check for vertical win
func (s state) checkVertical(col int8) bool {
	for row, counter := 0, 0; row < 6; row++ {
		// increment or reset counter
		if s.board[row][col] == (s.turn) {
			counter++
		} else {
			counter = 0
		}

		// check counter
		if counter == 4 {
			return true
		}
	}
	return false
}

// check for diagonal win LT -> RB
func (s state) checkDiagonalLT(row int8, col int8) bool {
	rowLT := row - int8(math.Min(float64(row), float64(col)))
	colLT := col - int8(math.Min(float64(row), float64(col)))

	for counter := 0; rowLT < 6 && colLT < 7; rowLT, colLT = rowLT+1, colLT+1 {
		if s.board[rowLT][colLT] == (s.turn) {
			counter++
		} else {
			counter = 0
		}

		if counter == 4 {
			return true
		}
	}

	return false
}

// check for diagonal win LB -> RT
func (s state) checkDiagonalLB(row int8, col int8) bool {
	offset := int8(math.Min(float64(5-row), float64(col)))
	rowRB := row + offset
	colRB := col - offset

	for counter := 0; rowRB >= 0 && colRB < 7; rowRB, colRB = rowRB-1, colRB+1 {
		if s.board[rowRB][colRB] == (s.turn) {
			counter++
		} else {
			counter = 0
		}

		if counter == 4 {
			return true
		}
	}

	return false
}

func main() {
	gameState := newGame()
	gameState.displayBoard(true, true)

	// REPL
	for {
		move := gameState.getMove()
		moveRow, moveCol := gameState.makeMove(move)
		gameState.displayBoard(true, true)

		if gameState.checkHorizontal(moveRow) ||
			gameState.checkVertical(moveCol) ||
			gameState.checkDiagonalLT(moveRow, moveCol) ||
			gameState.checkDiagonalLB(moveRow, moveCol) {
			winner := gameState.turn
			fmt.Printf("Winner: Player %d\n", winner)
			return
		}

		gameState.updateTurn()
	}
}
