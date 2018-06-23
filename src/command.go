package main

import (
	"regexp"
)

// Command represents 1 line of input
type Command struct {
	Name   string
	Action string
	Args   []string
}

// Parse takes in 1 line and converts it to a command
func Parse(line string) Command {
	s := regexp.MustCompile(": | |,").Split(line, -1)
	c := Command{
		Name:   s[0],
		Action: s[1],
		Args:   s[2:],
	}

	return c
}
