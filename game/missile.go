package game

type Missile struct {
	Target Location
}

func NewMissile(x_coordinate, y_coordinate int) *Missile {
	return &Missile{Target: Location{x_coordinate: x_coordinate, y_coordinate: y_coordinate}}
}
