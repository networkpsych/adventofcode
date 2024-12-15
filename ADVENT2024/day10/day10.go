package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coor struct {
	Y int
	X int
}

func main() {
	file, err := os.Open("day10.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	arr := []int{}
	tMap := [][]int{}
	for s.Scan() {
		line := strings.Split(s.Text(), "")

		for _, col := range line {
			val, _ := strconv.Atoi(col)
			arr = append(arr, val)
		}
		tMap = append(tMap, arr)
		arr = nil
	}

	PartOne(tMap)
}

func PartOne(chart [][]int) {
	p1Chart := []coor{}
	total := 0
	for y, val := range chart {
		for x, n := range val {
			if n == 0 {
				p1Chart = append(p1Chart, coor{Y: y, X: x})
			}
		}
	}
	for _, points := range p1Chart {
		total += FindRoutesP1(chart, points)
	}
	fmt.Println("P1 Total: ", total)
	total = 0
	for _, points := range p1Chart {
		total += FindRoutesP2(chart, points)
	}
	fmt.Println("P2 Total: ", total)
}

func FindRoutesP1(chart [][]int, sPoints coor) int {
	maxR := len(chart)
	maxC := len(chart[0])
	total := 0
	look := [][]int{ // NESW
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	queue := []coor{sPoints}
	seen := map[coor]bool{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if chart[c.Y][c.X] == 9 && !seen[c] {
			seen[c] = true
			total += 1
			continue
		}
		if seen[c] {
			continue
		} else {
			seen[c] = true
		}
		for _, d := range look {
			row, col := c.Y+d[0], c.X+d[1]
			if row >= 0 && col >= 0 && row < maxR && col < maxC {
				if chart[row][col] == chart[c.Y][c.X]+1 {
					queue = append(queue, coor{row, col})
				}
			}
		}
	}
	return total
}

func FindRoutesP2(chart [][]int, sPoints coor) int {
	maxR := len(chart)
	maxC := len(chart[0])
	total := 0
	look := [][]int{ // NESW
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	queue := []coor{sPoints}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if chart[c.Y][c.X] == 9 {
			total += 1
			continue
		}
		for _, d := range look {
			row, col := c.Y+d[0], c.X+d[1]
			if row >= 0 && col >= 0 && row < maxR && col < maxC {
				if chart[row][col] == chart[c.Y][c.X]+1 {
					queue = append(queue, coor{row, col})
				}
			}
		}
	}
	return total
}
