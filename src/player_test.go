package main_test

import (
	"fmt"
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestNewPlayer(t *testing.T) {
	name := "Mary"
	var direction Direction = East

	p := NewPlayer(name, direction)
	if p.Name != name {
		t.Errorf("Expected player name to be '%v', got '%v'", name, p.Name)
	}
	if p.Direction != direction {
		t.Errorf("Expected player direction to be '%v', got '%v'", direction, p.Direction)
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
			p := NewPlayer("Test", tC.direction)
			p.Rotate(tC.rotate)

			if p.Direction != tC.expected {
				t.Fatalf("Expected player to be facing %v, got %v", tC.expected, p.Direction)
			}
		})
	}
}

func ExamplePlayer_String() {
	fmt.Print(NewPlayer("Harry", West))
	// Output:
	// Harry: WEST
}
