package main

import (
	. "Threes/game"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(459)
	var numGames int
	var cumScore float64
	for i := 0; i < 10000; i++ {
		game := NewGame()
		var score float64
		for !game.IsOver {
			direction := game.BestMove()
			//fmt.Printf("Moving %s\n", direction)
			game.Swipe(direction)
			score = game.BoardScore()
			//fmt.Printf("Score is now %v\n", score)
		}
		numGames++
		cumScore += score
	}
	fmt.Printf("Avg score: %v\n", cumScore/float64(numGames))
}
