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
		Players: [6][6]*Player{},
	}

	p := NewPlayer(
		"Mary",
		East,
	)

	err := m.AddPlayer(p, 0, 0)
	if err != nil {
		t.Fatalf("Expected player to be added, got err '%v'", err)
	}

	err = m.AddPlayer(p, 0, 0)
	if err == nil {
		t.Fatalf("Expected not to be able to add player, did not get err")
	}
}

func ExampleMap_String() {
	m := NewMap()
	m.AddPlayer(NewPlayer("Joe", East), 0, 0)
	m.AddPlayer(NewPlayer("Mary", West), 1, 0)
	m.AddPlayer(NewPlayer("Moses", North), 5, 2)

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

	p, _, _ := m.FindPlayerByName("DAVE")
	if p == nil {
		t.Fatal("Expected to find player 'DAVE' but didn't")
	}
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
