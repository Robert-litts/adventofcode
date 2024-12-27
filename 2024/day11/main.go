package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day11.input", "Relative file path to use as input.")
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
	var rocks Rocks
	contents := string(bytes)
	num := strings.Fields(strings.TrimSpace(contents))
	numLine := []int{}

	// Convert the first part (left of the colon) to the target integer
	for _, val := range num {
		target, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			fmt.Println("Error converting target to integer:", err)
			return err
		}
		numLine = append(numLine, target)
	}

	for i := 0; i < 25; i++ {

		newNumLine := []int{}

		for i := 0; i < len(numLine); i++ {

			if checkEvenDigits(numLine[i]) {
				leftInt, rightInt, _ := splitInt(numLine[i])

				newNumLine = append(newNumLine, leftInt)
				newNumLine = append(newNumLine, rightInt)
				continue

			}

			if numLine[i] == 0 {
				newNumLine = append(newNumLine, 1)

				continue
			}
			newNumLine = append(newNumLine, numLine[i]*2024)

		}

		numLine = newNumLine

	}
	rock_total := 0
	for _, rock := range rocks.Rocks {
		rock_total += rock.Count
	}

	fmt.Println("Rocks: ", rock_total)
	fmt.Println("Rock Length: ", len(numLine))

	fmt.Println("Result:", result)
	return nil
}

func Part2(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	//result := 0
	var rocks Rocks
	contents := string(bytes)
	num := strings.Fields(strings.TrimSpace(contents))

	// Convert the first part (left of the colon) to the target integer
	for _, val := range num {
		target, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			fmt.Println("Error converting target to integer:", err)
			return err
		}

		rocks.AddRock(target, 1)

	}
	fmt.Println("Initial Numbers:", rocks)

	for j := 0; j < 75; j++ {
		var newRocks Rocks
		fmt.Println("Blink Value: ", j)

		for i := 0; i < len(rocks.Rocks); i++ {

			if checkEvenDigits(rocks.Rocks[i].Value) {
				leftInt, rightInt, _ := splitInt(rocks.Rocks[i].Value)

				newRocks.AddRock(leftInt, rocks.Rocks[i].Count)
				newRocks.AddRock(rightInt, rocks.Rocks[i].Count)

				continue

			}

			if rocks.Rocks[i].Value == 0 {
				newRocks.AddRock(1, rocks.Rocks[i].Count)
				continue

			}
			newRocks.AddRock(rocks.Rocks[i].Value*2024, rocks.Rocks[i].Count)

		}

		rocks = newRocks

	}

	rock_total := 0
	for _, rock := range rocks.Rocks {
		rock_total += rock.Count
	}

	fmt.Println("Rocks: ", rock_total)

	return nil
}

type Rock struct {
	Value int
	Count int
}

// Struct to represent a list of Rocks
type Rocks struct {
	Rocks []Rock
}

func (r *Rocks) AddRock(value, count int) {
	// Iterate over the existing rocks to find if it already exists
	for i := range r.Rocks {
		if r.Rocks[i].Value == value {
			// If found, increment the count
			r.Rocks[i].Count += count
			return
		}

	}
	// If no rock with value found, append a new one with count equal to count
	r.Rocks = append(r.Rocks, Rock{Value: value, Count: count})
}

func checkEvenDigits(n int) bool {
	// Convert the number to a string
	str := strconv.Itoa(n)
	return len(str)%2 == 0
}

func splitInt(n int) (int, int, error) {
	str := strconv.Itoa(n)

	// Check if the number is even
	if len(str)%2 != 0 {
		// Return an error if the number is odd
		return 0, 0, errors.New("Number is not even")
	}

	// Count the length of the string
	length := len(str)
	// If the length is less than 2, we can't split it
	if length < 2 {
		return 0, 0, errors.New("Number is too small to split")
	}

	// Find the middle of the string
	half := length / 2
	// Split the string into two halves
	leftStr := str[:half]
	rightStr := str[half:]

	// Convert back to int
	leftInt, err := strconv.Atoi(leftStr)
	if err != nil {
		return 0, 0, err
	}
	rightInt, err := strconv.Atoi(rightStr)
	if err != nil {
		return 0, 0, err
	}

	// Return the two integers and a nil error
	return leftInt, rightInt, nil
}
