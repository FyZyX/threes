package tile

import (
	"testing"
)

func TestCreateTile(t *testing.T) {
	var tile Tile

	if !tile.IsEmpty() {
		t.Error("New tiles should be empty")
	}

	if tile.Value() != 0 {
		t.Error("Empty tiles should have value 0")
	}
}

func TestSetTileValue(t *testing.T) {
	var tile Tile
	var value uint8 = 0
	tile.SetValue(value)

	if tile.value != value {
		t.Errorf("Unexpected value: should be %v, but got %v instead", tile.value, value)
	}
}
