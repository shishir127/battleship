package game

import (
	"errors"
	"fmt"
)

type Battleship struct {
	player1Board    Board
	player2Board    Board
	player1Missiles []Missile
	player2Missiles []Missile
}

func NewBattleship(setup *BattleshipSetup) (*Battleship, error) {
	return nil, nil
}

func (battleship *Battleship) RunSimulation() error {
	err := battleship.player1Board.HitMissiles(battleship.player2Missiles)
	if err != nil {
		return errors.New(fmt.Sprintf("Player 2's turn - %s", err.Error()))
	}

	err = battleship.player2Board.HitMissiles(battleship.player1Missiles)
	if err != nil {
		return errors.New(fmt.Sprintf("Player 1's turn - %s", err.Error()))
	}

	return nil
}
