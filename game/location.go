package game

import (
	"fmt"
)

type Location struct {
	x_coordinate int
	y_coordinate int
}

func (location *Location) Hash() string {
	return fmt.Sprintf("%d:%d", location.x_coordinate, location.y_coordinate)
}
