package othello

func NewMinMaxAIPlayer() Player {
	return &minmaxPlayer{}
}

type minmaxPlayer struct{}

func (b *minmaxPlayer) rankBoard(board *Board, player Color) int {
	score := board.Score()
	if player == Black {
		return int(score[0]) - int(score[1])
	} else {
		return int(score[1]) - int(score[0])
	}
}

func (b *minmaxPlayer) minmax(board *Board, originalPlayer Color, currentPlayer Color, currentDepth int) int {
	if currentDepth == 6 {
		return b.rankBoard(board, originalPlayer)
	}

	moves := FindPossibleTurns(board, currentPlayer)
	if len(moves) == 0 { // if no moves skip to next player's turn
		return b.minmax(board, originalPlayer, currentPlayer.Opposite(), currentDepth+1)
	}

	bestMoveVal := -99999 // for finding max
	if originalPlayer != currentPlayer {
		bestMoveVal = 99999 // for finding min
	}

	for _, move := range moves {
		childBoard := board.Clone()
		childBoard.MakeTurn(move.Row, move.Column, currentPlayer)

		val := b.minmax(childBoard, originalPlayer, currentPlayer.Opposite(), currentDepth+1)

		if currentPlayer == originalPlayer {
			if val > bestMoveVal {
				bestMoveVal = val
			}
		} else {
			if val < bestMoveVal {
				bestMoveVal = val
			}
		}
	}

	return bestMoveVal
}

func (b *minmaxPlayer) NextTurn(board *Board, currentPlayer Color) Turn {
	moves := FindPossibleTurns(board, currentPlayer)

	selectedMove := moves[0]
	bestMoveVal := -9999

	for _, move := range moves {
		childBoard := board.Clone()
		childBoard.MakeTurn(move.Row, move.Column, currentPlayer)

		val := b.minmax(childBoard, currentPlayer, currentPlayer.Opposite(), 1)
		if val > bestMoveVal {
			selectedMove = move
			bestMoveVal = val
		}
	}

	return selectedMove
}
