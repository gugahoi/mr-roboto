package main

import (
	"fmt"
)

// Map contains the players in the game
type Map struct {
	Players [6][6]*Player
}

// Map Layout (x,y)
// x coordinate is west -> east incrementally
// y coordinate is south -> north incrementally
//                         NORTH
//   | 0,N-1 | 1,N-1 | 2,N-1 | 3,N-1 | 4,N-1 | N-1,N-1 |
//	 |   .	 |   .   |   .   |   .   |   .   |   .     |
// W |   .	 |   .   |   .   |   .   |   .   |   .     | E
// E |   .	 |   .   |   .   |   .   |   .   |   .     | A
// S |  0,4  |  1,4  |  2,4  |  3,4  |  4,4  |  5,4    | S
// T |  0,3  |  1,3  |  2,3  |  3,3  |  4,3  |  5,3    | T
//   |  0,2  |  1,2  |  2,2  |  3,2  |  4,2  |  5,2    |
//   |  0,1  |  1,1  |  2,1  |  3,1  |  4,1  |  5,1    |
//   |  0,0  |  1,0  |  2,0  |  3,0  |  4,0  |  5,0    |
//                         SOUTH

// NewMap creates a new map of given size (square map)
func NewMap() *Map {
	return &Map{
		Players: [6][6]*Player{},
	}
}

// AddPlayer adds a new player to the map, unless a player of same name already exists
// or the chosen location is already taken
func (m *Map) AddPlayer(p *Player, x, y int) error {
	if p, _, _ := m.FindPlayerByName(p.Name); p != nil {
		return fmt.Errorf("player %v already in map, skipping", p)
	}

	if m.Players[x][y] != nil {
		return fmt.Errorf("location not empty: (%v,%v) %v", x, y, p.Name)
	}

	m.Players[x][y] = p
	return nil
}

// FindPlayerByName returns the player if the player with the given name
// has already been placed in the map
func (m *Map) FindPlayerByName(name string) (*Player, int, int) {
	for x, row := range m.Players {
		for y, player := range row {
			if player != nil && player.Name == name {
				return player, x, y
			}
		}
	}
	return nil, 0, 0
}

func (m *Map) String() string {
	var s string
	for x, row := range m.Players {
		for y, player := range row {
			if player != nil {
				s += fmt.Sprintf("%s: %d,%d,%s\n", player.Name, x, y, player.Direction)
			}
		}
	}
	return s
}

// Run runs a given command from the current state of the map
func (m *Map) Run(c Command) {
	switch c.Action {
	case "PLACE":
		x, y, d := ParseArgs(c.Args)
		m.AddPlayer(NewPlayer(c.Name, d), x, y)
	case "REPORT":
		m.Report(c.Name)
	case "MOVE":
		// m.FindPlayerByName(c.Name).Move()
	case "LEFT", "RIGHT":
		p, _, _ := m.FindPlayerByName(c.Name)
		p.Rotate(c.Action)
	}
}

// Report ...
func (m *Map) Report(name string) {
	p, x, y := m.FindPlayerByName(name)
	if p != nil {
		fmt.Printf("%s: %d,%d,%s\n", p.Name, x, y, p.Direction)
	}
}
