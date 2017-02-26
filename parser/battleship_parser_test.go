package parser

import (
	"errors"
	"game"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	input := []string{
		"10",
		"2",
		"1:2,4:3",
		"3:7,5:9",
		"2",
		"5:2,9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{
		GridSize:      10,
		NumberOfShips: 2,
		Player1ShipPositions: []*game.Location{
			game.NewLocation(1, 2),
			game.NewLocation(4, 3),
		},
		Player2ShipPositions: []*game.Location{
			game.NewLocation(3, 7),
			game.NewLocation(5, 9),
		},
		NumberOfMissiles: 2,
		Player1sMissiles: []*game.Missile{
			game.NewMissile(5, 2),
			game.NewMissile(9, 6),
		},
		Player2sMissiles: []*game.Missile{
			game.NewMissile(5, 5),
			game.NewMissile(7, 6),
		},
	}

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.NoError(t, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenGridSizeIsMalformed(t *testing.T) {
	input := []string{
		"foo",
		"2",
		"1:2,4:3",
		"3:7,5:9",
		"2",
		"5:2,9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Grid size - strconv.ParseInt: parsing \"foo\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenNumberOfShipsIsMalformed(t *testing.T) {
	input := []string{
		"10",
		"foo",
		"1:2,4:3",
		"3:7,5:9",
		"2",
		"5:2,9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Number of ships - strconv.ParseInt: parsing \"foo\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenNumberOfMissilesIsMalformed(t *testing.T) {
	input := []string{
		"10",
		"2",
		"1:2,4:3",
		"3:7,5:9",
		"foo",
		"5:2,9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Number of missiles - strconv.ParseInt: parsing \"foo\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenPlayer1ShipLocationIsMalformed(t *testing.T) {
	input := []string{
		"10",
		"2",
		"1:2<,4:3",
		"3:7,5:9",
		"2",
		"5:2,9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Player 1 ship locations - strconv.ParseInt: parsing \"2<\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenPlayer2ShipLocationIsMalformed(t *testing.T) {
	input := []string{
		"10",
		"2",
		"1:2,4:3",
		"3:_7,5:9",
		"2",
		"5:2,9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Player 2 ship locations - strconv.ParseInt: parsing \"_7\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenPlayer2MissilesIsMalformed(t *testing.T) {
	input := []string{
		"10",
		"2",
		"1:2,4:3",
		"3:7,5:9",
		"2",
		"5:2,9:6",
		"5:!5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Player 2 missiles - strconv.ParseInt: parsing \"!5\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestParserWhenPlayer1MissilesIsMalformed(t *testing.T) {
	input := []string{
		"10",
		"2",
		"1:2,4:3",
		"3:7,5:9",
		"2",
		"5:2,@9:6",
		"5:5,7:6",
	}

	expectedOutput := &game.BattleshipSetup{}
	expectedError := errors.New("Player 1 missiles - strconv.ParseInt: parsing \"@9\": invalid syntax")

	parser := BattleShipParser{input: input}
	err := parser.Parse()
	assert.Equal(t, expectedError, err)
	battleShipSetup := parser.Tokens()
	assert.Equal(t, expectedOutput, battleShipSetup)
}

func TestNewBattleShipParser(t *testing.T) {
	testInput := make([]string, 1)
	parser := NewBattleShipParser(testInput)
	assert.Equal(t, reflect.TypeOf(&BattleShipParser{}), reflect.TypeOf(parser))
	assert.Equal(t, testInput, parser.input)
}
