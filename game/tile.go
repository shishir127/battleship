package game

type Tile interface {
	Hit()
	Representation() string
}
