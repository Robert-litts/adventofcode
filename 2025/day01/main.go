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
	var inputFile = flag.String("inputFile", "../input/day1.ex.input", "Relative file path to use as input.")
	flag.Parse()
	fmt.Println("Running Part 1:")
	if err := Part1(*inputFile); err != nil {
		fmt.Println("Error in Part 1:", err)
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
	//result := 0
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	pattern := `[RL]`
	re := regexp.MustCompile(pattern)

	position := 50
	password := 0

	for _, line := range lines {
		direction := re.FindAllString(line, -1)[0]
		turns := re.Split(line, -1)
		fmt.Println(direction)
		numStr := turns[1]
		fmt.Println(numStr)

		num_turns, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		if num_turns > 99 {
			num_turns = num_turns % 100
		}
		fmt.Println(num_turns)

		if direction == "R" {
			position = position + num_turns
		} else {
			position = position - num_turns
		}

		if position > 99 {
			position = position - 100
		} else if position < 0 {
			position = 100 + position

		}
		if position == 0 {
			password += 1
		}
		fmt.Println("Position: ", position)
	}
	//}
	fmt.Println("Result:", password)
	return nil
}

func Part2(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	result := 0

	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		fmt.Println(line)

		fmt.Println("Result Part 2:", result)
	}
	return nil
}
