package board

import (
	. "Threes/tile"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	if len(board) != 4 {
		t.Error("Board should have 4 rows")
	}

	for _, row := range board {
		if len(row) != 4 {
			t.Error("Row should have 4 entires")
		}
	}
}

func TestBoard_AddTile(t *testing.T) {
	var tile Tile
	tile.SetValue(3)

	board := NewBoard()
	index := Index{0, 0}

	board.AddTile(tile, index)

	if board[index.Row][index.Column].IsEmpty() {
		t.Error("Tile should not be empty")
	}
}
