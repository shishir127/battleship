package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPresentShouldAlwaysReturnTrue(t *testing.T) {
	ship := Ship{}
	assert.True(t, ship.IsPresent())
}

func TestHitShouldReturnShipWithHitRepresentation(t *testing.T) {
	ship := Ship{}
	expectedShip := &Ship{representation: HIT}
	assert.Equal(t, expectedShip, ship.Hit())
}

func TestNewShipShouldReturnAShipThatIsNotHit(t *testing.T) {
	nothing := NewNothing()
	assert.Equal(t, NOT_HIT, nothing.representation)
}
