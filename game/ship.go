package game

const (
	NOT_HIT = "-"
	HIT     = "X"
)

type Ship struct {
	representation string
}

func (ship *Ship) Hit() {
	ship.representation = HIT
}

func NewShip() *Ship {
	return &Ship{representation: NOT_HIT}
}
