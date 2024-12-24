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
	var inputFile = flag.String("inputFile", "../input/day17.input", "Relative file path to use as input.")
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
	regA := 0
	regB := 0
	regC := 0
	program := []int{}
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	contents := string(bytes)

	lines := strings.Split(contents, "\n")

	// Iterate through lines in blocks of 3
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i]) // Remove spaces
		if i < 3 {                          // Skip first 3 lines
			fmt.Sscanf(line, "Register A: %d", &regA)
			fmt.Sscanf(line, "Register B: %d", &regB)
			fmt.Sscanf(line, "Register C: %d", &regC)
		}

		if line == "" {
			continue
		}
		if i >= 3 { // Start reading the program
			result := strings.Split(line, ":")
			programString := strings.Split(result[1], ",")
			for k := range len(programString) {
				programString[k] = strings.TrimSpace(programString[k])
				num, err := strconv.Atoi(programString[k])
				if err != nil {
					return err
				}
				program = append(program, num)
			}

		}
	}

	fmt.Println("Program:", program)
	fmt.Println("Registers:", regA, regB, regC)

	return nil
}

func adv(regA, comboOperand float64) (int, error) {
	result := int(regA / (math.Pow(2, comboOperand)))

	return result, nil

}

func bxl(regB, litOperand int) (int, error) {
	result := regB ^ litOperand
	return result, nil
}

func bst(comboOperand int) (int, error) {
	result := comboOperand % 8
	return result, nil
}
