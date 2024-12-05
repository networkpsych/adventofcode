package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var data []string

	file, err := os.Open("./day4.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	d := bufio.NewScanner(file)

	for d.Scan() {
		data = append(data, d.Text())

	}
	FindRecurrences(data)
}

func FindRecurrences(data []string) {

	p1 := 0
	p2 := 0
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[0]); col++ {
			if data[row][col] == 88 {
				p1 += PartOne(data, row, col)
			}
		}
	}

	for row := 1; row < len(data)-1; row++ {
		for col := 0; col < len(data[row]); col++ {
			if data[row][col] == 65 {
				p2 += PartTwo(data, row, col)
			}
		}
	}

	fmt.Printf("Part One: %d | Part Two: %d", p1, p2)

}

func PartOne(d []string, row, col int) int {
	// One x can have multiple xmas attachments
	c := 0
	if col >= 3 {
		if d[row][col] == 88 && d[row][col-1] == 77 && d[row][col-2] == 65 && d[row][col-3] == 83 {
			c++
		}
	}
	if col <= (len(d[0]) - 1 - 3) {
		if d[row][col] == 88 && d[row][col+1] == 77 && d[row][col+2] == 65 && d[row][col+3] == 83 {
			c++
		}
	}
	if row >= 3 {
		if d[row][col] == 88 && d[row-1][col] == 77 && d[row-2][col] == 65 && d[row-3][col] == 83 {
			c++
		}
	}
	if row <= (len(d[0]) - 1 - 3) {
		if d[row][col] == 88 && d[row+1][col] == 77 && d[row+2][col] == 65 && d[row+3][col] == 83 {
			c++
		}
	}
	if row >= 3 && col <= (len(d[0])-1-3) {
		if d[row][col] == 88 && d[row-1][col+1] == 77 && d[row-2][col+2] == 65 && d[row-3][col+3] == 83 {
			c++
		}
	}
	if row >= 3 && col >= 3 {

		if d[row][col] == 88 && d[row-1][col-1] == 77 && d[row-2][col-2] == 65 && d[row-3][col-3] == 83 {
			c++
		}
	}
	if row <= (len(d[0])-1-3) && col <= (len(d[0])-1-3) {
		if d[row][col] == 88 && d[row+1][col+1] == 77 && d[row+2][col+2] == 65 && d[row+3][col+3] == 83 {
			c++
		}
	}
	if row <= (len(d[0])-1-3) && col >= 3 {
		if d[row][col] == 88 && d[row+1][col-1] == 77 && d[row+2][col-2] == 65 && d[row+3][col-3] == 83 {
			c++
		}
	}
	return c
}

func PartTwo(d []string, row, col int) int {
	c := 0
	/*
		STRUCTURE:

		M.M S.M M.S S.S
		.A. .A. .A. .A.
		S.S S.M M.S M.M

	*/
	if col < 1 || col > len(d[row])-2 {
		return 0
	}

	if d[row-1][col+1] == 77 && d[row-1][col-1] == 77 && d[row+1][col-1] == 83 && d[row+1][col+1] == 83 {
		c++
	}

	if d[row-1][col+1] == 83 && d[row-1][col-1] == 83 && d[row+1][col-1] == 77 && d[row+1][col+1] == 77 {
		c++
	}

	if d[row-1][col+1] == 83 && d[row-1][col-1] == 77 && d[row+1][col-1] == 77 && d[row+1][col+1] == 83 {
		c++
	}

	if d[row-1][col+1] == 77 && d[row-1][col-1] == 83 && d[row+1][col-1] == 83 && d[row+1][col+1] == 77 {
		c++
	}

	return c
}
