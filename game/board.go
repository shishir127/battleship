package game

import (
	"errors"
	"fmt"
)

type Board struct {
	grid map[string]Tile
	size int
}

func NewBoard(shipsLocation map[Location]*Ship, size int) (*Board, error) {
	grid := createBlankGrid(size)

	for location, ship := range shipsLocation {
		if _, ok := grid[location.Hash()]; ok {
			grid[location.Hash()] = ship
		} else {
			return &Board{}, errors.New(fmt.Sprintf("Ship location %s is outside board bounds", location.Hash()))
		}
	}

	return &Board{grid: grid, size: size}, nil
}

func (board *Board) HitMissiles(missiles []Missile) error {
	for _, missile := range missiles {
		if _, ok := board.grid[missile.Target.Hash()]; ok {
			board.grid[missile.Target.Hash()].Hit()
		} else {
			return errors.New(fmt.Sprintf("Missile target %s is not on the board", missile.Target.Hash()))
		}
	}
	return nil
}

func createBlankGrid(size int) map[string]Tile {
	grid := make(map[string]Tile)
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			location := Location{x_coordinate: i, y_coordinate: j}
			grid[location.Hash()] = NewNothing()
		}
	}
	return grid
}
