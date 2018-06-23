package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	if len(os.Args) != 2 {
		usage()
	}
}

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	m := NewMap()
	for scanner.Scan() {
		c := Parse(scanner.Text())
		m.Run(c)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage:
	%s path/to/file.txt
`, os.Args[0])
	os.Exit(0)
}
