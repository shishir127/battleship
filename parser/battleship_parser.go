package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"game"
)

type BattleShipParser struct {
	Input []string
	setup game.BattleshipSetup
}

func NewBattleShipParser(stringInput []string) *BattleShipParser {
	return &BattleShipParser{Input: stringInput}
}

func (parser *BattleShipParser) Parse() error {
	var err error
	gameSetup := game.BattleshipSetup{}
	gameSetup.GridSize, err = strconv.Atoi(parser.Input[0])
	if err != nil {
		return errors.New(fmt.Sprintf("Grid size - %s", err.Error()))
	}

	gameSetup.NumberOfShips, err = strconv.Atoi(parser.Input[1])
	if err != nil {
		return errors.New(fmt.Sprintf("Number of ships - %s", err.Error()))
	}

	gameSetup.NumberOfMissiles, err = strconv.Atoi(parser.Input[4])
	if err != nil {
		return errors.New(fmt.Sprintf("Number of missiles - %s", err.Error()))
	}

	gameSetup.Player1ShipPositions, err = parseLocations(parser.Input[2])
	if err != nil {
		return errors.New(fmt.Sprintf("Player 1 ship locations - %s", err.Error()))
	}

	gameSetup.Player2ShipPositions, err = parseLocations(parser.Input[3])
	if err != nil {
		return errors.New(fmt.Sprintf("Player 2 ship locations - %s", err.Error()))
	}

	gameSetup.Player1sMissiles, err = parseMissileTargets(parser.Input[5])
	if err != nil {
		return errors.New(fmt.Sprintf("Player 1 missiles - %s", err.Error()))
	}

	gameSetup.Player2sMissiles, err = parseMissileTargets(parser.Input[6])
	if err != nil {
		return errors.New(fmt.Sprintf("Player 2 missiles - %s", err.Error()))
	}

	parser.setup = gameSetup
	return nil
}

func parseLocations(location string) ([]*game.Location, error) {
	var locationObjects []*game.Location
	locationInfo, err := parseLocationInfo(location)
	if err != nil {
		return locationObjects, err
	}

	for _, rawLocation := range locationInfo {
		locationObjects = append(locationObjects, game.NewLocation(rawLocation[0], rawLocation[1]))
	}
	return locationObjects, nil
}

func parseMissileTargets(location string) ([]*game.Missile, error) {
	var missiles []*game.Missile
	locationInfo, err := parseLocationInfo(location)
	if err != nil {
		return missiles, err
	}
	for _, rawLocation := range locationInfo {
		missiles = append(missiles, game.NewMissile(rawLocation[0], rawLocation[1]))
	}
	return missiles, nil
}

func parseLocationInfo(locations string) ([][]int, error) {
	locationList := strings.Split(locations, ",")
	var location [][]int
	for _, locationInfo := range locationList {
		ordinates := strings.Split(locationInfo, ":")
		x_coordinate, err := strconv.Atoi(ordinates[0])
		if err != nil {
			return location, err
		}

		y_coordinate, err := strconv.Atoi(ordinates[1])
		if err != nil {
			return location, err
		}
		location = append(location, []int{x_coordinate, y_coordinate})
	}
	return location, nil
}

func (parser *BattleShipParser) Tokens() *game.BattleshipSetup {
	return &parser.setup
}
