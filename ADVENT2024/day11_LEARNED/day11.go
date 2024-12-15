package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
	1. If a stone has a 0, then it is replaced with a 1
	2. If a stone has an even number of digits:
		a. Left half of digits goes to a new stone.
		b. Right half of digits goes to a new stone.
		c. Numbers do not keep any leading zeroes.
	3. If neither rule apply, multiple value by 2024
		and engrave on a new stone
Use a sync WaitGroup to track existing items during the recursion.
If an item already exists, you can count the recurrance of a stone.
Solution credit: https://github.com/omotto/AdventOfCode2024/blob/main/src/day11/main.go
Reddit Comment: https://tinyurl.com/38cny3vp
Golang Sync package: https://pkg.go.dev/sync
	Sync WaitGroup: https://pkg.go.dev/sync#WaitGroup
*/

func main() {
	file, err := os.Open("day11.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	inFile := []string{}
	for s.Scan() {
		inFile = append(inFile, s.Text())
	}

	//fmt.Println(inFile)
	fmt.Println("Test @25 iterations: ", Start(inFile, 25))
	fmt.Println("Test @75 iterations: ", Start(inFile, 75))
}

func Start(v []string, iterations int) int {
	g := sync.WaitGroup{}
	checked := &sync.Map{}
	item := strings.Split(v[0], " ")
	nums := make([]int, len(item))
	channel := make(chan int, len(nums))
	for idx, val := range item {
		g.Add(1)
		nums[idx], _ = strconv.Atoi(val)
		go func(stone int) {
			defer g.Done()
			channel <- Stones(stone, iterations, checked)
		}(nums[idx])
	}
	g.Wait()
	close(channel)
	r := 0
	for val := range channel {
		//fmt.Println(val)
		r += val
	}
	return r
}

func Stones(val int, iter int, checked *sync.Map) int {
	// recursion for each value in the arr.
	r := 0
	if iter > 0 {
		if v, ok := checked.Load(fmt.Sprintf("%d:%d", val, iter)); ok {
			r += v.(int)
		} else {

			if val == 0 {
				r += Stones(1, iter-1, checked)
			} else if d := strconv.Itoa(val); len(d)%2 == 0 {
				left, _ := strconv.Atoi(d[:len(d)/2])
				right, _ := strconv.Atoi(d[len(d)/2:])
				r += Stones(left, iter-1, checked)
				r += Stones(right, iter-1, checked)

			} else {
				r += Stones(val*2024, iter-1, checked)
			}

			checked.Store(fmt.Sprintf("%d:%d", val, iter), r)
		}
	} else {
		return 1
	}

	return r

}
