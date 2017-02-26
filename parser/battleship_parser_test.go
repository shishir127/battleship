package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestParser(t *testing.T) {
// 	input := []string{
// 		"10",
// 		"2",
// 		"1:2,4:3",
// 		"3:7,5:9",
// 		"2",
// 		"5:2,9:6",
// 		"5:5,7:6",
// 	}

// 	parser = BattleShipParser{input: input}

// 	assert.Equal()
// }

func TestNewBattleShipParser(t *testing.T) {
	testInput := make([]string, 1)
	parser := NewBattleShipParser(testInput)
	assert.Equal(t, reflect.TypeOf(&BattleShipParser{}), reflect.TypeOf(parser))
	assert.Equal(t, testInput, parser.input)
}
