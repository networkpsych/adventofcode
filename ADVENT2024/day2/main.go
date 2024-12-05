package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	/*
		1. Level is only increasing or decreasing
		2. any two adjacent levels are 1 - 3 in difference
	*/
	nums := make(map[int][]int, 1000)
	data, err := os.Open("./day2.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	s := bufio.NewScanner(data)
	idx := 0
	for s.Scan() {
		var tmpArr []int
		curr_line := strings.Split(s.Text(), " ")
		for _, i := range curr_line {
			tmp, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
			}
			tmpArr = append(tmpArr, tmp)
		}
		nums[idx] = tmpArr
		idx++
	}
	//PartOne(nums)
	PartTwo(nums)
}

func PartOne(items map[int][]int) {
	total := 0
	for _, val := range items {
		//fmt.Println(val)
		valid := false
		for i := 0; i < len(val); i++ {
			if i+1 >= len(val) {
				break
			} else {
				//fmt.Printf("%d <--> %d\n", val[i], val[i+1])
				if i == 0 {
					if val[i] < val[i+1] {
						if val[i+1]-val[i] > 3 || val[i+1]-val[i] < 1 {
							valid = false
							break
						}
					} else if val[i] > val[i+1] {
						if val[i]-val[i+1] > 3 || val[i]-val[i+1] < 1 {
							valid = false
							break
						}
					} else {
						valid = false
						break
					}
				} else {
					if val[i] > val[i+1] && val[i] > val[i-1] || val[i] < val[i-1] && val[i] < val[i+1] {
						valid = false
						break
					} else if val[i] < val[i+1] {
						if val[i+1]-val[i] > 3 || val[i+1]-val[i] < 1 {
							valid = false
							break
						}
					} else if val[i] > val[i+1] {
						if val[i]-val[i+1] > 3 || val[i]-val[i+1] < 1 {
							valid = false
							break
						}
					} else {
						valid = false
						break
					}
				}
			}
			valid = true
		}
		if valid {
			total += 1
		}
	}
	fmt.Println(total)
}

func PartTwo(items map[int][]int) {
	//var valid bool
	total := 0

	for _, val := range items {
		if CheckValid(val) {
			total += 1
		} else if SecondaryCheck(val) {
			total += 1
		}
	}

	fmt.Println(total)
}

func CheckValid(items []int) bool {
	check := true
	for i := 1; i < len(items); i++ {
		diff := items[i] - items[i-1]
		if diff < 0 {
			diff *= -1
		}
		if diff > 3 || diff < 1 {
			check = false
		}
	}
	if !Increasing(items) && !Decreasing(items) {
		check = false
	}
	return check
}

func Increasing(items []int) bool {
	for i := 1; i < len(items); i++ {
		if items[i] <= items[i-1] {
			return false
		}
	}
	return true
}

func Decreasing(items []int) bool {
	for i := 1; i < len(items); i++ {
		if items[i] >= items[i-1] {
			return false
		}
	}
	return true
}

func NotRepeating(items []int) bool {
	slices.Sort(items)
	for i := 1; i < len(items); i++ {
		if items[i] == items[i-1] {
			return false
		}
	}
	return true
}

func SecondaryCheck(items []int) bool {
	for i := range items {
		n := make([]int, len(items))
		copy(n, items)
		slices.Delete(n, i, i+1)
		n = n[:len(n)-1]
		if CheckValid(n) {
			return true
		}
	}
	return false
}
