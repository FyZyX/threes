package board

import (
	. "Threes/tile"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type (
	Direction uint8

	Row   [4]Tile
	Board [4]Row

	Index struct {
		Row    int
		Column int
	}

	TileGenerator interface {
		Generate() Tile
	}
)

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (board *Board) AddTileAt(tile Tile, index Index) {
	board[index.Row][index.Column] = tile
}

func (board *Board) TileAt(index Index) Tile {
	return board[index.Row][index.Column]
}

func (index Index) rotate(direction Direction) Index {
	switch direction {
	case Left: /* clockwise */
		return Index{
			Row:    3 - index.Column,
			Column: index.Row,
		}
	case Right: /* counter-clockwise */
		return Index{
			Row:    index.Column,
			Column: 3 - index.Row,
		}
	case Down: /* equivalent to two LEFT or RIGHT rotations */
		return Index{
			Row:    3 - index.Row,
			Column: 3 - index.Column,
		}
	default:
		return index
	}
}

func (board *Board) Slide(direction Direction) (indices []Index, err error) {
	/* This algorithm is based on the following observations:
	 * swipe LEFT is the same as swipe UP after a clockwise rotation
	 * swipe RIGHT is the same as swipe UP after a counter-clockwise rotation
	 * swipe DOWN is the same as swipe UP after two LEFT or RIGHT rotations
	 *
	 * All comments refer to swipe UP, though the indices are rotated in the specified direction
	 */

	var moved [len(board)]bool

	// since the first row is at the bounds, loop over only the latter three rows
	for i, row := range board[1:] {
		for j := range row {
			// OBO: index i starts at 0 in this looping construct,
			// yet we are starting with the row at index 1.
			// Row 1 is first merged with row 0, then row 2 with row 1, and so on.
			source := Index{i + 1, j}.rotate(direction)
			destination := Index{i, j}.rotate(direction)

			tile, err := Merge(board.TileAt(source), board.TileAt(destination))
			if err != nil {
				// tiles cannot be merged, move on
				continue
			}

			// place the merged tile in the destination index
			board.AddTileAt(tile, destination)
			// place an empty tile in the source index
			board.AddTileAt(Tile{}, source)

			// mark column as moved
			moved[j] = true
		}
	}

	// return all possible indices where a new tile can be generated
	for i, wasMoved := range moved {
		if wasMoved {
			indices = append(indices, Index{3, i}.rotate(direction))
		}
	}

	if len(indices) == 0 {
		err = errors.New(fmt.Sprintf("cannot slide board %s", direction))
		return
	}

	return
}

func (board Board) Score() (score float64) {
	for _, row := range board {
		for _, tile := range row {
			if tile.IsEmpty() {
				continue
			}
			score += math.Pow(3, float64(tile.Merges()))
		}
	}
	return score
}

func (row Row) String() string {
	return fmt.Sprintf(`| %5v | %5v | %5v | %5v |`, row[0], row[1], row[2], row[3])
}

func (board Board) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v", board[0], board[1], board[2], board[3])
}
