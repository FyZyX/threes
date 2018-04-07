package board

import (
	. "Threes/tile"
)

type (
	Board [][]Tile

	Index struct {
		Row    uint8
		Column uint8
	}

	TileGenerator interface {
		Generate() Tile
	}
)

func NewBoard() Board {
	board := make([][]Tile, 4, 4)
	for row := range board {
		board[row] = make([]Tile, 4, 4)
	}

	return board
}

func (board Board) AddTile(tile Tile, index Index) {
	board[index.Row][index.Column] = tile
}
