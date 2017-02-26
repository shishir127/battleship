package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHitChangeStateToHit(t *testing.T) {
	ship := Ship{representation: NOT_HIT}
	ship.Hit()
	assert.Equal(t, HIT, ship.representation)
}

func TestNewShipShouldReturnAShipThatIsNotHit(t *testing.T) {
	nothing := NewNothing()
	assert.Equal(t, NOT_HIT, nothing.representation)
}
