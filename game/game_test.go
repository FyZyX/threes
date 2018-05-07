package game

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame()

	if len(game.board) != 4 {
		t.Error("Incorrect board size")
	}
}

func TestGame_BestMove(t *testing.T) {
	game := NewGame()
	game.BestMove()
}
