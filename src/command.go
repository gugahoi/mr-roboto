package main

import (
	"log"
	"regexp"
	"strconv"
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
	if len(s) < 2 {
		log.Printf("unable to parse command: '%v'", line)
		return Command{}
	}
	c := Command{
		Name:   s[0],
		Action: s[1],
		Args:   s[2:],
	}

	return c
}

// ParseArgs ...
func ParseArgs(a []string) (int, int, Direction) {
	x, _ := strconv.Atoi(a[0])
	y, _ := strconv.Atoi(a[1])
	return x, y, StringToDirection(a[2])
}
