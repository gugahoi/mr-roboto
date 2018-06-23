package main_test

import (
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestParseLine(t *testing.T) {
	testCases := []struct {
		desc, name, action, args string
	}{
		{
			desc:   "ALICE: PLACE 1,2,EAST",
			name:   "ALICE",
			action: "PLACE",
			args:   "1,2,EAST",
		},
		{
			desc:   "ALICE: MOVE",
			name:   "ALICE",
			action: "MOVE",
			args:   "",
		},
		{
			desc:   "ALICE: LEFT",
			name:   "ALICE",
			action: "LEFT",
			args:   "",
		},
		{
			desc:   "ALICE: REPORT",
			name:   "ALICE",
			action: "REPORT",
			args:   "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := ParseLine(tC.desc)
			if c.Name != tC.name {
				t.Errorf("Expected name to be %v, got %v", tC.name, c.Name)
			}
			if c.Action != tC.action {
				t.Errorf("Expected action to be %v, got %v", tC.action, c.Action)
			}
			if c.Args != tC.args {
				t.Errorf("Expected args to be %v, got %v", tC.args, c.Args)
			}
		})
	}
}
