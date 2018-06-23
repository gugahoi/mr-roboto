package main_test

import "github.com/gugahoi/mr-roboto/src/objects"

func ExampleScenario1() {
	// scenario 1
	m, _ := objects.NewMap(10)
	// ALICE: PLACE 0,0,NORTH
	m.AddPlayer(objects.NewPlayer("ALICE", 0, 0, "NORTH"))
	// ALICE: MOVE
	m.FindPlayerByName("ALICE").Move()
	// ALICE: REPORT
	m.FindPlayerByName("ALICE").Report()
	// Output:
	// ALICE: 0,1,NORTH
}

func ExampleScenario2() {
	// scenario 2
	m, _ := objects.NewMap(10)
	// BOB: PLACE 0,0,NORTH
	m.AddPlayer(objects.NewPlayer("BOB", 0, 0, "NORTH"))
	// BOB: LEFT
	m.FindPlayerByName("BOB").Rotate("LEFT")
	// BOB: REPORT
	m.FindPlayerByName("BOB").Report()
	// Output:
	// BOB: 0,0,WEST

}

func ExampleScenario3() {
	// scenario 3
	m, _ := objects.NewMap(10)
	// ALICE: PLACE 1,2,EAST
	m.AddPlayer(objects.NewPlayer("ALICE", 1, 2, "EAST"))
	// ALICE: MOVE
	m.FindPlayerByName("ALICE").Move()
	// ALICE: MOVE
	m.FindPlayerByName("ALICE").Move()
	// ALICE: LEFT
	m.FindPlayerByName("ALICE").Rotate("LEFT")
	// BOB: PLACE 3,3,EAST
	m.AddPlayer(objects.NewPlayer("BOB", 3, 3, "EAST"))
	// BOB: MOVE
	m.FindPlayerByName("BOB").Move()
	// ALICE: MOVE
	m.FindPlayerByName("ALICE").Move()
	// ALICE: REPORT
	m.FindPlayerByName("ALICE").Report()
	// BOB: RIGHT
	m.FindPlayerByName("BOB").Rotate("RIGHT")
	// BOB: MOVE
	m.FindPlayerByName("BOB").Move()
	// BOB: REPORT
	m.FindPlayerByName("BOB").Report()
	// Output:
	// ALICE: 3,3,NORTH
	// BOB: 4,2,SOUTH
}

func ExampleScenario4() {
	// scenario 4
	m, _ := objects.NewMap(10)
	// BOB: PLACE 1,3,SOUTH
	m.AddPlayer(objects.NewPlayer("BOB", 1, 3, "SOUTH"))
	// ALICE: PLACE 0,1,EAST
	m.AddPlayer(objects.NewPlayer("ALICE", 0, 1, "EAST"))
	// ALICE: MOVE
	m.FindPlayerByName("ALICE").Move()
	// BOB: MOVE
	m.FindPlayerByName("BOB").Move()
	// BOB: MOVE
	m.FindPlayerByName("BOB").Move()
	// ALICE: MOVE
	m.FindPlayerByName("ALICE").Move()
	// BOB: MOVE
	m.FindPlayerByName("BOB").Move()
	// BOB: LEFT
	m.FindPlayerByName("BOB").Rotate("LEFT")
	// ALICE: REPORT
	m.FindPlayerByName("ALICE").Report()
	// BOB: REPORT
	m.FindPlayerByName("BOB").Report()
	// Output:
	// ALICE: 2,1,EAST
	// BOB: 1,1,EAST
}