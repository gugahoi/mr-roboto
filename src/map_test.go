package main_test

import (
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
		Players: make(map[[2]int]*Player),
	}

	p := NewPlayer(
		"Mary",
		East,
	)

	err := m.AddPlayer(p, [2]int{0, 0})
	if err != nil {
		t.Fatalf("Expected player to be added, got err '%v'", err)
	}

	err = m.AddPlayer(p, [2]int{2, 3})
	if err == nil {
		t.Fatalf("Expected not to be able to add player, did not get err")
	}

	err = m.AddPlayer(NewPlayer("Kaleb", West), [2]int{0, 0})
	if err == nil {
		t.Fatalf("Expected not to be able to add player, did not get err")
	}

	err = m.AddPlayer(NewPlayer("Tristan", East), [2]int{-1, 0})
	if err == nil {
		t.Fatalf("Expected not to be able to add player as x is out of bounds, did not get err")
	}

	err = m.AddPlayer(NewPlayer("Tristan", East), [2]int{1, -1})
	if err == nil {
		t.Fatalf("Expected not to be able to add player as y is out of bounds, did not get err")
	}
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
		p, pos := m.FindPlayerByName("DAVE")
		if p.Direction != c.direction {
			t.Fatalf("Expected player 'DAVE' to be facing '%v', found '%v", c.direction, p.Direction)
		}
		if pos[0] != c.x {
			t.Fatalf("Expected player 'DAVE' to have x position '%v', found '%v", c.x, pos[0])
		}
		if pos[1] != c.y {
			t.Fatalf("Expected player 'DAVE' to have y position '%v', found '%v", c.y, pos[1])
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
			m.AddPlayer(NewPlayer("Hazard", tC.d), [2]int{tC.startX, tC.startY})
			m.AddPlayer(NewPlayer("Lukaku", North), [2]int{1, 0})

			m.Move("Hazard")
			_, pos := m.FindPlayerByName("Hazard")
			if pos[0] != tC.expectedX {
				t.Fatalf("Expected player to have moved in x coordinate to %v, got %v", tC.expectedX, pos[0])
			}
			if pos[1] != tC.expectedY {
				t.Fatalf("Expected player to have moved in y coordinate to %v, got %v", tC.expectedY, pos[1])
			}
			m.Move("Lee")
		})
	}
}
