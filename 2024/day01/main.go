package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running Part 1:")
	column1, column2 := readInputFile() // Read input file and get the columns
	Part1(column1, column2)             // Pass columns to Part1

	Part2(column1, column2) // Pass columns to Part2
}

func readInputFile() ([]int, []int) {
	var inputFile = flag.String("inputFile", "../input/day1.input", "Relative file path to use as input.")
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}
	contents := string(bytes)

	var column1 []int
	var column2 []int

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		parts := strings.Fields(line)

		// Convert the first part to an integer and append to column1
		if len(parts) >= 2 {
			firstColumn, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting first column value:", err)
				continue
			}
			column1 = append(column1, firstColumn)

			// Convert the second part to an integer and append to column2
			secondColumn, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting second column value:", err)
				continue
			}
			column2 = append(column2, secondColumn)
		}
	}
	return column1, column2
}

func Part1(column1, column2 []int) {
	// Sort the columns low to high
	sort.Ints(column1)
	sort.Ints(column2)

	dif := 0
	for i, each := range column1 {
		dif += max((each - column2[i]), (column2[i] - each))
	}
	fmt.Println("Part 1 Total:", dif)
}

func Part2(column1, column2 []int) {
	// Sort the columns low to high
	sort.Ints(column1)
	sort.Ints(column2)

	similarity := 0
	for _, each := range column1 {
		count := 0
		for _, k := range column2 {
			if each == k {
				count++
			}
		}
		similarity += each * count
	}
	fmt.Println("Part 2 Total:", similarity)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
