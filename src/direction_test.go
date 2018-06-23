package main_test

import (
	"fmt"
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestStringToDirection(t *testing.T) {
	directions := []Direction{Uknown, North, East, West, South}
	directionsString := []string{"UKNOWN", "NORTH", "EAST", "WEST", "SOUTH"}
	for idx := range directions {
		if directions[idx] != StringToDirection(directionsString[idx]) {
			t.Errorf("Expected direction string to be '%s', got '%s'", directionsString[idx], directions[idx])
		}
	}
}

func TestDirection_String(t *testing.T) {
	directions := []Direction{Uknown, North, East, West, South}
	directionsString := []string{"UKNOWN", "NORTH", "EAST", "WEST", "SOUTH"}
	for idx := range directions {
		if fmt.Sprint(directions[idx]) != directionsString[idx] {
			t.Errorf("Expected direction string to be '%s', got '%s'", directionsString[idx], directions[idx])
		}
	}
}
