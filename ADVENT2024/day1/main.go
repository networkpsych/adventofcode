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
	var left []int
	var right []int
	data, err := os.Open("./day1.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	s := bufio.NewScanner(data)
	i := 0
	for s.Scan() {
		l := strings.Split(s.Text(), ",")
		num1, err := strconv.Atoi(l[0])
		if err != nil {
			fmt.Println(err)
		}
		num2, err := strconv.Atoi(l[1])
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, num1)
		right = append(right, num2)
		i++
	}

	slices.Sort(left)
	slices.Sort(right)
	PartTwo(left, right)
	if err := s.Err(); err != nil {
		panic(err)
	}
}

func PartOne(l []int, r []int) {
	var t []int
	var tmp int

	for j := 0; j < len(l); j++ {
		tmp = l[j] - r[j]
		if tmp < 0 {
			tmp = tmp * -1
		}
		t = append(t, tmp)
	}

	//fmt.Println(total)

	total := 0
	for k := 0; k < len(t); k++ {
		total += t[k]
	}
	fmt.Println(total)
}

func PartTwo(l []int, r []int) {
	r_count := Counter(r)

	total := 0
	for _, val := range l {

		_, ok := r_count[val]

		if ok {
			total += (val * r_count[val])
		}
	}

	fmt.Println(total)

}

func Counter(item []int) map[int]int {
	tmp := make(map[int]int, len(item))

	for _, val := range item {
		_, ok := tmp[val]

		if ok {
			tmp[val] += 1
		} else {
			tmp[val] = 1
		}
	}

	return tmp
}
