package main

import (
	. "Threes/game"
	"fmt"
	"math/rand"
)

var numGames int
var cumScore float64

func main() {
	rand.Seed(42)
	for i := 0; i < 10000; i++ {
		playGame()
	}
	fmt.Printf("Avg score: %v\n", cumScore/float64(numGames))
}

func playGame() {
	game := NewGame()
	var score float64
	for !game.IsOver {
		//fmt.Println(game)
		//fmt.Println()
		direction := game.BestMove()
		//fmt.Printf("Moving %s\n", direction)
		game.Swipe(direction)
		score = game.BoardScore()
		//fmt.Printf("Score is now %v\n", score)
	}
	numGames++
	cumScore += score
}
