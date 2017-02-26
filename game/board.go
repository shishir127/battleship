package game

type Board struct {
	grid map[string]Tile
	size int
}

func NewBoard(shipsLocation map[Location]*Ship, size int) (*Board, error) {
	grid := make(map[string]Tile)
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			location := Location{x_coordinate: i, y_coordinate: j}
			grid[location.Hash()] = NewNothing()
		}
	}

	for location, ship := range shipsLocation {
		grid[location.Hash()] = ship
	}

	return &Board{grid: grid, size: size}, nil
}
