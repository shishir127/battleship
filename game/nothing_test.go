package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNothingShouldAlwaysReturnFalse(t *testing.T) {
	nothing := Nothing{}
	assert.False(t, nothing.IsPresent())
}

func TestNothingShouldReturnSelfWhenHitIsCalled(t *testing.T) {
	nothing := &Nothing{}
	assert.Equal(t, nothing, nothing.Hit())
}

func TestNewNothingShouldReturnAnInstanceOfNothing(t *testing.T) {
	nothing := NewNothing()
	assert.Equal(t, "-", nothing.representation)
}
