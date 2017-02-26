package game

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	boardSize := 2
	input := []Location{Location{x_coordinate: 1, y_coordinate: 2}}
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
	input := []Location{Location{x_coordinate: 3, y_coordinate: 2}}
	board, err := NewBoard(input, boardSize)
	expectedError := errors.New("Ship location 3:2 is outside board bounds")
	assert.Equal(t, expectedError, err)
	assert.Equal(t, 0, board.size)
	assert.Equal(t, 0, len(board.grid))
}

func TestHitMissilesShouldHitTheTileThatIsTheTargetOfTheMissile(t *testing.T) {
	mockedTile1 := &MockedTile{}
	mockedTile2 := &MockedTile{}
	boardSize := 1
	grid := map[string]Tile{
		"0:0": NewNothing(),
		"0:1": mockedTile1,
		"1:0": NewNothing(),
		"1:1": mockedTile2,
	}
	missiles := []Missile{
		Missile{Target: Location{x_coordinate: 0, y_coordinate: 1}},
		Missile{Target: Location{x_coordinate: 1, y_coordinate: 1}},
	}

	mockedTile1.On("Hit")
	mockedTile2.On("Hit")

	board := &Board{grid: grid, size: boardSize}
	err := board.HitMissiles(missiles)

	assert.NoError(t, err)
	mockedTile1.AssertExpectations(t)
	mockedTile2.AssertExpectations(t)
}

func TestHitMissilesShouldReturnAnErrorIfMissileTargetIsOutOfBounds(t *testing.T) {
	boardSize := 1
	grid := map[string]Tile{
		"0:0": NewNothing(),
		"0:1": NewNothing(),
		"1:0": NewNothing(),
		"1:1": NewNothing(),
	}
	missiles := []Missile{
		Missile{Target: Location{x_coordinate: 2, y_coordinate: 1}},
	}
	expectedError := errors.New("Missile target 2:1 is not on the board")

	board := &Board{grid: grid, size: boardSize}
	err := board.HitMissiles(missiles)

	assert.Equal(t, expectedError, err)
	board.HitMissiles(missiles)
}

func TestBoardDisplay(t *testing.T) {
	boardSize := 1
	grid := map[string]Tile{
		"0:0": NewNothing(),
		"0:1": NewNothing(),
		"1:0": NewNothing(),
		"1:1": NewNothing(),
	}
	board := &Board{grid: grid, size: boardSize}
	expectedOutput := []string{
		"-,-",
		"-,-",
	}

	output := board.Display()

	assert.Equal(t, expectedOutput, output)
}
