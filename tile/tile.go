package tile

import (
	"fmt"
)

type (
	Value uint8

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

func Merge(t1 Tile, t2 Tile) (tile Tile, err error) {
	if t1.value < 3 && (t1.value == t2.value) {
		return tile, IncompatibleTilesError{t1, t2}
	}

	if t1.value+t2.value == 3 {
		return Tile{3}, err
	}

	if t1.value != t2.value {
		return tile, IncompatibleTilesError{t1, t2}
	}

	return Tile{t1.value + t2.value}, err
}

func (tile Tile) String() string {
	return fmt.Sprintf("%v", tile.value)
}

func (e IncompatibleTilesError) Error() (err string) {
	return fmt.Sprintf("Incompatible tile values: %v and %v", e.t1.value, e.t2.value)
}
