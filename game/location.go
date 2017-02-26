package game

import (
	"fmt"
)

type Location struct {
	x_coordinate int
	y_coordinate int
}

func NewLocation(x_coordinate, y_coordinate int) *Location {
	return &Location{x_coordinate: x_coordinate, y_coordinate: y_coordinate}
}

func (location *Location) Hash() string {
	return fmt.Sprintf("%d:%d", location.x_coordinate, location.y_coordinate)
}
