package main

import "fmt"

// Player should be an interface for extensibility and testability?

// Position ...
type Position struct {
	X, Y int
}

// Player represents a player in the map with name, position and direction it is facing
type Player struct {
	Name      string
	Pos       Position
	Direction Direction
}

// NewPlayer creates a new player with appropriate properties
func NewPlayer(name string, x int, y int, d Direction) *Player {
	return &Player{
		name,
		Position{x, y},
		d,
	}
}

// Move moves the player in the direction set
// x coordinate is West -> East incrementally
// y coordinate is South -> North incrementally
func (p *Player) Move() error {
	switch p.Direction {
	case North:
		p.Pos.Y++
	case East:
		p.Pos.X++
	case South:
		p.Pos.Y--
	case West:
		p.Pos.X--
	default:
		return fmt.Errorf("cannot move in this direction: %v", p.Direction)
	}
	return nil
}

// Rotate changes the direction the player is facing
func (p *Player) Rotate(to string) {
	switch p.Direction {
	case North:
		switch to {
		case "LEFT":
			p.Direction = West
		case "RIGHT":
			p.Direction = East
		}
	case East:
		switch to {
		case "LEFT":
			p.Direction = North
		case "RIGHT":
			p.Direction = South
		}
	case South:
		switch to {
		case "LEFT":
			p.Direction = East
		case "RIGHT":
			p.Direction = West
		}
	case West:
		switch to {
		case "LEFT":
			p.Direction = South
		case "RIGHT":
			p.Direction = North
		}
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s: %d,%d,%s", p.Name, p.Pos.X, p.Pos.Y, p.Direction)
}

// Report prints the player information
func (p *Player) Report() {
	fmt.Println(p)
}
