package game

import "github.com/stretchr/testify/mock"

type Tile interface {
	Hit()
}

type MockedTile struct {
	mock.Mock
}

func (mockedTile *MockedTile) Hit() {
	mockedTile.Called()
}
