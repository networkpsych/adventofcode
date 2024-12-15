package helper

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Coor struct {
	x, y int
}

func LoadFile(filePath string) ([]string, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	output := []string{}
	for s.Scan() {
		output = append(output, s.Text())
	}

	return output, nil

}

func LoadFileStringChart(filePath string, delim string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	output := [][]string{}
	for s.Scan() {
		tmp := strings.Split(s.Text(), delim)
		output = append(output, tmp)
	}

	return output, nil
}

func LoadFileIntChart(filePath string, delim string) ([][]int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	output := [][]int{}
	for s.Scan() {
		line := []int{}
		for _, item := range strings.Split(s.Text(), delim) {
			tmp, _ := strconv.Atoi(item)
			line = append(line, tmp)

		}
		output = append(output, line)
	}

	return output, nil
}

/*
func LoadFileCoor(filePath string, delim string) ([][]Coor, error){
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	output := [][]Coor{}
	for s.Scan() {
		line := []int{}
		for _, item := range  {
			tmp, _ := strconv.Atoi(item)
			line = append(line, tmp)

		}
		output = append(output, line)
	}

	return output, nil
}
*/
