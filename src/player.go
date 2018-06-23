package main

// Player should be an interface for extensibility and testability?

// Player represents a player in the map with name, position and direction it is facing
type Player struct {
	Name      string
	Direction Direction
}

// NewPlayer creates a new player with appropriate properties
func NewPlayer(name string, d Direction) *Player {
	return &Player{
		name,
		d,
	}
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
