package othello

// Player represents one participant in the othell game
type Player interface {
	// NextTurn has to calculate the next Turn of this player based on the given Board state and the color of this player
	NextTurn(*Board, Color) Turn
}

// FuncPlayer is a helper type to turn simple functions into players
type FuncPlayer func(*Board, Color) Turn

func (f FuncPlayer) NextTurn(b *Board, currentPlayer Color) Turn {
	return f(b, currentPlayer)
}
