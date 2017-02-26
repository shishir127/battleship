package game

type Object interface {
	IsPresent() bool
	Hit() *Object
}
