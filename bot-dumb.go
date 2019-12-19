package othello

func NewBraindeadAIPlayer() Player {
	return &braindeadPlayer{}
}

type braindeadPlayer struct{}

func (b *braindeadPlayer) NextTurn(board *Board, currentPlayer Color) Turn {
	moves := FindPossibleTurns(board, currentPlayer)

	return moves[0]
}
