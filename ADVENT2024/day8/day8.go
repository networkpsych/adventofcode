package main

import (
	"bufio"
	"fmt"
	"os"
)

type loc struct {
	x, y int
	val  string
}

type aloc struct {
	x, y int
}

type chart struct {
	cMap [][]string
}

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	pchart := chart{}

	for s.Scan() {
		line := []string{}
		for _, char := range s.Text() {
			line = append(line, string(char))
		}
		pchart.cMap = append(pchart.cMap, line)
	}

	PartOne(pchart)

}

func PartOne(chrt chart) {
	ta := []loc{}
	for row, line := range chrt.cMap {
		for col, c := range line {
			if c == "0" || c == "A" {
				item := loc{
					x:   col,
					y:   row,
					val: c,
				}
				ta = append(ta, item)
			}
		}
	}
	points := CheckPoints(len(chrt.cMap), len(chrt.cMap[0])-1, ta, &chrt)
	for _, p := range points {
		chrt.cMap[p[0]][p[1]] = "#"
	}
	for _, line := range chrt.cMap {
		l := ""
		for _, char := range line {
			l += char
		}
		fmt.Println(l)
	}
}

func CheckPoints(maxRow, maxCol int, r []loc, c *chart) [][]int {

	newP := make(map[aloc]bool)
	points := [][]int{}

	for _, p := range r {
		for o := 1; o < len(r); o++ {
			if p.val == r[0].val {
				j := CheckValidPoint(p, r[o], maxRow, maxCol, c)
				if len(j) != 0 {
					for _, vp := range j {
						newP[vp] = true
						tmp := []int{vp.y, vp.x}
						points = append(points, tmp)
					}
				}
			}
		}

	}

	fmt.Println(newP)
	fmt.Println(len(newP))

	return points
}

func CheckValidPoint(curr loc, iter loc, maxRow int, maxCol int, c *chart) []aloc {
	validPoints := []aloc{}

	testP := []aloc{
		{x: curr.x - iter.x, y: curr.y - iter.y},
		{x: curr.x + iter.x, y: curr.y - iter.y},
		{x: curr.x - iter.x, y: curr.y + iter.y},
		{x: curr.x + iter.x, y: curr.y + iter.y},
	}

	for _, p := range testP {
		if p.x > 0 && p.y > 0 && p.x <= maxCol && p.y <= maxRow {
			if c.Get(p.y, p.x) != "0" && c.Get(p.y, p.x) != "A" {
				validPoints = append(validPoints, p)
			}
		}
	}

	return validPoints

}

func (chrt chart) Get(row, col int) string {
	return chrt.cMap[row][col]
}
