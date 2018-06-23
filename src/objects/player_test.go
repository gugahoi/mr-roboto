package objects_test

import (
	"fmt"
	"testing"

	"github.com/gugahoi/mr-roboto/src/objects"
)

func TestNewPlayer(t *testing.T) {
	name := "Mary"
	x := 5
	y := 5
	direction := "EAST"

	p := objects.NewPlayer(name, x, y, direction)
	if p.Name != name {
		t.Fatalf("Expected player name to be '%v', got '%v'", name, p.Name)
	}
	if p.Direction != direction {
		t.Fatalf("Expected player direction to be '%v', got '%v'", direction, p.Direction)
	}
}

func TestMove(t *testing.T) {
	p := objects.NewPlayer("Shish", 0, 0, "NORTH")

	p.Move()
	if p.Pos.Y != 1 {
		t.Fatalf("Expected player to have moved north, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}

	p.Direction = "EAST"
	p.Move()
	if p.Pos.X != 1 {
		t.Fatalf("Expected player to have moved west, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}
	p.Direction = "SOUTH"
	p.Move()
	if p.Pos.Y != 0 {
		t.Fatalf("Expected player to have moved south, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}
	p.Direction = "WEST"
	p.Move()
	if p.Pos.X != 0 {
		t.Fatalf("Expected player to have moved west, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}

	p.Direction = "NOT VALID"
	err := p.Move()
	if err == nil {
		t.Fatal("Expected invalid direction error, got nil")
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		direction string
		rotate    string
		expected  string
	}{
		{
			direction: "NORTH",
			rotate:    "LEFT",
			expected:  "WEST",
		},
		{
			direction: "NORTH",
			rotate:    "RIGHT",
			expected:  "EAST",
		},
		{
			direction: "SOUTH",
			rotate:    "LEFT",
			expected:  "EAST",
		},
		{
			direction: "SOUTH",
			rotate:    "RIGHT",
			expected:  "WEST",
		},
		{
			direction: "EAST",
			rotate:    "LEFT",
			expected:  "NORTH",
		},
		{
			direction: "EAST",
			rotate:    "RIGHT",
			expected:  "SOUTH",
		},
		{
			direction: "WEST",
			rotate:    "LEFT",
			expected:  "SOUTH",
		},
		{
			direction: "WEST",
			rotate:    "RIGHT",
			expected:  "NORTH",
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("player is facing %s rotates %s", tC.direction, tC.rotate), func(t *testing.T) {
			p := objects.NewPlayer("Test", 0, 0, tC.direction)
			p.Rotate(tC.rotate)

			if p.Direction != tC.expected {
				t.Fatalf("Expected player to be facing %v, got %v", tC.expected, p.Direction)
			}
		})
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
