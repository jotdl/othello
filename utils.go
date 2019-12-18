package othello

type Field struct {
	Column int
	Row    int

	Value Color
}

type Turn struct {
	Column int
	Row    int
}

type Score [2]uint8

func FindPossibleMoves(board *Board, color Color) []Turn {
	turns := make([]Turn, 0, 20)

	for row := 0; row < board.Dimension; row++ {
		for column := 0; column < board.Dimension; column++ {
			if err := board.IsValidMove(row, column, color); err != nil {
				continue
			}

			turns = append(turns, Turn{Row: row, Column: column})
		}
	}

	return turns
}
