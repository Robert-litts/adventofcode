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
	var inputFile = flag.String("inputFile", "../input/day3.input", "Relative file path to use as input.")
	flag.Parse()
	fmt.Println("Running Part 1:")
	if err := Part1(*inputFile); err != nil {
		fmt.Println("Error in Part 2:", err)
		return
	}
	fmt.Println("Running Part 2:")
	if err := Part2(*inputFile); err != nil {
		fmt.Println("Error in Part 2:", err)
		return
	}
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

func Part2(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	result := 0
	contents := string(bytes)
	pattern_dont := `(?:^|[^a-zA-Z])don't\(\)`
	pattern_mul := `mul\((-?\d+(\.\d+)?),(-?\d+(\.\d+)?)\)`
	pattern_do := `do\(\)`
	re_dont := regexp.MustCompile(pattern_dont)
	re_mul := regexp.MustCompile(pattern_mul)
	re_do := regexp.MustCompile(pattern_do)
	do := true

	index_dont := re_dont.FindAllStringIndex(contents, -1)
	index_mul := re_mul.FindAllStringIndex(contents, -1)
	index_do := re_do.FindAllStringIndex(contents, -1)
	dont_index := 0
	do_index := 0
	mul_index := 0

	// loop thru each char in the line until the end
	for i := 0; i < len(contents); {

		//If we reach a dont,
		if dont_index < len(index_dont) && i == index_dont[dont_index][0] {
			do = false
			i = index_dont[dont_index][1]
			dont_index++
			continue
		}

		//If we reach a do
		if do_index < len(index_do) && i == index_do[do_index][0] {
			do = true
			i = index_do[do_index][1]
			do_index++
			continue
		}

		// Check for mul()
		if mul_index < len(index_mul) && i == index_mul[mul_index][0] {
			//If do is enabled
			if do {
				mulMatch := re_mul.FindStringSubmatch(contents[index_mul[mul_index][0]:index_mul[mul_index][1]])
				if len(mulMatch) > 0 {
					//fmt.Println("MulMatch: ", mulMatch)
					var x, y int
					fmt.Sscanf(mulMatch[1], "%d", &x)
					fmt.Sscanf(mulMatch[3], "%d", &y)
					result += x * y
					//fmt.Println("Result: ", result)
				}
			}
			//If do is disabled, skip the mull...move i to the end of current mul & increment mul_index to the next mul
			i = index_mul[mul_index][1]
			mul_index++
			continue
		}
		i++

	}

	fmt.Println("Result Part 2:", result)
	return nil
}
