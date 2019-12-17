package main

import "time"

import "fmt"

func main() {
	bot1 := othello.NewSimpleBot()
	bot2 := MyBot{}

	game := othello.NewGame(bot1, bot2)

	for !game.Finished() {
		fmt.Println(game)

		game.DoNextMove()

		time.Sleep(2 * time.Millisecond)
	}

	fmt.Printf("And the winner is %q with a final score of %v:%v", game.Winner(), game.Player1Score(), game.Player2Score())
}
