package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day22.input", "Relative file path to use as input.")
	flag.Parse()
	start := time.Now()
	fmt.Println("Running Part 1:")
	if err := Part1(*inputFile); err != nil {
		fmt.Println("Error in Part 2:", err)
		return
	}
	duration := time.Since(start)
	fmt.Printf("Execution Time: %s\n", duration)
}

func Part1(inputFile string) error {
	nums := []int{}
	secretsOut := []int{}
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	depth := 2000

	contents := string(bytes)

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		nums = append(nums, number)
	}
	fmt.Println(nums)

	for _, num := range nums {
		secret := num
		for i := 0; i < depth; i++ {

			secret = findSecret(secret)
		}
		secretsOut = append(secretsOut, secret)
		//fmt.Println("2000th secret of ", num, " is ", secret)
	}

	fmt.Println(secretsOut)
	result := 0
	for j := range secretsOut {
		result += secretsOut[j]
	}
	fmt.Println("Part 1: ", result)
	return nil
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Coordinate struct {
	X int
	Y int
}

func findSecret(inSecret int) int {

	// Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
	// Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
	// Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
	step1 := ((inSecret * 64) ^ inSecret) % 16777216
	step2 := (int(math.Floor(float64(step1)/32)) ^ step1) % 16777216
	step3 := ((step2 * 2048) ^ step2) % 16777216

	return step3
}
