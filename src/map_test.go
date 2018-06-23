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
	_, x, y := m.FindPlayerByName("DAVE")
	if x != 1 {
		t.Fatalf("Expected x position to be 1, got %v", x)
	}
	if y != 0 {
		t.Fatalf("Expected y position to be 0, got %v", y)
	}

	// TODO: add interface to output to be able to test
	c = Command{Action: "REPORT", Name: "DAVE"}
	m.Run(c)

	c = Command{Action: "LEFT", Name: "DAVE"}
	m.Run(c)
	if p.Direction != North {
		t.Fatalf("Expected player 'DAVE' to be facing '%v', found '%v", North, p.Direction)
	}
	c = Command{Action: "RIGHT", Name: "DAVE"}
	m.Run(c)
	if p.Direction != East {
		t.Fatalf("Expected player 'DAVE' to be facing '%v', found '%v", North, p.Direction)
	}
}

func TestMap_Move(t *testing.T) {
	testCases := []struct {
		desc                 string
		expectedX, expectedY int
		d                    Direction
	}{
		{
			desc:      "should move north",
			d:         North,
			expectedX: 0,
			expectedY: 1,
		},
		{
			desc:      "should fail to move South out of the board",
			d:         South,
			expectedX: 0,
			expectedY: 0,
		},
		{
			desc:      "should fail to move into other player",
			d:         East,
			expectedX: 0,
			expectedY: 0,
		},
		{
			desc:      "should fail move west out of the board",
			d:         West,
			expectedX: 0,
			expectedY: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := NewMap()
			m.AddPlayer(NewPlayer("Hazard", tC.d), 0, 0)
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
