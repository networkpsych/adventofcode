package main

import (
	"bufio"
	"fmt"
	"os"
)

type loc struct {
	r, c int
}

func main() {
	maze := [][]string{}

	file, err := os.Open("./day6.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		tmp := []string{}
		for i := 0; i < len(line); i++ {
			tmp = append(tmp, string(line[i]))
		}
		maze = append(maze, tmp)
	}

	visited, idx := PartOne(maze)
	PartTwo(maze, visited, idx)

}

func PartOne(maze [][]string) ([]loc, loc) {
	var sIdx loc
	var output []loc
	var total int

	for row, line := range maze {
		for col := 0; col < len(line); col++ {
			if line[col] == "^" {
				sIdx.r = row
				sIdx.c = col
			}
		}
	}
	output, total = Walk(maze, sIdx.r, sIdx.c, len(maze[0]))
	fmt.Println("PART ONE: ", total)
	return output, sIdx
}

func PartTwo(m [][]string, visited []loc, sIdx loc) {
	total := 0
	for _, idx := range visited {
		if (idx == sIdx) || m[idx.r][idx.c] == "#" {
			continue
		}
		modArr(&m, idx.r, idx.c, "#")
		if LoopedGuard(m, sIdx) {
			total += 1
		}
		modArr(&m, idx.r, idx.c, ".")
	}
	fmt.Println("PART TWO: ", total)
}

func modArr(arr *[][]string, row, col int, val string) {
	(*arr)[row][col] = val
}

func WallCheck(arr [][]string, row, col int, face *int) bool {
	if row <= len(arr)-1 && row >= 0 && col <= len(arr[0])-1 && col >= 0 {
		if string(arr[row][col]) == "#" {
			if (*face) >= 3 {
				(*face) = 0
			} else {
				(*face) += 1
			}
			return true
		}
	}
	return false
}

func Walk(m [][]string, row int, col int, wBounds int) ([]loc, int) {
	w := m //walking
	d := map[int]string{
		0: "^",
		1: ">",
		2: "v",
		3: "<",
	}
	visited := make(map[loc]bool)
	mDir := 0
	//total := 1
	completed := false
	for !completed {
		x := col
		y := row
		curr := loc{y, x}
		guard := d[mDir]
		if guard == "^" && !WallCheck(w, y-1, x, &mDir) {
			row -= 1
		}
		if guard == ">" && !WallCheck(w, y, x+1, &mDir) {
			col += 1
		}
		if guard == "v" && !WallCheck(w, y+1, x, &mDir) {
			row += 1
		}
		if guard == "<" && !WallCheck(w, y, x-1, &mDir) {
			col -= 1
		}
		if row > len(m)-1 || row < 0 || col > len(m[0])-1 || col < 0 {
			visited[curr] = true
			completed = true
			break
		}
		modArr(&w, row, col, d[mDir])
		modArr(&w, y, x, "X")
		visited[curr] = true
	}
	for _, row := range w {
		fmt.Printf("%s\n", row)
	}
	vIdx := make([]loc, 0, len(visited))
	for index := range visited {
		vIdx = append(vIdx, index)
	}
	return vIdx, len(visited)
}

func LoopedGuard(m [][]string, sIdx loc) bool {
	// code based on https://github.com/arjunpathak072/aoc-2024/blob/main/day-6/main.go
	//
	visit := make(map[loc]loc)
	curr := sIdx
	d := []loc{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	mDir := 0
	hMax := len(m)
	lMax := len(m[0])

	for curr.r >= 0 && curr.c >= 0 && curr.r < hMax && curr.c < lMax {
		if visit[curr] == d[mDir] {
			return true
		}

		visit[curr] = d[mDir]

		if m[curr.r][curr.c] == "#" {
			curr.r -= d[mDir].r
			curr.c -= d[mDir].c
			mDir = (mDir + 1) % len(d)
			continue
		}
		modArr(&m, curr.r, curr.c, "X")
		curr.r += d[mDir].r
		curr.c += d[mDir].c
	}

	return false
}
