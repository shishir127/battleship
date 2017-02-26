package game

import "github.com/stretchr/testify/mock"

type Tile interface {
	Hit()
	Representation() string
}

type MockedTile struct {
	mock.Mock
}

func (mockedTile *MockedTile) Hit() {
	mockedTile.Called()
}

func (mockedTile *MockedTile) Representation() string {
	args := mockedTile.Called()
	return args.String(0)
}
