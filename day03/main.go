package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day3.example", "Relative file path to use as input.")
	flag.Parse()
	fmt.Println("Running Part 1:")
	if err := Part1(*inputFile); err != nil {
		fmt.Println("Error in Part 2:", err)
		return
	}
	// fmt.Println("Running Part 2:")
	// if err := Part2(*inputFile); err != nil {
	// 	fmt.Println("Error in Part 2:", err)
	// 	return
	// }
}

func Part1(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	result := 0
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	pattern := `mul\((-?\d+(\.\d+)?),(-?\d+(\.\d+)?)\)`
	pattern2 := `\d+`
	re := regexp.MustCompile(pattern)
	re2 := regexp.MustCompile(pattern2)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				matches2 := re2.FindAllString(match, -1)
				store := []int{}
				for _, numStr := range matches2 {

					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("Error converting string to int:", err)
					} else {
						// Store the number in the slice
						store = append(store, num)
					}
				}
				multiply := store[0] * store[1]
				result += multiply
				//fmt.Println("Multiplication result:", multiply)
			}
		}
	}
	fmt.Println("Result:", result)
	return nil
}
