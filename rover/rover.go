package main

import (
	"fmt"
)

type Rover struct {
	column  int
	row     int
	heading string
}

func main() {
	input := "MMRMMLM"
	// input := "MMMMMMMMMM"

  mars := new(Rover)
  mars.Init()
	mars.CommandRover(input)
  mars.GetPosition()
}

func (r *Rover) Init(){
  r.row = 0
  r.column = 0
  r.heading = "N"
}

func (r *Rover) GetPosition(){
  fmt.Printf("%d:%d:%s \n",r.column,r.row,r.heading)
}

func (r *Rover) CommandRover(input string) {
	for i := 0; i < len(input); i++ {
		command := string(input[i])
		if command == "M" {
			r.rowUpdate()
			r.colUpdate()
		}
		if command == "R" || command == "L" {
			r.headingUpdate(command)
		}
	}
}

func (r *Rover) rowUpdate() {
	if r.heading == "N" {
		r.row = r.row + 1
    if r.row == 10 {
      r.row = 0
    }
	}
	if r.heading == "S" {
		r.row = r.row - 1
    if r.row == -1{
      r.row = 9
    }
	}
}

func (r *Rover) colUpdate(){
	if r.heading == "E" {
		r.column = r.column + 1
    if r.column == 10 {
      r.column = 0
    }
	}
	if r.heading == "W" {
		r.column = r.column - 1
    if r.column == -1{
      r.column = 9
    }
	}
}

func (r *Rover) headingUpdate(direction string) {
	if direction == "L" {
		if r.heading == "N" {
			r.heading = "W"
		} else if (r.heading == "W") {
			r.heading = "S"
		} else if (r.heading == "S") {
			r.heading = "E"
		} else if (r.heading == "E") {
			r.heading = "N"
		}
	}
	if direction == "R" {
		if (r.heading == "N") {
			r.heading = "E"
		} else if (r.heading == "W") {
			r.heading = "N"
		} else if (r.heading == "S") {
			r.heading = "W"
		} else if (r.heading == "E") {
			r.heading = "S"
		}
	}
}
