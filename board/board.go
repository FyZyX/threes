package board

import (
	. "Threes/tile"
	"fmt"
)

type (
	row   [4]Tile
	Board [4]row

	Index struct {
		Row    uint8
		Column uint8
	}

	TileGenerator interface {
		Generate() Tile
	}
)

func (board *Board) AddTile(tile Tile, index Index) {
	board[index.Row][index.Column] = tile
}

func (row row) String() string {
	return fmt.Sprintf(`| %5v | %5v | %5v | %5v |`, row[0], row[1], row[2], row[3])
}

func (board Board) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v", board[0], board[1], board[2], board[3])
}
