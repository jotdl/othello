package othello

func NewSimpleAIPlayer() Player {
	return &simplePlayer{}
}

type simplePlayer struct{}

func (b *simplePlayer) NextTurn(board *Board, currentPlayer Color) Turn {
	moves := FindPossibleMoves(board, currentPlayer)

	currentScore := board.ScoreOf(currentPlayer)

	bestScore := currentScore
	selectedMove := moves[0]

	for _, move := range moves {
		board := board.Clone()
		board.MakeTurn(move.Row, move.Column, currentPlayer)

		score := board.ScoreOf(currentPlayer)
		if score > bestScore {
			bestScore = score
			selectedMove = move
		}
	}

	return selectedMove
}
