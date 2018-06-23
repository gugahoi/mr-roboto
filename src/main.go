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
	// set up input
	var i string
	if len(os.Args) == 2 {
		i = os.Args[1]
	}
	scanner := bufio.NewScanner(input(i))

	// run the game
	m := NewMap()
	for scanner.Scan() {
		c := Parse(scanner.Text())
		m.Run(c)
	}
}

func input(file string) io.Reader {
	if file != "" {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		return strings.NewReader(string(content))
	}
	return os.Stdin
}
