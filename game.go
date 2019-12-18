package othello

import (
	"bytes"
	"errors"
	"fmt"
)

type Game struct {
	board *Board

	finished bool

	currentPlayer Color
	player1       Player
	player2       Player

	MoveCount int
}

func NewGame(player1, player2 Player) *Game {
	return &Game{
		board:         newBoard(8),
		currentPlayer: Black,
		player1:       player1,
		player2:       player2,
	}
}

func (g *Game) Finished() bool {
	return g.finished
}

func (g *Game) Board() *Board {
	return g.board.Clone()
}

func (g *Game) CurrentPlayer() Color {
	return g.currentPlayer
}

func (g *Game) String() string {
	buf := &bytes.Buffer{}

	score := g.Board().Score()
	fmt.Fprintf(buf, "Move #%v\t\tScore: %v:%v\n\n", g.MoveCount, score[0], score[1])
	fmt.Fprint(buf, g.board)

	return buf.String()
}

func (g *Game) DoNextMove() (Turn, error) {
	if g.Finished() {
		return Turn{}, errors.New("game already finished")
	}
	g.MoveCount++

	var player Player
	switch g.currentPlayer {
	case Black:
		player = g.player1
	case White:
		player = g.player2
	}

	turn := player.NextTurn(g.Board(), g.currentPlayer)

	err := g.board.MakeTurn(turn.Row, turn.Column, g.currentPlayer)
	if err != nil {
		return turn, err
	}

	g.currentPlayer = g.currentPlayer.Opposite()

	if len(FindPossibleMoves(g.board, g.currentPlayer)) == 0 {
		g.currentPlayer = g.currentPlayer.Opposite()

		if len(FindPossibleMoves(g.board, g.currentPlayer)) == 0 {
			g.finished = true
		}
	}

	return turn, nil
}

func (g *Game) Winner() Color {
	if !g.Finished() {
		return None
	}

	score := g.board.Score()
	switch {
	case score[0] < score[1]:
		return White
	case score[1] < score[0]:
		return Black
	}
	return None
}
