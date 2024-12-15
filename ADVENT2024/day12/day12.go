package main

import (
	"advent2024/helper"
	"fmt"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("advent/ADVENT2024/day12/test.txt")
	file, err := helper.LoadFileStringChart(path)

	if err != nil {
		panic(err)
	}

	fmt.Println(file)
}
