package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {

	}

	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {

	}
}
