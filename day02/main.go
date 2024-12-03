package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day2.input", "Relative file path to use as input.")
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

func Part2(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var safe int
	for _, line := range lines {
		parts := strings.Fields(line)
		inc_count := 0
		dec_count := 0
		dif := 0
		var decreasing, increasing bool
		for i := 0; i < len(parts)-1; i++ {
			j := parts[i]
			jInt, err := strconv.Atoi(j)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return err
			}

			partsInt, err := strconv.Atoi(parts[i+1])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return err
			}
			if max(jInt-partsInt, partsInt-jInt) < 1 || max(jInt-partsInt, partsInt-jInt) > 3 {
				dif++
			}
			if jInt < partsInt {
				inc_count++
			} else if jInt > partsInt {
				dec_count++
			}
		}

		if inc_count == len(parts)-1 && dec_count == 0 {
			increasing = true
		} else if dec_count == len(parts)-1 && inc_count == 0 {
			decreasing = true
		}

		if dif == 0 && ((increasing && !decreasing) || (decreasing && !increasing)) {
			safe++
			//fmt.Println("safe")
		} else {
			safeCount, err := fixSafe(line)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			safe += safeCount
		}

	}
	fmt.Println("Safety Number Part 2: ", safe)
	return nil
}

func fixSafe(line string) (int, error) {
	safe := 0
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			return 0, fmt.Errorf("error converting string to int: %v", err)
		}
		numbers[i] = num
	}
	//Remove each number and check the remaining sequence
	for i := 0; i < len(numbers); i++ {
		remaining := make([]int, 0, len(numbers)-1)
		remaining = append(remaining, numbers[:i]...)
		remaining = append(remaining, numbers[i+1:]...)
		//fmt.Printf("Testing sequence without %d: %v\n", numbers[i], remaining)

		inc_count := 0
		dec_count := 0
		dif := 0
		var decreasing, increasing bool

		for k := 0; k < len(remaining)-1; k++ {

			if max(remaining[k]-remaining[k+1], remaining[k+1]-remaining[k]) < 1 || max(remaining[k]-remaining[k+1], remaining[k+1]-remaining[k]) > 3 {
				dif++
			}
			if remaining[k] < remaining[k+1] {
				inc_count++
			} else if remaining[k] > remaining[k+1] {
				dec_count++
			}
		}

		if inc_count == len(remaining)-1 && dec_count == 0 {
			increasing = true
		} else if dec_count == len(remaining)-1 && inc_count == 0 {
			decreasing = true
		}

		if dif == 0 && ((increasing && !decreasing) || (decreasing && !increasing)) {
			safe++

		}
		//If iterating through the different removal scenarios produces exactly 1 safe result, return 1, otherwise return 0
		if safe == 1 {
			return 1, nil
		}
	}
	return 0, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Part1(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		//fmt.Println("Error reading file:", err)
		return err
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var safe int
	for _, line := range lines {
		parts := strings.Fields(line)
		inc_count := 0
		dec_count := 0
		dif := 0
		var decreasing, increasing bool
		for i := 0; i < len(parts)-1; i++ {
			j := parts[i]
			jInt, err := strconv.Atoi(j)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return err
			}

			partsInt, err := strconv.Atoi(parts[i+1])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return err
			}
			if max(jInt-partsInt, partsInt-jInt) < 1 || max(jInt-partsInt, partsInt-jInt) > 3 {
				dif++
			}
			if jInt < partsInt {
				inc_count++
			} else if jInt > partsInt {
				dec_count++
			}
		}
		if inc_count == len(parts)-1 && dec_count == 0 {
			increasing = true
		} else if dec_count == len(parts)-1 && inc_count == 0 {
			decreasing = true
		}

		if dif == 0 && ((increasing && !decreasing) || (decreasing && !increasing)) {
			safe++
			//fmt.Println("safe")
		} else {
			//fmt.Println("not safe")
		}

	}
	fmt.Println("Safety Number Part 1: ", safe)
	return nil
}
