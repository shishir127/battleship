package main

import (
	"fmt"
	"interfaces"
)

func main() {
	fmt.Println("Starting up...")
	// API design
	fileInterface := &interfaces.FileInterface{InputFileName: "/Users/shishir/side_projects/battleship/input.txt", OutputFileName: "bar"}
	fmt.Println("Reading input file...")
	input, err := fileInterface.ReadInput()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(len(input))
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
	// parser = &BattlshipParser{}
	// parser.Parse(fileInterface)
	// battleShip = NewBattleship(parser)
	fmt.Println("Writing to output file...")
	// fileInterface.WriteOutput(battleShip.Simulate())
	fmt.Println("Done!")
}
