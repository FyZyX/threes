package tile

import "errors"

type Tile struct {
	value uint8
}

func (tile Tile) Value() (value uint8) {
	return 0
}

func (tile Tile) SetValue(value uint8) {
	tile.value = value
}

func (tile Tile) IsEmpty() bool {
	return true
}

func Merge(t1 Tile, t2 Tile) (tile Tile, err error) {
	if t1.value != t2.value {
		return tile, errors.New("incompatible tile values")
	}

	return Tile{t1.value + t2.value}, err
}
