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
	x int
	y int
}

type chart struct {
	cMap [][]string
}

func main() {

	// Big O of don't run too much on this ;-;

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

	ta := map[string][]aloc{}
	ta1 := []loc{}
	for row, line := range pchart.cMap {
		for col, c := range line {
			if c != "." {
				item := loc{
					x:   col,
					y:   row,
					val: c,
				}
				ta1 = append(ta1, item)
				ta[c] = append(ta[c], aloc{col, row})
			}
		}
	}

	//PartOne(pchart, ta)
	PartTwo(pchart, ta)

}

func PartOne(chrt chart, ta []loc) {

	points := CheckPoints(len(chrt.cMap), len(chrt.cMap[0]), ta, &chrt)
	for _, p := range points {
		chrt.cMap[p.y][p.x] = "#"
	}
	for _, line := range chrt.cMap {
		l := ""
		for _, char := range line {
			l += char
		}
		fmt.Println(l)
	}
}

func CheckPoints(maxRow, maxCol int, r []loc, c *chart) map[aloc]aloc {

	newP := make(map[aloc]aloc)
	total := 0

	for _, p := range r {
		for o := 1; o < len(r); o++ {
			if p.val == r[o].val {
				j := CheckValidPoint(p, r[o], maxRow, maxCol, c)
				if len(j) != 0 {
					for _, vp := range j {
						newP[vp] = vp
					}
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
	fmt.Println(len(newP))

	return newP
}

func PartTwo(c chart, freq map[string][]aloc) {
	maxRow := len(c.cMap)
	maxCol := len(c.cMap[0])
	t := [][]aloc{}
	j := map[aloc]int{}
	//part2 := 0

	for _, val := range freq {
		for _, p1 := range val {
			for _, p2 := range val {
				if p2 == p1 {
					continue
				} else {
					t = append(t, CheckPart2(p1, p2, maxRow, maxCol))
				}

			}
		}
	}

	for _, val := range t {
		for _, idx := range val {
			if j[idx] >= 1 {
				j[idx] += 1
			} else {
				j[idx] = 1
			}
		}
	}
	//fmt.Println(freq)
	//fmt.Println(t)
	fmt.Println(len(j))
}

func CheckValidPoint(curr loc, iter loc, maxRow int, maxCol int, c *chart) []aloc {
	validPoints := []aloc{}

	testP := []aloc{
		{x: iter.x - (iter.x - curr.x), y: iter.y - (iter.y - curr.y)},
		{x: iter.x - (iter.x - curr.x), y: iter.y - (iter.y - curr.y)},
		{x: curr.x + (curr.x - iter.x), y: curr.y + (curr.y - iter.y)},
		{x: curr.x + (curr.x - iter.x), y: curr.y + (curr.y - iter.y)},
	}

	for _, p := range testP {
		if p.x >= 0 && p.y >= 0 && p.x < maxCol && p.y < maxRow {
			if c.Get(p.y, p.x) == "." {
				validPoints = append(validPoints, p)
			}
		}
	}

	return validPoints

}

func CheckPart2(curr aloc, iter aloc, maxRow, maxCol int) []aloc {
	var x2 int
	var y2 int
	output := []aloc{}
	x := curr.x
	y := curr.y

	x2 = iter.x - curr.x
	y2 = iter.y - curr.y
	//fmt.Println(iter, curr)
	//fmt.Println("x2 & y2 are: ", reflect.TypeOf(x2).Kind(), reflect.TypeOf(y2).Kind())
	//fmt.Println(curr.x-iter.x, iter.x-curr.x)
	for x > 0 && x < maxCol && y > 0 && y < maxRow {
		output = append(output, aloc{x, y})
		x -= x2
		y -= y2

	}
	x = iter.x
	y = iter.y
	for x >= 0 && x < maxCol && y >= 0 && y < maxRow {
		output = append(output, aloc{x, y})
		x += x2
		y += y2

	}
	//fmt.Println(len(output))
	return output
}

func (chrt chart) Get(row, col int) string {
	return chrt.cMap[row][col]
}
