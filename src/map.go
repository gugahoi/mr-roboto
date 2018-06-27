package main

import (
	"fmt"
	"log"
)

const size = 6

// Map contains the players in the game
type Map struct {
	Players map[[2]int]*Player
}

// Map Layout (X,Y)
// X coordinate is west -> east incrementally
// Y coordinate is south -> north incrementally
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
		Players: make(map[[2]int]*Player),
	}
}

// AddPlayer adds a new player to the map, unless a player of same name already exists
// or the chosen location is already taken
func (m *Map) AddPlayer(p *Player, pos [2]int) error {
	if pos[0] < 0 || pos[0] >= size {
		return fmt.Errorf("Invalid X coordinate: %v", pos[0])
	}

	if pos[1] < 0 || pos[1] >= size {
		return fmt.Errorf("Invalid Y coordinate: %v", pos[1])
	}

	if p, _ := m.FindPlayerByName(p.Name); p != nil {
		return fmt.Errorf("player %v already in map, skipping", p)
	}

	if _, ok := m.Players[pos]; ok {
		return fmt.Errorf("location not empty: (%v) %v", pos, p.Name)
	}

	m.Players[pos] = p
	return nil
}

// FindPlayerByName returns the player if the player with the given name
// has already been placed in the map
func (m *Map) FindPlayerByName(name string) (*Player, *[2]int) {
	for pos, player := range m.Players {
		if player != nil && player.Name == name {
			return player, &pos
		}
	}
	return nil, nil
}

// Run runs a given command from the current state of the map
func (m *Map) Run(c Command) {
	switch c.Action {
	case "PLACE":
		X, Y, d := ParseArgs(c.Args)
		m.AddPlayer(NewPlayer(c.Name, d), [2]int{X, Y})
	case "REPORT":
		m.Report(c.Name)
	case "MOVE":
		m.Move(c.Name)
	case "LEFT", "RIGHT":
		p, _ := m.FindPlayerByName(c.Name)
		if p != nil {
			p.Rotate(c.Action)
		}
	default:
		log.Printf("Unsuported command: %v\n", c.Action)
	}
}

// Report prints the player report including name, position and direction
func (m *Map) Report(name string) {
	p, pos := m.FindPlayerByName(name)
	if p != nil {
		fmt.Printf("%s: %d,%d,%s\n", p.Name, pos[0], pos[1], p.Direction)
	}
}

// Move moves the player by name if available
func (m *Map) Move(name string) {
	p, pos := m.FindPlayerByName(name)
	if p == nil {
		log.Println("Player not found, skipping move")
		return
	}

	nextPos := nextPosition(p.Direction, *pos)
	if _, ok := m.Players[nextPos]; ok {
		log.Printf("Position already taken by %v, skipping move\n", m.Players[nextPos].Name)
		return
	}

	delete(m.Players, *pos)
	m.Players[nextPos] = p
}

func nextPosition(d Direction, pos [2]int) [2]int {
	switch d {
	case North:
		if pos[1] < size-1 {
			pos[1]++
		}
	case South:
		if pos[1] > 0 {
			pos[1]--
		}
	case East:
		if pos[0] < size-1 {
			pos[0]++
		}
	case West:
		if pos[0] > 0 {
			pos[0]--
		}
	}
	return pos
}
