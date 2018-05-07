package tile_generator

import (
	. "Threes/tile"
	"testing"
)

func TestSimpleTileGenerator_Generate(t *testing.T) {
	generator := NewSimpleTileGenerator()
	tile := generator.Generate()

	if len(generator.deck) != 11 {
		t.Error("Generated tile not removed from deck")
	}

	if tile.IsEmpty() {
		t.Error("Generated tiles cannot be empty")
	}

	for i := 0; i < 11; i++ {
		generator.Generate()
	}

	if len(generator.deck) != 12 {
		t.Error("New deck not created on last pop")
	}
}

func TestNewBonusTileGenerator(t *testing.T) {
	var maxTileValue Value = 384

	generator := NewBonusTileGenerator(maxTileValue)

	if generator.maxValue != maxTileValue/8 {
		t.Error("Max value should be 1/8 of the maximum tile value")
	}
}

func TestBonusTileGenerator_PossibleValues(t *testing.T) {
	generator := NewBonusTileGenerator(48)
	values := generator.PossibleValues()

	expected := [1]Value{6}
	if len(values) != 1 || values[0] != 6 {
		t.Errorf("expected %v, got %v instead", expected, values)
	}
}

func TestBonusTileGenerator_SetMaxValue(t *testing.T) {
	var generator BonusTileGenerator
	var maxTileValue Value = 192

	generator.SetMaxValue(maxTileValue)

	if generator.maxValue != 24 {
		t.Error("Max value should be 1/8 of the maximum tile value")
	}
}

func TestBonusTileGenerator_Generate(t *testing.T) {
	generator := NewBonusTileGenerator(48)
	tile := generator.Generate()

	if tile.IsEmpty() {
		t.Error("Generated tiles cannot be empty")
	}

	if tile.Value() > generator.maxValue {
		t.Errorf("Tile value is %v, but it cannot exceed %v", tile.Value(), generator.maxValue)
	}
}
