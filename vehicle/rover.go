package vehicle

import (
	"fmt"
	"errors"
)

type direction int

const (
	north direction = iota
	west
	south
	east
)

func (d direction) String() string {
	return [...]string{"N", "W", "S", "E"}[d]
}

type grid [10][10]int

type Rover struct {
	column  int
	row     int
	heading direction
  grid grid
}


func NewRover() *Rover {
	return &Rover{
		row:     0,
		column:  0,
		heading: north,
    grid: grid{},
	}
}

func (r *Rover) AddObstacle(col int, row int){
   r.grid[col][row] = 1
}

func (r Rover) GetPosition() {
	fmt.Printf("%d:%d:%s \n", r.column, r.row, r.heading.String())
}

func (r Rover) ObstacleHit() error{
  return fmt.Errorf("%s:%d:%d:%s \n", "O", r.column, r.row, r.heading.String())
}

func (r *Rover) CommandRover(input string) error{
	for i := 0; i < len(input); i++ {
		command := string(input[i])
		if command == "M" {
      r.rowUpdate()
      r.colUpdate()
      err := r.checkForObstacles()
      if (err != nil){
        return r.ObstacleHit()
      }
		}
		if command == "R" || command == "L" {
			r.headingUpdate(command)
		}
	}
  return nil
}

func (r Rover)checkForObstacles() error{
  if (r.grid[r.column][r.row] == 1){
    return errors.New("ObstacleHit")
  }
  return nil
}

func (r *Rover) rowUpdate() {
	if r.heading == north {
		increaseCoordinate(&r.row)
	}
	if r.heading == south {
		decreaseCoordinate(&r.row)
	}
}

func (r *Rover) colUpdate() {
	if r.heading == east {
		increaseCoordinate(&r.column)
	}
	if r.heading == west {
		decreaseCoordinate(&r.column)
	}
}

func (r *Rover) headingUpdate(direction string) {
	if direction == "L" {
		r.leftCommanded()
	}
	if direction == "R" {
		r.rightCommanded()
	}
}

func (r *Rover) leftCommanded() {
	if r.heading == east {
		r.heading = north
		return
	}
	r.heading = r.heading + 1
}
func (r *Rover) rightCommanded() {
	if r.heading == north {
		r.heading = east
		return
	}
	r.heading = r.heading - 1
}

func increaseCoordinate(val *int) {
	*val = *val + 1
	if *val == 10 {
		*val = 0
	}
}

func decreaseCoordinate(val *int) {
	*val = *val - 1
	if *val == -1 {
		*val = 9
	}
}
