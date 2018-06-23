package main_test

import (
	"fmt"
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestNewPlayer(t *testing.T) {
	name := "Mary"
	x := 5
	y := 5
	var direction Direction = East

	p := NewPlayer(name, x, y, direction)
	if p.Name != name {
		t.Errorf("Expected player name to be '%v', got '%v'", name, p.Name)
	}
	if p.Pos.X != x {
		t.Errorf("Expected player x coordinate to be '%v', got '%v'", x, p.Pos.X)
	}
	if p.Pos.Y != y {
		t.Errorf("Expected player y coordinate to be '%v', got '%v'", y, p.Pos.Y)
	}
	if p.Direction != direction {
		t.Errorf("Expected player direction to be '%v', got '%v'", direction, p.Direction)
	}
}

func TestPlayer_Move(t *testing.T) {
	p := NewPlayer("Shish", 0, 0, North)

	p.Move()
	if p.Pos.Y != 1 {
		t.Fatalf("Expected player to have moved north, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}

	p.Direction = East
	p.Move()
	if p.Pos.X != 1 {
		t.Fatalf("Expected player to have moved west, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}
	p.Direction = South
	p.Move()
	if p.Pos.Y != 0 {
		t.Fatalf("Expected player to have moved south, position is: %v x %v", p.Pos.X, p.Pos.Y)
	}
	p.Direction = West
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

func TestPlayer_Rotate(t *testing.T) {
	testCases := []struct {
		direction Direction
		rotate    string
		expected  Direction
	}{
		{
			direction: North,
			rotate:    "LEFT",
			expected:  West,
		},
		{
			direction: North,
			rotate:    "RIGHT",
			expected:  East,
		},
		{
			direction: South,
			rotate:    "LEFT",
			expected:  East,
		},
		{
			direction: South,
			rotate:    "RIGHT",
			expected:  West,
		},
		{
			direction: East,
			rotate:    "LEFT",
			expected:  North,
		},
		{
			direction: East,
			rotate:    "RIGHT",
			expected:  South,
		},
		{
			direction: West,
			rotate:    "LEFT",
			expected:  South,
		},
		{
			direction: West,
			rotate:    "RIGHT",
			expected:  North,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("player is facing %s rotates %s", tC.direction, tC.rotate), func(t *testing.T) {
			p := NewPlayer("Test", 0, 0, tC.direction)
			p.Rotate(tC.rotate)

			if p.Direction != tC.expected {
				t.Fatalf("Expected player to be facing %v, got %v", tC.expected, p.Direction)
			}
		})
	}
}

func ExamplePlayer_String() {
	fmt.Print(NewPlayer("Harry", 0, 0, West))
	// Output:
	// Harry: 0,0,WEST
}
