package main

import rover "github.com/zachatttack/clean-code/vehicle"

func main() {
	input := "MMRMMLM"
	// input := "MLLLLLLLLL"
  // input := "MMMRRRRR"
	// input := "MMMMMMMMMM"

	mars := rover.NewRover()
	mars.CommandRover(input)
	mars.GetPosition()
}
