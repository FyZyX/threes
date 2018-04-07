package tile

import (
	"fmt"
	"testing"
)

func TestTile_Value(t *testing.T) {
	tile := Tile{0}

	if tile.Value() != tile.value {
		t.Errorf("Unexpected value: wanted %v, but got %v instead", tile.value, tile.Value())
	}
}

func TestTile_SetValue(t *testing.T) {
	var tile Tile
	var value Value
	tile.SetValue(value)

	if tile.value != value {
		t.Errorf("Unexpected value: should be %v, but got %v instead", value, tile.value)
	}
}

func TestTile_IsEmpty(t *testing.T) {
	var tile Tile

	if !tile.IsEmpty() {
		t.Error("New tiles should be empty")
	}

	if tile.Value() != 0 {
		t.Error("Empty tiles should have value 0")
	}
}

func TestMerge(t *testing.T) {
	t1 := Tile{3}
	t2 := Tile{3}

	tile, err := Merge(t1, t2)
	if err != nil {
		t.Error("Tiles of equal value were not merged")
	}

	if tile.value != t1.value+t2.value {
		t.Errorf("Unepexted merge value: wanted %v, but got %v", t1.value+t2.value, tile.value)
	}

	t1.SetValue(6)

	tile, err = Merge(t1, t2)
	if err == nil {
		t.Error("Tiles of unequal value cannot be merged")
	}

	t1.SetValue(1)
	t2.SetValue(2)

	tile, err = Merge(t1, t2)
	if err != nil {
		t.Error("Tile with value 1 not merged with tile of value 2")
	}

	if tile.value != 3 {
		t.Errorf("Unepexted merge value: wanted %v, but got %v", 3, tile.value)
	}

	t1.SetValue(1)
	t2.SetValue(1)

	tile, err = Merge(t1, t2)
	if err == nil {
		t.Error("Tiles of value 1 can only be merged with tiles of value 2")
	}

	t1.SetValue(2)
	t2.SetValue(2)

	tile, err = Merge(t1, t2)
	if err == nil {
		t.Error("Tiles of value 1 can only be merged with tiles of value 2")
	}
}

func TestTile_String(t *testing.T) {
	tile := Tile{3}
	tileString := `3`

	expected := tileString
	actual := fmt.Sprint(tile)
	if actual != expected {
		t.Errorf("New tile printed incorecctly. Got:\n%v\nWanted:\n%v\n", actual, expected)
	}
}

func TestIncompatibleTilesError_Error(t *testing.T) {
	t1 := Tile{3}
	t2 := Tile{3}
	err := IncompatibleTilesError{t1, t2}.Error()

	if err == "" {
		t.Error("Error does not display incompatible values")
	}
}
