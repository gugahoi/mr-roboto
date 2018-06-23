package main

import (
	"bufio"
	"strings"
)

func main() {

	input := `ALICE: PLACE 1,2,EAST
ALICE: MOVE
ALICE: MOVE
ALICE: LEFT
BOB: PLACE 3,3,EAST
BOB: MOVE
ALICE: MOVE
ALICE: REPORT
BOB: RIGHT
BOB: MOVE
BOB: REPORT
`
	scanner := bufio.NewScanner(strings.NewReader(input))
	m, _ := NewMap(6)
	for scanner.Scan() {
		c := Parse(scanner.Text())
		m.Run(c)
	}
}
