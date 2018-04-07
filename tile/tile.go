package tile

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
