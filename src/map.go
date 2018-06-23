package main

import (
	"fmt"
	"log"
)

// Map contains the players in the game
type Map struct {
	Players []*Player
}

// Map Layout (x,y)
// x coordinate is west -> east incrementally
// y coordinate is south -> north incrementally
//                NORTH
//   | 0,N-1 | 1,N-1 | 2,N-1 | 3,N-1 | 4,N-1 | N-1,N-1 |
//	 |   .	 |   .   |   .   |   .   |   .   |   .     |
//	 |   .	 |   .   |   .   |   .   |   .   |   .     |
//	 |   .	 |   .   |   .   |   .   |   .   |   .     |
// W |  0,4  |  1,4  |  2,4  |  3,4  |  4,4  |  5,4    | E
// E |  0,3  |  1,3  |  2,3  |  3,3  |  4,3  |  5,3    | A
// S |  0,2  |  1,2  |  2,2  |  3,2  |  4,2  |  5,2    | S
// T |  0,1  |  1,1  |  2,1  |  3,1  |  4,1  |  5,1    | T
//   |  0,0  |  1,0  |  2,0  |  3,0  |  4,0  |  5,0    |
//                SOUTH

// NewMap creates a new map of given size (square map)
func NewMap() *Map {
	return &Map{
		Players: []*Player{},
	}
}

// AddPlayer adds a new player to the map, unless a player of same name already exists
func (m *Map) AddPlayer(p *Player) {
	if m.FindPlayerByName(p.Name) != nil {
		log.Printf("Player %v already in map, skipping\n", p)
		return
	}

	m.Players = append(m.Players, p)
}

// FindPlayerByName returns the player if the player with the given name
// has already been placed in the map
func (m *Map) FindPlayerByName(name string) *Player {
	for _, player := range m.Players {
		if player.Name == name {
			return player
		}
	}
	return nil
}

func (m *Map) String() string {
	var s string
	for _, p := range m.Players {
		s += fmt.Sprintln(p)
	}
	return s
}

// Run runs a given command from the current state of the map
func (m *Map) Run(c Command) {
	switch c.Action {
	case "PLACE":
		x, y, d := ParseArgs(c.Args)
		m.AddPlayer(NewPlayer(c.Name, x, y, d))
	case "REPORT":
		m.FindPlayerByName(c.Name).Report()
	case "MOVE":
		m.FindPlayerByName(c.Name).Move()
	case "LEFT", "RIGHT":
		m.FindPlayerByName(c.Name).Rotate(c.Action)
	}
}
