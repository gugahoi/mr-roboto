package main

// Direction indicates where a player is facing
type Direction int

const (
	// Uknown is the default direction
	Uknown Direction = iota
	// North direction
	North
	// East direction
	East
	// South direction
	South
	// West direction
	West
)

// StringToDirection converts a string into a direction enum
func StringToDirection(d string) Direction {
	switch d {
	case "NORTH":
		return North
	case "EAST":
		return East
	case "SOUTH":
		return South
	case "WEST":
		return West
	default:
		return Uknown
	}
}

func (d Direction) String() string {
	switch d {
	case 1:
		return "NORTH"
	case 2:
		return "EAST"
	case 3:
		return "SOUTH"
	case 4:
		return "WEST"
	default:
		return "UKNOWN"
	}
}
