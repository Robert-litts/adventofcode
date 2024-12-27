package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day7.example", "Relative file path to use as input.")
	flag.Parse()
	start := time.Now()

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
	duration := time.Since(start)
	fmt.Printf("Execution Time: %s\n", duration)
}

func Part1(inputFile string) error {
	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		// Handle error in reading file
		return err
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	calibration := 0

	// Process each line in the file
	for _, line := range lines {
		// Trim any leading or trailing whitespace
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line at the colon
		split := strings.Split(line, ":")
		if len(split) < 2 {
			// If the line doesn't contain a colon, skip it
			continue
		}

		// Convert the first part (left of the colon) to the target integer
		target, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			fmt.Println("Error converting target to integer:", err)
			return err
		}
		// Split the second part (right of the colon) into the "check" values
		checkStrings := strings.Fields(strings.TrimSpace(split[1]))

		// Convert each "check" value from string to integer
		var check []int
		for _, str := range checkStrings {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting check value to integer:", err)
				return err
			}
			check = append(check, num)
		}

		// Print the target and check array for verification
		//fmt.Printf("Target: %d, Check: %v\n", target, check)

		// Call the recursive function to find the path, start with first value in check as the current value and the first index as 1
		// Start with is_addition = true, meaning "add" is the first operation
		if recursive_solve(target, check, check[0], 1) {
			calibration += target
		}
	}
	fmt.Println("Calibration Score, Part 1: ", calibration)

	return nil
}

func Part2(inputFile string) error {
	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		// Handle error in reading file
		return err
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	calibration := 0

	// Process each line in the file
	for _, line := range lines {
		// Trim any leading or trailing whitespace
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line at the colon
		split := strings.Split(line, ":")
		if len(split) < 2 {
			// If the line doesn't contain a colon, skip it
			continue
		}

		// Convert the first part (left of the colon) to the target integer
		target, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			fmt.Println("Error converting target to integer:", err)
			return err
		}
		// Split the second part (right of the colon) into the "check" values
		checkStrings := strings.Fields(strings.TrimSpace(split[1]))

		// Convert each "check" value from string to integer
		var check []int
		for _, str := range checkStrings {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting check value to integer:", err)
				return err
			}
			check = append(check, num)
		}

		// Print the target and check array for verification
		//fmt.Printf("Target: %d, Check: %v\n", target, check)

		// Call the recursive function to find the path, start with first value in check as the current value and the first index as 1
		// Start with is_addition = true, meaning "add" is the first operation
		if recursive_solve2(target, check, check[0], 1) {
			calibration += target
		}
	}
	fmt.Println("Calibration Score, Part 2: ", calibration)

	return nil
}

// Recursive function to find if any combination of operators (addition or multiplication) produces the target value
func recursive_solve(target int, numbers []int, current_value int, current_index int) bool {
	// If we've processed all numbers, check if the current value matches the target and if so, return true
	if current_index == len(numbers) {
		return current_value == target
	}

	// Get next num to check
	num := numbers[current_index]

	// Try both operations: addition and multiplication, return True if any pass, otherwise false
	result := recursive_solve(target, numbers, current_value+num, current_index+1) ||
		recursive_solve(target, numbers, current_value*num, current_index+1)

	// Return the result of the operations
	return result
}

// Recursive function  for Part 2 to find if any combination of operators (addition, multiplication, concat) produces the target value
func recursive_solve2(target int, numbers []int, current_value int, current_index int) bool {

	// If we've processed all numbers, check if the current value matches the target and if so, return true
	if current_index == len(numbers) {
		return current_value == target
	}

	// Get next num to check
	num := numbers[current_index]

	// Try 3 operations: addition, multiplication, concatenation, return True if any pass, otherwise false
	result := recursive_solve2(target, numbers, current_value+num, current_index+1) ||
		recursive_solve2(target, numbers, current_value*num, current_index+1) ||
		recursive_solve2(target, numbers, concat_ints(current_value, num), current_index+1)

	// If no valid solution found, return false
	return result
}

// Function to concat two integers
func concat_ints(a, b int) int {
	// Convert b to string
	b_str := strconv.Itoa(b)
	// Convert a to string
	a_str := strconv.Itoa(a)
	// Concatenate a and b
	concat := a_str + b_str
	// Convert concat to int
	concat_int, _ := strconv.Atoi(concat)
	// Return concat_int
	return concat_int
}
