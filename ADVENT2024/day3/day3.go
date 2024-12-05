package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

/*
PART ONE
1. Filter results to find mul(X,Y)
	a. Valid results are only mul(X,Y).
	b. Valid results are encases within parentheses.
2. Add results of each mul(X,Y)
	(X*Y + X2*Y2 + ... Xn*Yn)

*/

func main() {
	//file, err := os.Open("./input.txt")
	file, err := os.ReadFile("./day3.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(file))
	//defer file.Close()
	pOne := FilterFile(string(file))
	pOne_nums := NumFilter(pOne)
	pTwo := FilterFile_p2(file)
	pTwo_nums := NumFilter(pTwo)
	AddValues(pOne_nums, "Part One")
	AddValues(pTwo_nums, "Part Two")
}

func FilterFile(s string) []string {
	f, _ := regexp.Compile("mul\\(\\d*,\\d*\\)")
	//fmt.Println(s)
	return f.FindAllString(s, -1)
	//fmt.Println(strings.Split(results, "mul"))
}

func FilterFile_p2(s []byte) []string {
	var output []string
	flag := true
	f := regexp.MustCompile("mul\\(\\d*,\\d*\\)|do\\(\\)|don't\\(\\)")
	t := f.FindAll(s, -1)
	for _, val := range t {
		if string(val) == "do()" {
			flag = true
		} else if string(val) == "don't()" {
			flag = false
		}
		fmt.Println(flag, string(val))
		if flag {
			if string(val) == "do()" {

				continue
			} else {
				output = append(output, string(val))
			}
		}
	}
	return output
}

func NumFilter(s []string) map[int][]int {
	nums := make(map[int][]int)
	f := regexp.MustCompile("\\D")
	for i, val := range s {
		var numArr []int
		tmp := f.Split(val, -1)
		//fmt.Println(tmp)
		//tmp = strings.Split(tmp, ",")
		for j := 0; j < len(tmp); j++ {
			//fmt.Println(tmp[j])
			n, err := strconv.Atoi(tmp[j])
			if err != nil {

			} else {
				numArr = append(numArr, n)
			}
		}
		nums[i] = numArr
	}
	return nums
}

func AddValues(n map[int][]int, s string) {
	total := 0
	for _, val := range n {
		total += (val[0] * val[1])
	}
	fmt.Printf("%s :: %d\n", s, total)
}
