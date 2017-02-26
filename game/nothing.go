package game

type Nothing struct{}

func (object *Nothing) IsPresent() bool {
	return false
}

func (object *Nothing) Hit() *Nothing {
	return object
}
