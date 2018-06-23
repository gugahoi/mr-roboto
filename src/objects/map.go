package objects

import (
	"errors"
	"fmt"
	"log"
)

// Map represents the map game is played in and contains the players in the game
type Map struct {
	Size    int
	Players []*Player
}

// Map Layout (x,y)
// x coordinate is west -> east incrementally
// y coordinate is south -> north incrementally
//                NORTH
//   | 0,5 | 1,5 | 2,5 | 3,5 | 4,5 | 5,5 |
// W | 0,4 | 1,4 | 2,4 | 3,4 | 4,4 | 5,4 | E
// E | 0,3 | 1,3 | 2,3 | 3,3 | 4,3 | 5,3 | A
// S | 0,2 | 1,2 | 2,2 | 3,2 | 4,2 | 5,2 | S
// T | 0,1 | 1,1 | 2,1 | 3,1 | 4,1 | 5,1 | T
//   | 0,0 | 1,0 | 2,0 | 3,0 | 4,0 | 5,0 |
//                SOUTH

// NewMap creates a new map of given size (square map)
func NewMap(size int) (*Map, error) {
	if size < 1 {
		return nil, errors.New("InvalidMapSize: map size must be bigger than 0")
	}

	return &Map{
		Size:    size,
		Players: []*Player{},
	}, nil
}

// AddPlayer adds a new player to the map
func (m *Map) AddPlayer(p Player) {
	if m.hasPlayer(p) {
		log.Printf("Player %v already in map, skipping\n", p)
		return
	}

	m.Players = append(m.Players, &p)
}

// hasPlayer returns true/false if the player has laready been placed in the map
func (m *Map) hasPlayer(p Player) bool {
	for _, player := range m.Players {
		fmt.Println(p.Name, player.Name)
		if player.Name == p.Name {
			return true
		}
	}
	return false
}