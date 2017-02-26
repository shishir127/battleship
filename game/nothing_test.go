package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNothingShouldDoNothinWhenHitIsCalled(t *testing.T) {
	nothing := &Nothing{}
	previousState := nothing.representation
	nothing.Hit()
	assert.Equal(t, previousState, nothing.representation)
}

func TestNewNothingShouldReturnAnInstanceOfNothing(t *testing.T) {
	nothing := NewNothing()
	assert.Equal(t, "-", nothing.representation)
}
