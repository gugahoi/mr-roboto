package main_test

import (
	"io"
	"os"
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestInput(t *testing.T) {
	testCases := []struct {
		desc     string
		args     []string
		expected io.Reader
	}{
		{
			desc:     "No input should read from stdin",
			args:     []string{"program"},
			expected: os.Stdin,
		},
		// TODO: mock the file to test
		// {
		// 	desc:     "Input should reaturn the first file",
		// 	args:     []string{"program", "inputFile"},
		// 	expected: os.Stdin,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Input(tC.args)
			if got != tC.expected {
				t.Fatalf("Expected %v, got %v", tC.expected, got)
			}
		})
	}
}

func Example_Scenario1() {
	// scenario 1
	m := NewMap()
	// ALICE: PLACE 0,0,NORTH
	m.AddPlayer(NewPlayer("ALICE", North), [2]int{0, 0})
	// ALICE: MOVE
	m.Move("ALICE")
	// ALICE: REPORT
	m.Report("ALICE")
	// Output:
	// ALICE: 0,1,NORTH
}

func Example_Scenario2() {
	// scenario 2
	m := NewMap()
	// BOB: PLACE 0,0,NORTH
	m.AddPlayer(NewPlayer("BOB", North), [2]int{0, 0})
	// BOB: LEFT
	p, _ := m.FindPlayerByName("BOB")
	p.Rotate("LEFT")
	// BOB: REPORT
	m.Report("BOB")
	// Output:
	// BOB: 0,0,WEST

}

func Example_Scenario3() {
	// scenario 3
	m := NewMap()
	// ALICE: PLACE 1,2,EAST
	m.AddPlayer(NewPlayer("ALICE", East), [2]int{1, 2})
	// ALICE: MOVE
	m.Move("ALICE")
	// ALICE: MOVE
	m.Move("ALICE")
	// ALICE: LEFT
	p, _ := m.FindPlayerByName("ALICE")
	p.Rotate("LEFT")
	// BOB: PLACE 3,3,EAST
	m.AddPlayer(NewPlayer("BOB", East), [2]int{3, 3})
	// BOB: MOVE
	m.Move("BOB")
	// ALICE: MOVE
	m.Move("ALICE")
	// ALICE: REPORT
	m.Report("ALICE")
	// BOB: RIGHT
	p, _ = m.FindPlayerByName("BOB")
	p.Rotate("RIGHT")
	// BOB: MOVE
	m.Move("BOB")
	// BOB: REPORT
	m.Report("BOB")
	// Output:
	// ALICE: 3,3,NORTH
	// BOB: 4,2,SOUTH
}

func Example_Scenario4() {
	// scenario 4
	m := NewMap()
	// BOB: PLACE 1,3,SOUTH
	m.AddPlayer(NewPlayer("BOB", South), [2]int{1, 3})
	// ALICE: PLACE 0,1,EAST
	m.AddPlayer(NewPlayer("ALICE", East), [2]int{0, 1})
	// ALICE: MOVE
	m.Move("ALICE")
	// BOB: MOVE
	m.Move("BOB")
	// BOB: MOVE
	m.Move("BOB")
	// ALICE: MOVE
	m.Move("ALICE")
	// BOB: MOVE
	m.Move("BOB")
	// BOB: LEFT
	p, _ := m.FindPlayerByName("BOB")
	p.Rotate("LEFT")
	// ALICE: REPORT
	m.Report("ALICE")
	// BOB: REPORT
	m.Report("BOB")
	// Output:
	// ALICE: 2,1,EAST
	// BOB: 1,1,EAST
}
