package main

import (
	"fmt"
	"game"
	"interfaces"
	"parser"
)

func main() {
	fmt.Println("Starting up...")
	fileInterface := &interfaces.FileInterface{InputFileName: "/Users/shishir/side_projects/battleship/input.txt", OutputFileName: "bar"}
	fmt.Println("Reading input file...")

	input, err := fileInterface.ReadInput()
	if err != nil {
		fmt.Print(err)
		return
	}

	battleShipParser := &parser.BattleShipParser{Input: input}
	err = battleShipParser.Parse()
	if err != nil {
		fmt.Print(err)
		return
	}

	battleShip, err := game.NewBattleship(battleShipParser.Tokens())
	err = battleShip.RunSimulation()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Writing to output file...")
	// fileInterface.WriteOutput(battleShip.TerminalState())
	fmt.Println("Done!")
}
