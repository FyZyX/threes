package board

import (
	. "Threes/tile"
	"fmt"
	"testing"
)

func TestBoard(t *testing.T) {
	var board Board

	if len(board) != 4 {
		t.Error("Board should have 4 rows")
	}

	for _, row := range board {
		if len(row) != 4 {
			t.Error("Row should have 4 entires")
		}
	}
}

func TestBoard_String(t *testing.T) {
	var board Board
	boardString := `|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |`

	expected := boardString
	actual := fmt.Sprint(board)
	if actual != expected {
		t.Errorf("New board printed incorecctly. Got:\n%v\nWanted:\n%v\n", actual, expected)
	}
}

func TestRow_String(t *testing.T) {
	var row row
	rowString := `|     0 |     0 |     0 |     0 |`

	expected := rowString
	actual := fmt.Sprint(row)
	if actual != expected {
		t.Errorf("New row printed incorecctly. Got:\n%v\nWanted:\n%v\n", actual, expected)
	}
}

func TestBoard_AddTile(t *testing.T) {
	var tile Tile
	tile.SetValue(3)

	var board Board
	index := Index{0, 0}

	board.AddTile(tile, index)

	if board[index.Row][index.Column].IsEmpty() {
		t.Error("Tile should not be empty")
	}
}
