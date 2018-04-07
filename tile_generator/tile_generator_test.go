package tile_generator

import (
	. "Threes/tile"
	"testing"
)

func TestSimpleTileGenerator_Generate(t *testing.T) {
	generator := SimpleTileGenerator{}
	tile := generator.Generate()

	if tile.IsEmpty() {
		t.Error("Generated tiles cannot be empty")
	}
}

func TestBonusTileGenerator_SetMaxValue(t *testing.T) {
	var maxTileValue Value = 48

	var generator BonusTileGenerator
	generator.SetMaxValue(maxTileValue)

	if generator.maxValue != maxTileValue/8 {
		t.Error("Max value should be 1/8 of the maximum tile value")
	}
}

func TestBonusTileGenerator_Generate(t *testing.T) {
	var generator BonusTileGenerator
	generator.SetMaxValue(48)
	tile := generator.Generate()

	if tile.IsEmpty() {
		t.Error("Generated tiles cannot be empty")
	}

	if tile.Value() > generator.maxValue {
		t.Errorf("Tile value is %v, but it cannot exceed %v", tile.Value(), generator.maxValue)
	}
}
