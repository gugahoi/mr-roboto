package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// run the game
	scanner := bufio.NewScanner(Input(os.Args))
	m := NewMap()
	for scanner.Scan() {
		c := Parse(scanner.Text())
		m.Run(c)
	}
}

// Input reads the first argument passed in as a file.
// If no arguments it reads continually from STDIN.
func Input(args []string) io.Reader {
	if len(args) < 2 {
		return os.Stdin
	}
	content, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}
	return strings.NewReader(string(content))
}
