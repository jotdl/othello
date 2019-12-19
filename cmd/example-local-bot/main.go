// package definition
package main

// import of dependencies
import (
	"fmt"

	"github.com/jotdl/othello"
)

// main function
func main() {
	// create a bot from our calculation method
	bot := othello.NewMinMaxAIPlayer() // othello.FuncPlayer(calculateNextTurn)

	// create a new othello game with our bot and a simple ai
	game := othello.NewGame(bot, othello.NewAdvancedAIPlayer())

	// as long as the game is not finished repeat game execution
	for !game.Finished() {
		fmt.Println(game) // print the current board

		_, err := game.DoNextMove() // calculate next move
		if err != nil {
			panic(fmt.Errorf("error during game execution %w", err))
		}

		//time.Sleep(500 * time.Millisecond)
	}

	fmt.Println(game) // print the final board

	// print the final result
	fmt.Printf("The winner is \"Player %v\"!!!\n", game.Winner())
}
