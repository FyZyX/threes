package tile_generator

import (
	. "Threes/tile"
	"math/rand"
)

type (
	Deck = []Value

	SimpleTileGenerator struct {
		deck []Value
	}

	BonusTileGenerator struct {
		maxValue       Value
		possibleValues []Value
	}
)

func newDeck() Deck {
	deck := []Value{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3}
	return shuffle(deck)
}

func shuffle(deck Deck) (shuffled Deck) {
	shuffled = make([]Value, len(deck))
	perm := rand.Perm(len(deck))

	for i, v := range perm {
		shuffled[v] = deck[i]
	}

	return shuffled
}

func NewSimpleTileGenerator() SimpleTileGenerator {
	return SimpleTileGenerator{newDeck()}
}

func (generator *SimpleTileGenerator) pop() (top Value) {
	if len(generator.deck) == 1 {
		top = generator.deck[0]
		generator.deck = newDeck()
		return top
	}

	top = generator.deck[0]
	generator.deck = generator.deck[1:]
	return top
}

func (generator *SimpleTileGenerator) Generate() (tile Tile) {
	tile.SetValue(generator.pop())
	return tile
}

func NewBonusTileGenerator(maxTileValue Value) BonusTileGenerator {
	generator := BonusTileGenerator{}
	generator.SetMaxValue(maxTileValue)
	generator.generatePossibleValues()
	return generator
}

func (generator BonusTileGenerator) generatePossibleValues() (possibleValues []Value) {
	value := generator.maxValue
	for value > 3 {
		possibleValues = append(possibleValues, value)
		value /= 2
	}

	return possibleValues
}

func (generator BonusTileGenerator) PossibleValues() (possibleValues []Value) {
	return generator.possibleValues
}

func (generator *BonusTileGenerator) SetMaxValue(maxTileValue Value) {
	generator.maxValue = maxTileValue / 8
}

func (generator BonusTileGenerator) ShouldGenerate() bool {
	return generator.maxValue > 0 && rand.Float32() < 1/21
}

func randomValue(values []Value) Value {
	return values[rand.Intn(len(values))]
}

func (generator BonusTileGenerator) Generate() (tile Tile) {
	tile.SetValue(randomValue(generator.generatePossibleValues()))
	return tile
}
