package game

type Nothing struct {
	representation string
}

func NewNothing() *Nothing {
	return &Nothing{representation: "-"}
}

func (nothing *Nothing) Hit() {
	return
}
