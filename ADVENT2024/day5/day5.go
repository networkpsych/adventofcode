package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
	/*
		1. Page ordering guide of X|Y
		2. First section specifies ordering rule
			X|Y <- Page ordering
			X needs to occur at some point before Y
		3. Second section specifies the page numbers of each update.
			i.e., N, N+1, N+2, ... N+N <- Page numbers

		4. Check which updates are in the right order.
	*/
	sect1 := [][]int{}
	sect2 := [][]int{}

	//file, err := os.Open("./test.txt")
	file, err := os.Open("./day5.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	f := bufio.NewScanner(file)
	swch := false
	for f.Scan() {
		if string(f.Text()) == "" {
			fmt.Println("switching")
			swch = true
			continue
		}
		if !swch {
			tmp := strings.Split(f.Text(), "|")
			x, _ := strconv.Atoi(tmp[0])
			y, _ := strconv.Atoi(tmp[1])
			tmpL := []int{x, y}
			sect1 = append(sect1, tmpL)
		} else {
			tmp := strings.Split(f.Text(), ",")
			tmpL := []int{}
			for _, n := range tmp {
				num, _ := strconv.Atoi(n)
				tmpL = append(tmpL, num)
			}
			sect2 = append(sect2, tmpL)
		}
	}

	PartOne(sect1, sect2)
	PartTwo(sect1, sect2)
}

func Validate(f [][]int, line []int) bool {
	for _, rule := range f {
		X := slices.Index(line, rule[0])
		Y := slices.Index(line, rule[1])
		if X == -1 || Y == -1 {
			continue
		} else if X > Y {
			return false
		}
	}
	return true
}

func ReOrder(f [][]int, line []int, idx int) []int {
	// this is a very stupid algo.
	copy := line
	for !Validate(f, copy) {
		for _, rule := range f {
			for i := 0; i < len(line); i++ {
				for j := i + 1; j < len(line); j++ {
					curr := []int{copy[j], copy[i]}
					if reflect.DeepEqual(curr, rule) {
						copy[j], copy[i] = copy[i], copy[j]
					}
				}
			}
		}
	}
	return copy
}

func PartOne(f [][]int, s [][]int) {
	total := 0

	for _, val := range s {
		if Validate(f, val) {
			total += val[len(val)/2]
		}
	}
	fmt.Println("Part One: ", total)
}

func PartTwo(f [][]int, s [][]int) {
	total := 0
	holder := [][]int{}
	for _, line := range s {
		if !Validate(f, line) {
			holder = append(holder, line)
		}
	}
	for idx, line := range holder {
		n := ReOrder(f, line, idx)
		total += n[len(n)/2]
	}
	fmt.Println("Part Two: ", total)
}
