package main

import (
	"advent/helper"
	"fmt"
	"path/filepath"
)

type coor struct {
	y, x int
}

type shape struct {
	letter    string
	area      []coor
	perimeter int
}

func main() {
	path, _ := filepath.Abs("test.txt")
	file, err := helper.LoadFileStringChart(path, "")

	if err != nil {
		panic(err)
	}

	fmt.Println(file)
	partOne(file)
}

func partOne(chart [][]string) {

	total := 0
	p1 := map[string][]coor{}
	for y, row := range chart {
		for x, col := range row {
			if arrCheck(p1[col], coor{y, x}) {
				continue
			} else {
				p := perimeter(chart, y, x, col)
				p1[col] = append(p1[col], p.area...)
				total += p.perimeter
			}

		}
	}
	//fmt.Println(p1)
	fmt.Println("PART ONE: ", total)
	total = 0
	p2 := map[string][]coor{}
	perim := map[int]int{}

	for y, row := range chart {
		for x, col := range row {
			if arrCheck(p2[col], coor{y, x}) {
				continue
			} else {
				p := perimeter(chart, y, x, col)
				p2[col] = append(p2[col], p.area...)
				//total += p.perimeter
			}
		}
	}
	pLen := 0
	for key, val := range p2 {
		for _, v := range val {
			if _, ok := perim[v.y]; !ok {
				perim[v.y] = v.x
				pLen += 1
			}
		}
		total += len(p2[key]) * pLen
	}
	fmt.Println(perim)
	fmt.Println("PART TWO: ", total)

}

func arrCheck(arr []coor, val coor) bool {
	for _, c := range arr {
		if c == val {
			return true
		}
	}
	return false
}

func perimeter(chart [][]string, r, c int, letter string) shape {
	maxR := len(chart)
	maxC := len(chart[0])
	look := [][]int{ // NESW
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	seen := map[coor]bool{}
	p := 0
	queue := []coor{{r, c}}
	for len(queue) > 0 {
		c := queue[0] // current queue item
		queue = queue[1:]
		if _, ok := seen[c]; ok {
			continue
		}
		if c.y >= maxR || c.x >= maxC || c.y < 0 || c.x < 0 || chart[c.y][c.x] != letter {
			p += 1
			continue
		}
		seen[c] = true
		for _, xy := range look {
			row := c.y + xy[0]
			col := c.x + xy[1]
			if _, ok := seen[coor{row, col}]; !ok {
				queue = append(queue, coor{row, col})
			}
		}
		//fmt.Println(seen)
	}
	area := []coor{}
	for key, _ := range seen {
		area = append(area, key)
	}
	output := shape{
		letter:    letter,
		area:      area,
		perimeter: p * len(area),
	}
	return output

}
