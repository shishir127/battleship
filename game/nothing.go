package game

type Nothing struct {
	representation string
}

func NewNothing() *Nothing {
	return &Nothing{representation: "-"}
}

func (object *Nothing) IsPresent() bool {
	return false
}

func (object *Nothing) Hit() *Nothing {
	return object
}
