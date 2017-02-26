package game

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunSimulation(t *testing.T) {
	board1 := Board{grid: createBlankGrid(2), size: 2}
	board2 := Board{grid: createBlankGrid(2), size: 2}
	missiles1 := []Missile{
		Missile{Target: Location{x_coordinate: 1, y_coordinate: 1}},
	}
	missiles2 := []Missile{
		Missile{Target: Location{x_coordinate: 0, y_coordinate: 1}},
	}
	battleship := &Battleship{
		player1Board:    board1,
		player2Board:    board2,
		player1Missiles: missiles1,
		player2Missiles: missiles2,
	}

	err := battleship.RunSimulation()
	assert.NoError(t, err)
}

func TestRunSimulationShouldHandeErrorsForPlayer2sTurn(t *testing.T) {
	board1 := Board{grid: createBlankGrid(2), size: 2}
	board2 := Board{grid: createBlankGrid(2), size: 2}
	missiles1 := []Missile{
		Missile{Target: Location{x_coordinate: 1, y_coordinate: 1}},
	}
	missiles2 := []Missile{
		Missile{Target: Location{x_coordinate: -1, y_coordinate: 1}},
	}
	battleship := &Battleship{
		player1Board:    board1,
		player2Board:    board2,
		player1Missiles: missiles1,
		player2Missiles: missiles2,
	}
	expectedError := errors.New("Player 2's turn - Missile target -1:1 is not on the board")

	err := battleship.RunSimulation()
	assert.Equal(t, expectedError, err)
}

func TestRunSimulationShouldHandeErrorsForPlayer1sTurn(t *testing.T) {
	board1 := Board{grid: createBlankGrid(2), size: 2}
	board2 := Board{grid: createBlankGrid(2), size: 2}
	missiles1 := []Missile{
		Missile{Target: Location{x_coordinate: 1, y_coordinate: 4}},
	}
	missiles2 := []Missile{
		Missile{Target: Location{x_coordinate: 0, y_coordinate: 1}},
	}
	battleship := &Battleship{
		player1Board:    board1,
		player2Board:    board2,
		player1Missiles: missiles1,
		player2Missiles: missiles2,
	}
	expectedError := errors.New("Player 1's turn - Missile target 1:4 is not on the board")

	err := battleship.RunSimulation()
	assert.Equal(t, expectedError, err)
}
