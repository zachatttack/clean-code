package main

import (
	"fmt"
)

type rover struct {
	column  int
	row     int
	heading string
}

func main() {
	// input := "MMRMMLM"
	input := "MMRMMLM"

	mars := rover{
		row:     0,
		column:  0,
		heading: "N",
	}

	parseInput(input, &mars)
	fmt.Println(mars)
}

func parseInput(input string, mars *rover) {
	for i := 0; i < len(input); i++ {
		command := string(input[i])
		if command == "M" {
			rowUpdate(mars)
			colUpdate(mars)
		}
		if command == "R" || command == "L" {
			headingUpdate(mars, command)
		}
	}
}

func rowUpdate(mars *rover) {
	if mars.heading == "N" {
		mars.row = mars.row + 1
	}
	if mars.heading == "S" {
		mars.row = mars.row - 1
	}
}

func colUpdate(mars *rover) {
	if mars.heading == "E" {
		mars.column = mars.column + 1
	}
	if mars.heading == "W" {
		mars.column = mars.column - 1
	}
}

func headingUpdate(mars *rover, direction string) {
	if direction == "L" {
		if mars.heading == "N" {
			mars.heading = "W"
		} else if (mars.heading == "W") {
			mars.heading = "S"
		} else if (mars.heading == "S") {
			mars.heading = "E"
		} else if (mars.heading == "E") {
			mars.heading = "N"
		}
	}
	if direction == "R" {
		if (mars.heading == "N") {
			mars.heading = "E"
		} else if (mars.heading == "W") {
			mars.heading = "N"
		} else if (mars.heading == "S") {
			mars.heading = "W"
		} else if (mars.heading == "E") {
			mars.heading = "S"
		}
	}
}
