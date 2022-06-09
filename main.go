package main

import (
	"fmt"
	rover "github.com/zachatttack/clean-code/vehicle"
)

func main() {
	input := "MMM"
	// input := "MMRMMLM"
	// input := "MLLLLLLLLL"
  // input := "MMMRRRRR"
	// input := "MMMMMMMMMM"

	mars := rover.NewRover()
	mars.AddObstacle(0,3) 
  err := mars.CommandRover(input)
  if (err != nil){
    fmt.Println(err)
    return
  }
	mars.GetPosition()
}
