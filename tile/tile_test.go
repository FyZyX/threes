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
	source := Tile{3}
	destination := Tile{3}

	tile, err := Merge(source, destination)
	if err != nil {
		t.Error("Tiles of equal value were not merged")
	}

	if tile.value != source.value+destination.value {
		t.Errorf("Unepexted merge value: wanted %v, but got %v", source.value+destination.value, tile.value)
	}

	source.SetValue(6)

	tile, err = Merge(source, destination)
	if err == nil {
		t.Error("Tiles of unequal value cannot be merged")
	}

	source.SetValue(1)
	destination.SetValue(2)

	tile, err = Merge(source, destination)
	if err != nil {
		t.Error("Tile with value 1 not merged with tile of value 2")
	}

	if tile.value != 3 {
		t.Errorf("Unepexted merge value: wanted %v, but got %v", 3, tile.value)
	}

	source.SetValue(1)
	destination.SetValue(1)

	tile, err = Merge(source, destination)
	if err == nil {
		t.Error("Tiles of value 1 can only be merged with tiles of value 2")
	}

	source.SetValue(2)
	destination.SetValue(2)

	tile, err = Merge(source, destination)
	if err == nil {
		t.Error("Tiles of value 1 can only be merged with tiles of value 2")
	}

	source.SetValue(2)
	destination.SetValue(0)

	tile, err = Merge(source, destination)
	if err != nil {
		t.Error("Tile not merged with empty tile")
	}

	source.SetValue(0)
	destination.SetValue(1)

	tile, err = Merge(source, destination)
	fmt.Println(err)
	if err == nil {
		t.Error("Empty source tile cannot be merged")
	}
}

func TestTile_Merges(t *testing.T) {
	tile := Tile{}
	tile.SetValue(1)
	if tile.Merges() != 0 {
		t.Errorf("Incorrect number of merges for tile with value %v", tile.value)
	}

	tile.SetValue(2)
	if tile.Merges() != 0 {
		t.Errorf("Incorrect number of merges for tile with value %v", tile.value)
	}

	tile.SetValue(3)
	if tile.Merges() != 1 {
		t.Errorf("Incorrect number of merges for tile with value %v", tile.value)
	}

	tile.SetValue(6)
	if tile.Merges() != 2 {
		t.Errorf("Incorrect number of merges for tile with value %v", tile.value)
	}

	tile.SetValue(48)
	if tile.Merges() != 5 {
		t.Errorf("Incorrect number of merges for tile with value %v", tile.value)
	}

	tile.SetValue(384)
	if tile.Merges() != 8 {
		t.Errorf("Incorrect number of merges for tile with value %v", tile.value)
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
