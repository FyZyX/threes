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
		maxValue Value
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

func (stg *SimpleTileGenerator) pop() (top Value) {
	if len(stg.deck) == 1 {
		top = stg.deck[0]
		stg.deck = newDeck()
		return top
	}

	top = stg.deck[0]
	stg.deck = stg.deck[1:]
	return top
}

func (stg *SimpleTileGenerator) Generate() (tile Tile) {
	tile.SetValue(stg.pop())
	return tile
}

func (generator *BonusTileGenerator) SetMaxValue(maxTileValue Value) {
	generator.maxValue = maxTileValue / 8
}

func (generator BonusTileGenerator) Generate() (tile Tile) {
	tile.SetValue(generator.maxValue)
	return tile
}
