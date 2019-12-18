package othello

// Field represents one field on an othello board
type Field struct {
	Column int // column of this field
	Row    int // row of this column

	Value Color // current value of this column
}

// Turn represents one player move in this game
type Turn struct {
	Column int // column of this move
	Row    int // row of this move
}

// Score represents the current score of a othello game. Index 0 represents score of player 1 while index 1 represents the score of player 2
type Score [2]uint8

// FindPossibleTurns returns all possible moves for the given board and color.
// Returns an empty slice if there are no more possible moves
func FindPossibleTurns(board *Board, color Color) []Turn {
	turns := make([]Turn, 0, 20)

	for row := 0; row < board.Dimension; row++ {
		for column := 0; column < board.Dimension; column++ {
			if err := board.IsValidTurn(row, column, color); err != nil {
				continue
			}

			turns = append(turns, Turn{Row: row, Column: column})
		}
	}

	return turns
}
