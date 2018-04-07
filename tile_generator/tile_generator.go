package tile_generator

import (
	. "Threes/tile"
)

type (
	SimpleTileGenerator struct{}
	BonusTileGenerator  struct {
		maxValue Value
	}
)

func (SimpleTileGenerator) Generate() (tile Tile) {
	tile.SetValue(1)

	return
}

func (generator *BonusTileGenerator) SetMaxValue(maxTileValue Value) {
	generator.maxValue = maxTileValue / 8
}

func (generator BonusTileGenerator) Generate() (tile Tile) {
	tile.SetValue(generator.maxValue)

	return
}
