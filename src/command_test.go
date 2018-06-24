package main_test

import (
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		desc, name, action string
		args               []string
	}{
		{
			desc:   "ALICE: PLACE 1,2,EAST",
			name:   "ALICE",
			action: "PLACE",
			args:   []string{"1", "2", "EAST"},
		},
		{
			desc:   "ALICE: MOVE",
			name:   "ALICE",
			action: "MOVE",
			args:   []string{},
		},
		{
			desc:   "ALICE: LEFT",
			name:   "ALICE",
			action: "LEFT",
			args:   []string{},
		},
		{
			desc:   "ALICE: REPORT",
			name:   "ALICE",
			action: "REPORT",
			args:   []string{},
		},
		{
			desc:   "INVALID~COMMAND:::::123123",
			name:   "",
			action: "",
			args:   []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := Parse(tC.desc)
			if c.Name != tC.name {
				t.Errorf("Expected name to be %v, got %v", tC.name, c.Name)
			}
			if c.Action != tC.action {
				t.Errorf("Expected action to be %v, got %v", tC.action, c.Action)
			}
			for idx := range c.Args {
				if c.Args[idx] != tC.args[idx] {
					t.Errorf("Expected arg at %v to be '%v', got '%v'", idx, tC.args, c.Args)
				}
			}
		})
	}
}
