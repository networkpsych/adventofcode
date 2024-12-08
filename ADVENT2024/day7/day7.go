package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type sets struct {
	total  int
	args   []int
	length int
}

func main() {
	prog_start := time.Now()
	file, err := os.Open("./day7.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	var inst []sets
	for s.Scan() {
		l := strings.Replace(s.Text(), ":", "", -1)
		t := strings.Split(l, " ")
		j := Convert(t)
		inst = append(inst, j)
	}

	start := time.Now()
	PartOne(inst)
	ela := time.Now()
	fmt.Println("Time: ", ela.Sub(start))
	start = time.Now()
	PartTwo(inst)
	ela = time.Now()
	fmt.Println("Time: ", ela.Sub(start))
	fmt.Println("Program time: ", ela.Sub(prog_start))
}

func PartOne(set []sets) {

	t := 0
	for _, val := range set {
		//fmt.Println(val.total)
		if Backwards(val, val.total, val.length-1, false) {
			//fmt.Println(val.total)
			t += val.total
		}
	}
	fmt.Println("PART ONE: ", t)
}

func PartTwo(set []sets) {
	t := 0
	for _, val := range set {
		if Backwards(val, val.total, val.length-1, true) {
			t += val.total
		}
	}
	fmt.Println("PART TWO: ", t)
}

func Convert(s []string) sets {
	holder := []int{}

	for _, val := range s {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		holder = append(holder, num)
	}

	convOut := sets{
		total:  holder[0],
		args:   holder[1:],
		length: len(holder[1:]),
	}
	return convOut
}

func Backwards(set sets, curr int, idx int, concat bool) bool {
	/*
		Algorithm learned from /u/Odd-Statician7023
		Take last number in array. Check recursively if that number can be subtracted
		or divided into the total. If it cannot, then that branch can be pruned.
		If it does, then continue along each index, until you cannot prune, or the current
		total is equal to zero.
	*/

	if idx == 0 {
		return curr == set.args[idx]
	}
	if curr > set.args[idx] && Backwards(set, curr-set.args[idx], idx-1, concat) {
		return true
	}
	if curr%set.args[idx] == 0 && Backwards(set, curr/set.args[idx], idx-1, concat) {
		return true
	}

	if concat {
		divided := math.Pow(10, math.Floor(math.Log10(float64(set.args[idx])))+1)

		if curr%int(divided) == set.args[idx] && Backwards(set, curr/int(divided), idx-1, concat) {
			return true
		}
	}
	return false
}
