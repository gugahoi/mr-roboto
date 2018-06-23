package main_test

import (
	"fmt"
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestNewMap(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "create a new map",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := NewMap()
			_ = m
		})
	}
}

func TestMap_AddPlayer(t *testing.T) {
	m := Map{
		Players: []*Player{},
	}

	p := NewPlayer(
		"Mary",
		0,
		0,
		East,
	)

	m.AddPlayer(p)
	if size := len(m.Players); size != 1 {
		t.Fatalf("Expected to have %v players, got %v", 1, size)
	}
}

func TestMap_AddSamePlayer(t *testing.T) {
	m := Map{
		Players: []*Player{},
	}

	p := NewPlayer(
		"Mary",
		0,
		0,
		East,
	)

	m.AddPlayer(p)
	m.AddPlayer(p)
	if size := len(m.Players); size != 1 {
		t.Fatalf("Expected to have %v players, got %v", 1, size)
	}
}

func ExampleMap_String() {
	m := NewMap()
	m.AddPlayer(NewPlayer("Joe", 0, 0, East))
	m.AddPlayer(NewPlayer("Mary", 1, 0, West))
	m.AddPlayer(NewPlayer("Moses", 5, 2, North))

	fmt.Print(m)
	// Output:
	// Joe: 0,0,EAST
	// Mary: 1,0,WEST
	// Moses: 5,2,NORTH
}

func TestMap_Run(t *testing.T) {
	m := NewMap()
	c := Command{Action: "PLACE", Args: []string{"0", "0", "EAST"}, Name: "DAVE"}
	m.Run(c)

	p := m.FindPlayerByName("DAVE")
	if p == nil {
		t.Fatal("Expected to find player 'DAVE' but didn't")
	}
	if p.Pos.X != 0 {
		t.Fatalf("Expected player 'DAVE' to be at x coordinate '0', found '%v", p.Pos.X)
	}
	if p.Pos.Y != 0 {
		t.Fatalf("Expected player 'DAVE' to be at Y coordinate '0', found '%v", p.Pos.Y)
	}
	// TODO: fix this direction
	if p.Direction != East {
		t.Fatalf("Expected player 'DAVE' to be facing '%v', found '%v", East, p.Direction)
	}

	c = Command{Action: "MOVE", Name: "DAVE"}
	m.Run(c)
	c = Command{Action: "REPORT", Name: "DAVE"}
	m.Run(c)
	c = Command{Action: "LEFT", Name: "DAVE"}
	m.Run(c)
	c = Command{Action: "RIGHT", Name: "DAVE"}
	m.Run(c)
}
