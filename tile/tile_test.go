package tile

import (
	"testing"
)

func TestTile_IsEmpty(t *testing.T) {
	var tile Tile

	if !tile.IsEmpty() {
		t.Error("New tiles should be empty")
	}

	if tile.Value() != 0 {
		t.Error("Empty tiles should have value 0")
	}
}

func TestTile_SetValue(t *testing.T) {
	var tile Tile
	var value uint8 = 0
	tile.SetValue(value)

	if tile.value != value {
		t.Errorf("Unexpected value: should be %v, but got %v instead", tile.value, value)
	}
}

func TestTile_Value(t *testing.T) {
	tile := Tile{0}

	if tile.Value() != tile.value {
		t.Errorf("Unexpected value: should be %v, but got %v instead", tile.Value(), tile.value)
	}
}
