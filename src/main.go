package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := readFile("fixtures/example1.txt")
	if err != nil {
		log.Fatalf("Failed reading file: %v\n", err)
	}

	err = readLine(f)
	if err != nil {
		log.Fatalf("Failed reading line: %v\n", err)
	}
}

func readFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readLine(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}
