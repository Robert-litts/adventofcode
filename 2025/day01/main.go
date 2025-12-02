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
	var inputFile = flag.String("inputFile", "../input/day1.input", "Relative file path to use as input.")
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
		numStr := turns[1]

		num_turns, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		if num_turns > 99 {
			num_turns = num_turns % 100
		}
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
		//fmt.Println("Position: ", position)
	}
	fmt.Println("Result:", password)
	return nil
}

func Part2(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	pattern := `[RL]`
	re := regexp.MustCompile(pattern)

	start_position := 50
	position := 50
	password := 0
	zero_passed := 0

	for _, line := range lines {
		direction := re.FindAllString(line, -1)[0]
		turns := re.Split(line, -1)
		numStr := turns[1]
		num_turns, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		//Handle large turn values
		if num_turns > 99 {
			zero_passed += num_turns / 100 //retain total number full turns completed
			num_turns = num_turns % 100    //reduce to within a single turn via modulo
		}

		if direction == "R" {
			position = position + num_turns
		} else {
			position = position - num_turns
		}

		if position > 99 {
			position = position - 100

			//Handle case where it doesn't land on 0, but also didnt start on 0
			if position != 0 && start_position != 0 {
				zero_passed += 1
			}
		} else if position < 0 {
			position = 100 + position

			//Handle case where it doesn't land on 0, but also didnt start on 0 (to avoid double counting)
			if position != 0 && start_position != 0 {
				zero_passed += 1
			}

		}
		//Handle case where it stops on 0
		if position == 0 {
			password += 1
		}
		//reset starting position
		start_position = position
	}
	total_score := password + zero_passed
	fmt.Println("Result:", total_score)
	return nil
}
