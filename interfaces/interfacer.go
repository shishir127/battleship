package interfaces

type Interfacer interface {
	ReadInput(i interface{}) ([]string, error)
	WriteOutput(i interface{}) ([]string, error)
}
