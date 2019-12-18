package othello

import "math"

func NewAdvancedAIPlayer() Player {
	return &advancedPlayer{}
}

type advancedPlayer struct{}

func (b *advancedPlayer) rankMove(turn Turn, dim float64) int {
	// it's a good idea to get the fields on the board so we rank them higher than in the middle (as they're easier to change)
	return int(math.Abs(dim/2.0-float64(turn.Row)) + math.Abs(dim/2.0-float64(turn.Column)))
}

func (b *advancedPlayer) rankBestEnemyReaction(board *Board, currentPlayer Color, move Turn) int {
	board = board.Clone()
	board.MakeTurn(move.Row, move.Column, currentPlayer)

	enemy := currentPlayer.Opposite()

	moves := FindPossibleTurns(board, enemy)
	if len(moves) == 0 {
		return -100000
	}

	dim := float64(board.Dimension)

	bestScore := b.rankMove(moves[0], dim)

	for _, move := range moves {
		score := b.rankMove(move, dim)
		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
}

func (b *advancedPlayer) NextTurn(board *Board, currentPlayer Color) Turn {
	moves := FindPossibleTurns(board, currentPlayer)

	dim := float64(board.Dimension)

	bestScore := b.rankMove(moves[0], dim) - b.rankBestEnemyReaction(board, currentPlayer, moves[0])
	selectedMove := moves[0]

	for _, move := range moves {
		score := b.rankMove(move, dim) - b.rankBestEnemyReaction(board, currentPlayer, move)
		if score > bestScore {
			bestScore = score
			selectedMove = move
		}
	}

	return selectedMove
}
