package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	var col1, col2 []int

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		row := make([]int, 2)
		if len(parts) == 2 {
			row[0], err = strconv.Atoi(parts[0])
			row[1], err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			col1 = append(col1, row[0])
			col2 = append(col2, row[1])
		}
	}

	return col1, col2, nil
}

func getDistance(col1, col2 []int) int {
	ans := 0
	sort.Ints(col1)
	sort.Ints(col2)
	for i := 0; i < len(col1); i++ {
		ans += max(col1[i], col2[i]) - min(col1[i], col2[i])
	}

	return ans
}

func getSimilarityScore(col1, col2 []int) int {
	frequencyMapCol2 := make(map[int]int)

	for _, val := range col2 {
		frequencyMapCol2[val]++
	}

	ans := 0
	for _, val := range col1 {
		ans += val * frequencyMapCol2[val]
	}

	return ans
}

func main() {
	col1, col2, err := readInput("input.txt")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(getDistance(col1, col2))
	fmt.Println(getSimilarityScore(col1, col2))
}
