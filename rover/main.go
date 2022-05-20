package main

import rover "myrover/rover"

func main() {
	input := "MMRMMLM"
	// input := "MMMMMMMMMM"

  mars := new(rover.Rover)
  mars.Init()
	mars.CommandRover(input)
  mars.GetPosition()
}
