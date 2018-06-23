package main_test

import (
	"testing"

	. "github.com/gugahoi/mr-roboto/src"
)

func TestNewMap(t *testing.T) {
	testCases := []struct {
		desc string
		size int
		err  bool
	}{
		{
			desc: "create a new map of size 0",
			size: 0,
			err:  true,
		},
		{
			desc: "create a new map of size 10",
			size: 10,
			err:  false,
		},
		{
			desc: "create a new map of size -1",
			size: -1,
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m, err := NewMap(tC.size)

			if tC.err {
				if err == nil {
					t.Fatalf("Expected error to be thrown but got none")
				}
			} else {
				if m.Size != tC.size {
					t.Fatalf("Map created with wrong dimensions: expected '%v', got %v", tC.size, m.Size)
				}
			}
		})
	}
}

func TestAddPlayer(t *testing.T) {
	m := Map{
		Size:    10,
		Players: []*Player{},
	}

	p := NewPlayer(
		"Mary",
		0,
		0,
		"EAST",
	)

	m.AddPlayer(p)
	if size := len(m.Players); size != 1 {
		t.Fatalf("Expected to have %v players, got %v", 1, size)
	}
}

func TestAddSamePlayer(t *testing.T) {
	m := Map{
		Size:    10,
		Players: []*Player{},
	}

	p := NewPlayer(
		"Mary",
		0,
		0,
		"EAST",
	)

	m.AddPlayer(p)
	m.AddPlayer(p)
	if size := len(m.Players); size != 1 {
		t.Fatalf("Expected to have %v players, got %v", 1, size)
	}
}
