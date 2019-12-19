package othello

import (
	"bytes"
	"errors"
	"fmt"
)

// Game is a simple helper struct to execute a game of othello
type Game struct {
	board *Board

	finished bool

	currentPlayer Color
	player1       Player
	player2       Player

	// Number of already made Moves/Turns. Basically current round
	MoveCount int
}

// NewGame of the 2 given players. Player1 will play as black while Player2 plays as white
// Creates a default board with 8x8 fields and the initial state like shown below:
//       0   1   2   3   4   5   6   7
//     ---------------------------------
//   0 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   1 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   2 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   3 |   |   |   | ● | ○ |   |   |   |
//     ---------------------------------
//   4 |   |   |   | ○ | ● |   |   |   |
//     ---------------------------------
//   5 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   6 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   7 |   |   |   |   |   |   |   |   |
//     ---------------------------------
func NewGame(player1, player2 Player) *Game {
	return &Game{
		board:         newBoard(8),
		currentPlayer: Black,
		player1:       player1,
		player2:       player2,
	}
}

// Finished returns true if no more turns can be mad and a winner with a final score is determined
func (g *Game) Finished() bool {
	return g.finished
}

// Board of the current game. Always returns a clone so nobody else can modify it
func (g *Game) Board() *Board {
	return g.board.Clone()
}

// CurrentPlayer returns the color (black/white) of the currently active player
func (g *Game) CurrentPlayer() Color {
	return g.currentPlayer
}

func (g *Game) String() string {
	buf := &bytes.Buffer{}

	score := g.Board().Score()
	fmt.Fprintf(buf, "Round #%2d\t\tScore: %2d:%2d\n", g.MoveCount, score[0], score[1])
	fmt.Fprintf(buf, "Current: Player %v\t(● : Player 1, ○ : Player 2)\n\n", int(g.currentPlayer))
	fmt.Fprint(buf, g.board)

	return buf.String()
}

// DoNextMove calculates the next move based on the turn calculated by currently active player. Returns the made turn on success. If an error happened it will return an error instead.
// Furthermore it:
// - Increases move counter
// - Switches player to whoever has to move next (if only one player can turn this player will stay active)
// - Recalculates finish game state
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

	if len(FindPossibleTurns(g.board, g.currentPlayer)) == 0 {
		g.currentPlayer = g.currentPlayer.Opposite()

		if len(FindPossibleTurns(g.board, g.currentPlayer)) == 0 {
			g.finished = true
		}
	}

	return turn, nil
}

// Winner returns the winner of this game if it is finished. Will return None until the game is finished.
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
