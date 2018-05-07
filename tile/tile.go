package tile

import (
	"errors"
	"fmt"
	"math"
)

type (
	Value uint16

	Tile struct {
		value Value
	}

	IncompatibleTilesError struct {
		t1 Tile
		t2 Tile
	}
)

func (tile Tile) Value() (value Value) {
	return tile.value
}

func (tile *Tile) SetValue(value Value) {
	tile.value = value
}

func (tile Tile) IsEmpty() bool {
	return tile.value == 0
}

func Merge(source Tile, destination Tile) (tile Tile, err error) {
	// Any tile can be merged with an empty tile.
	if source.value == 0 {
		return tile, errors.New("source tile cannot be empty")
	}

	// Any tile can be merged with an empty tile.
	if source.value != 0 && destination.value == 0 {
		return Tile{source.value + destination.value}, err
	}

	// 1s cannot be merged with 1s.
	// 2s cannot be merged with 2s.
	if source.value < 3 && (source.value == destination.value) {
		return tile, IncompatibleTilesError{source, destination}
	}

	// 1s can be merged with 2s.
	if source.value+destination.value == 3 {
		return Tile{3}, err
	}

	// Tiles with unequal values cannot be merged.
	if source.value != destination.value {
		return tile, IncompatibleTilesError{source, destination}
	}

	return Tile{source.value + destination.value}, err
}

func (tile Tile) Merges() int {
	if tile.value < 3 {
		return 0
	}

	base := float64(tile.value / 3)
	return int(math.Log2(base)) + 1
}

func (tile Tile) String() string {
	return fmt.Sprintf("%v", tile.value)
}

func (e IncompatibleTilesError) Error() (err string) {
	return fmt.Sprintf("Incompatible tile values: %v and %v", e.t1.value, e.t2.value)
}
