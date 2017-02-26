package game

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	boardSize := 2
	input := map[Location]*Ship{
		Location{x_coordinate: 1, y_coordinate: 2}: NewShip(),
	}
	expectedGrid := map[string]Tile{
		"0:0": NewNothing(),
		"0:1": NewNothing(),
		"0:2": NewNothing(),
		"1:0": NewNothing(),
		"1:1": NewNothing(),
		"1:2": NewShip(),
		"2:0": NewNothing(),
		"2:1": NewNothing(),
		"2:2": NewNothing(),
	}
	board, err := NewBoard(input, boardSize)
	assert.NoError(t, err)
	assert.Equal(t, boardSize, board.size)
	assert.Equal(t, 9, len(board.grid))
	assert.Equal(t, expectedGrid, board.grid)
}

func TestNewBoardShouldReturnAnErrorIfShipLocationIsOutsideBounds(t *testing.T) {
	boardSize := 2
	input := map[Location]*Ship{
		Location{x_coordinate: 3, y_coordinate: 2}: NewShip(),
	}
	board, err := NewBoard(input, boardSize)
	expectedError := errors.New("Ship location 3:2 is outside board bounds")
	assert.Equal(t, expectedError, err)
	assert.Equal(t, 0, board.size)
	assert.Equal(t, 0, len(board.grid))
}
