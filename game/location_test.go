package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashingWorks(t *testing.T) {
	location := Location{x_coordinate: 1, y_coordinate: 3}
	assert.Equal(t, "1:3", location.Hash())
}
