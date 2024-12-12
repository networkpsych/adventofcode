package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day9.txt")
	if err != nil {
		panic(err)
	}
	PartOne(strings.Split(string(file), ""))
	PartTwo(strings.Split(string(file), ""))
}

func PartOne(file []string) {
	total := int64(0)
	output := []string{}
	for idx, char := range file {
		numSp, _ := strconv.ParseInt(string(char), 10, 32)
		if idx%2 == 0 {
			val := strconv.Itoa(idx / 2)
			for i := 0; i < int(numSp); i++ {
				output = append(output, val)
			}
		} else {
			for j := 0; j < int(numSp); j++ {
				output = append(output, ".")
			}
		}
	}
	for i, j := 0, len(output)-1; i <= j; {
		if output[i] == "." && output[j] != "." {
			output[i], output[j] = output[j], output[i]
			i++
			j--
		} else {
			if output[i] != "." {
				i++
			} else if output[j] != "." {
				j--
			} else if output[i] == "." && output[j] == "." {
				j--
			}
		}
	}

	for idx := range output {
		//fmt.Printf("%v", output[idx])
		if output[idx] != "." {
			n, _ := strconv.ParseInt(output[idx], 10, 32)
			total += int64(idx) * n
		}
	}
	fmt.Println(total)

}

func PartTwo(file []string) {
	total := int64(0)
	dMap := []string{}
	memIdx := [][]int64{}
	freeIdx := [][]int64{}

	for idx, char := range file {
		if idx%2 == 0 {
			memItem, _ := strconv.ParseInt(char, 10, 32)
			val := strconv.Itoa(idx / 2)
			memIdx = append(memIdx, []int64{int64(len(dMap)), int64(len(dMap)-1) + memItem})
			for i := 0; i < int(memItem); i++ {
				dMap = append(dMap, val)
			}
		} else {
			spaceItem, _ := strconv.ParseInt(char, 10, 32)
			freeIdx = append(freeIdx, []int64{int64(len(dMap)), int64(len(dMap)-1) + spaceItem})
			for i := 0; i < int(spaceItem); i++ {
				dMap = append(dMap, ".")
			}
		}
	}

	for m := (len(memIdx) - 1); m >= 0; m-- {
		for f := 0; f < len(freeIdx)-1; f++ {
			mem := memIdx[m]
			free := freeIdx[f]
			if free[1] < mem[1] && free[1]-free[0] >= mem[1]-mem[0] {
				for item := free[0]; item <= free[0]+mem[1]-mem[0]; item++ {
					dMap[item] = strconv.Itoa(m)
				}
				for o := mem[0]; o <= mem[1]; o++ {
					dMap[o] = "."
				}
				free[0] = free[0] + mem[1] - mem[0] + 1
				break
			}
		}
	}

	for idx, char := range dMap {
		if char != "." {
			num, _ := strconv.ParseInt(char, 10, 32)
			total += int64(idx) * num
		}

	}
	fmt.Println(total)
	//fmt.Println(dMap)
	//fmt.Println(memIdx)
	//fmt.Println(freeIdx)

}
