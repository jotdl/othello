package othello

import (
	"fmt"
)

func NewTerminalPlayer() Player {
	return &terminalPlayer{}
}

type terminalPlayer struct{}

func (b *terminalPlayer) NextTurn(board *Board, currentPlayer Color) Turn {

	row, column := 0, 0
	for gotValidMove := false; !gotValidMove; {
		row, column = b.askForNextMove(board.Dimension)

		err := board.IsValidMove(row, column, currentPlayer)
		gotValidMove = err == nil
		if err != nil {
			fmt.Printf("Move was not valid, due %q. \nPlease try again!\n\n", err)
		}
	}

	return Turn{Row: row, Column: column}
}

func (b *terminalPlayer) askForNextMove(dimension int) (int, int) {
	fmt.Printf("Please enter a number between 0 and %v and press ⏎ afterwards.\n", dimension-1)
	row := b.readNumber("Row", dimension-1)
	column := b.readNumber("Column", dimension-1)
	return row, column
}

func (b *terminalPlayer) readNumber(name string, maxValue int) (value int) {
	fmt.Printf("%v : ", name)
	_, err := fmt.Scan(&value)

	for err != nil || value < 0 || value > maxValue {
		fmt.Printf("\n\nInput was incorrect. \nPlease enter a number between 0 and %v and press ⏎ afterwards. Please try again! \n\n%v : ", maxValue, name)

		_, err = fmt.Scan(&value)
	}

	return value
}
