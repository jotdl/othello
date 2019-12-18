package othello

type Player interface {
	NextTurn(*Board, Color) Turn
}

type FuncPlayer func(*Board, Color) Turn

func (f FuncPlayer) NextTurn(b *Board, currentPlayer Color) Turn {
	return f(b, currentPlayer)
}
