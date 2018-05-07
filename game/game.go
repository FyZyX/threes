package game

import (
	. "Threes/board"
	. "Threes/tile"
	. "Threes/tile_generator"
	"fmt"
	"math/rand"
)

type (
	Game struct {
		board        Board
		nextTile     Tile
		maxTileValue Value
		IsOver       bool
	}
)

var (
	generator      = NewSimpleTileGenerator()
	bonusGenerator = NewBonusTileGenerator(0)
)

func NewGame() Game {
	var game Game

	var tilesAdded int
	for tilesAdded < 8 {
		if index := NewRandomIndex(); game.board.TileAt(index).IsEmpty() {
			tile := generator.Generate()
			game.board.AddTileAt(tile, index)
			tilesAdded++
		}
	}

	game.nextTile = game.generateNextTile()
	game.updateMaxTileValue()

	return game
}

func randomIndex(indices []Index) Index {
	return indices[rand.Intn(len(indices))]
}

func (game *Game) generateNextTile() Tile {
	var tile Tile

	if bonusGenerator.ShouldGenerate() {
		tile = bonusGenerator.Generate()
	} else {
		tile = generator.Generate()
	}

	return tile
}

func (game *Game) addNextTileAt(index Index) {
	game.board.AddTileAt(game.nextTile, index)
}

func (game *Game) Swipe(direction Direction) {
	indices, err := game.board.Slide(direction)
	if err != nil {
		fmt.Sprintf("cannot swipe board %s", direction)
	} else {
		game.addNextTileAt(randomIndex(indices))
		game.updateMaxTileValue()
	}
}

func (game *Game) updateMaxTileValue() Value {
	var maxValue Value
	for _, row := range game.board {
		for _, tile := range row {
			if value := tile.Value(); value > maxValue {
				maxValue = value
			}
		}
	}

	bonusGenerator.SetMaxValue(maxValue)

	return maxValue
}

func (game Game) BoardScore() float64 {
	return game.board.Score()
}

func (game *Game) BestMove() Direction {
	var possibleDirections []Direction

	currentScore := game.board.Score()
	bestScore := currentScore
	var bestDirection Direction
	directions := []Direction{Up, Down, Left, Right}
	for _, direction := range directions {
		game := *game
		indices, _ := game.board.Slide(direction)
		score := game.board.Score()
		//fmt.Printf("Checking %s, score is %v\n", direction, score)
		if len(indices) > 0 {
			possibleDirections = append(possibleDirections, direction)

			if score > bestScore {
				//fmt.Printf("%s is better\n", direction)
				bestScore = score
				bestDirection = direction
			}
		}
	}

	if bestScore == currentScore {
		if len(possibleDirections) == 0 {
			game.IsOver = true
		} else {
			bestDirection = possibleDirections[rand.Intn(len(possibleDirections))]
		}
	}

	return bestDirection
}

func (game Game) String() string {
	return game.board.String()
}
