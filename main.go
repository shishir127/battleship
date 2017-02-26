package main

import "fmt"

func main() {
	fmt.Println("Starting up...")
	// API design
	// fileInterface = &FileInterface{InputFileName: "foo", OutputFileName: "bar"}
	// fileInterface.CheckInputFile()
	fmt.Println("Reading input file...")
	// fileInterface.ReadInput()
	// parser = &BattlshipParser{}
	// parser.Parse(fileInterface)
	// battleShip = NewBattleship(parser)
	fmt.Println("Writing to output file...")
	// fileInterface.WriteOutput(battleShip.Simulate())
	fmt.Println("Done!")
}
