// package definition
package main

// import of dependencies
import (
	"fmt"
	"github.com/jotdl/othello"
	"time"
)

// main function
func main() {
	// create a bot from our calculation method
	bot := othello.FuncPlayer(calculateNextTurn)

	// create a new othello game with our bot and a simple ai
	game := othello.NewGame(othello.NewTerminalPlayer(), bot)

	// as long as the game is not finished repeat game execution
	for !game.Finished() {
		fmt.Println(game) // print the current board

		err := game.DoNextMove() // calculate next move
		if err != nil {
			panic(fmt.Errorf("error during game execution %w", err))
		}

		time.Sleep(1 * time.Second)
	}

	// print the final result
	fmt.Printf("The winner is \"Player %v\"!!!\n", game.Winner())
}

// calculateNextTurn calculates the next turn of our bot based on the given board
func calculateNextTurn(board *othello.Board, currentPlayer othello.Color) othello.Turn {
	moves := othello.FindPossibleMoves(board, currentPlayer)

	return moves[0]
}
