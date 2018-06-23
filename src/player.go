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
	Direction string
}

// NewPlayer creates a new player with appropriate properties
func NewPlayer(name string, x int, y int, direction string) *Player {
	return &Player{
		name,
		Position{x, y},
		direction,
	}
}

// Move moves the player in the direction set
// x coordinate is west -> east incrementally
// y coordinate is south -> north incrementally
func (p *Player) Move() error {
	switch p.Direction {
	case "NORTH":
		p.Pos.Y++
	case "EAST":
		p.Pos.X++
	case "SOUTH":
		p.Pos.Y--
	case "WEST":
		p.Pos.X--
	default:
		return fmt.Errorf("cannot move in this direction: %v", p.Direction)
	}
	return nil
}

// Rotate changes the direction the player is facing
func (p *Player) Rotate(to string) {
	switch p.Direction {
	case "NORTH":
		switch to {
		case "LEFT":
			p.Direction = "WEST"
		case "RIGHT":
			p.Direction = "EAST"
		}
	case "EAST":
		switch to {
		case "LEFT":
			p.Direction = "NORTH"
		case "RIGHT":
			p.Direction = "SOUTH"
		}
	case "SOUTH":
		switch to {
		case "LEFT":
			p.Direction = "EAST"
		case "RIGHT":
			p.Direction = "WEST"
		}
	case "WEST":
		switch to {
		case "LEFT":
			p.Direction = "SOUTH"
		case "RIGHT":
			p.Direction = "NORTH"
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
