package game

const (
	NOT_HIT = "-"
	HIT     = "X"
)

type Ship struct {
	representation string
}

func (ship *Ship) IsPresent() bool {
	return true
}

func (ship *Ship) Hit() *Ship {
	return &Ship{representation: HIT}
}

func NewShip() *Ship {
	return &Ship{representation: NOT_HIT}
}
