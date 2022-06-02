package main

import rover "myrover/rover"

func main() {
	input := "MMRMMLM"
	// input := "MLLLLLLLLL"
  // input := "MMMRRRRR"
	// input := "MMMMMMMMMM"

	mars := rover.NewRover()
	mars.CommandRover(input)
	mars.GetPosition()
}
