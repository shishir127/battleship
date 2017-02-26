package game

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestNewMissile(t *testing.T) {
	missile := NewMissile(1, 2)
	expectedMissile := &Missile{Target: Location{x_coordinate: 1, y_coordinate: 2}}
	assert.Equal(t, expectedMissile, missile)
}
