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

	err = m.AddPlayer(p, 2, 3)
	if err == nil {
		t.Fatalf("Expected not to be able to add player, did not get err")
	}

	err = m.AddPlayer(NewPlayer("Kaleb", West), 0, 0)
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
	cases := []struct {
		command   Command
		direction Direction
		x, y      int
	}{
		{
			command:   Command{Action: "PLACE", Args: []string{"0", "0", "EAST"}, Name: "DAVE"},
			direction: East,
			x:         0,
			y:         0,
		},
		{
			command:   Command{Action: "MOVE", Name: "DAVE"},
			direction: East,
			x:         1,
			y:         0,
		},
		{
			command:   Command{Action: "REPORT", Name: "DAVE"},
			direction: East,
			x:         1,
			y:         0,
		},
		{
			command:   Command{Action: "LEFT", Name: "DAVE"},
			direction: North,
			x:         1,
			y:         0,
		},
		{
			command:   Command{Action: "RIGHT", Name: "DAVE"},
			direction: East,
			x:         1,
			y:         0,
		},
		{
			command:   Command{Action: "SOMUnsupoorted\\Action"},
			direction: East,
			x:         1,
			y:         0,
		},
	}

	m := NewMap()
	// Note these cases are run in sequence
	for _, c := range cases {
		m.Run(c.command)
		p, x, y := m.FindPlayerByName("DAVE")
		if p.Direction != c.direction {
			t.Fatalf("Expected player 'DAVE' to be facing '%v', found '%v", c.direction, p.Direction)
		}
		if x != c.x {
			t.Fatalf("Expected player 'DAVE' to have x position '%v', found '%v", c.x, x)
		}
		if y != c.y {
			t.Fatalf("Expected player 'DAVE' to have y position '%v', found '%v", c.y, y)
		}
	}
}

func TestMap_Move(t *testing.T) {
	testCases := []struct {
		desc                 string
		startX, startY       int
		expectedX, expectedY int
		d                    Direction
	}{
		{
			desc:      "should move north",
			startX:    0,
			startY:    0,
			d:         North,
			expectedX: 0,
			expectedY: 1,
		},
		{
			desc:      "should fail to move South out of the board",
			startX:    0,
			startY:    0,
			d:         South,
			expectedX: 0,
			expectedY: 0,
		},
		{
			desc:      "should fail to move into other player",
			startX:    0,
			startY:    0,
			d:         East,
			expectedX: 0,
			expectedY: 0,
		},
		{
			desc:      "should fail to move west out of the board",
			startX:    0,
			startY:    0,
			d:         West,
			expectedX: 0,
			expectedY: 0,
		},
		{
			desc:      "should move west",
			startX:    4,
			startY:    4,
			d:         West,
			expectedX: 3,
			expectedY: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := NewMap()
			m.AddPlayer(NewPlayer("Hazard", tC.d), tC.startX, tC.startY)
			m.AddPlayer(NewPlayer("Lukaku", North), 1, 0)

			m.Move("Hazard")
			_, x, y := m.FindPlayerByName("Hazard")
			if x != tC.expectedX {
				t.Fatalf("Expected player to have moved in x coordinate to %v, got %v", tC.expectedX, x)
			}
			if y != tC.expectedY {
				t.Fatalf("Expected player to have moved in y coordinate to %v, got %v", tC.expectedY, y)
			}
			m.Move("Lee")
		})
	}
}
