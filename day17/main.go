package main

import (
	"errors"
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
	var jump bool
	program := []int{}
	output := []int{}
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
	//i = instruction pointer, increases by 2 except for jump instructions
	for i := 0; i < len(program); {
		litOperand := program[i+1]
		comboOperand, err := findComboOperand(program[i+1], regA, regB, regC)
		if err != nil {
			return err
		}

		switch {
		case program[i] == 0:
			regA, err = adv(regA, comboOperand)
			if err != nil {
				return err
			}
		case program[i] == 1:
			regB, err = bxl(regB, litOperand)
			if err != nil {
				return err
			}
		case program[i] == 2:
			regB, err = bst(comboOperand)
			if err != nil {
				return err
			}
		case program[i] == 3:
			jump = jnz(regA)
			if jump {
				i = litOperand
				continue
			}
		case program[i] == 4:
			regB, err = bxc(regB, regC)
			if err != nil {
				return err
			}
		case program[i] == 5:
			output = append(output, out(comboOperand))

		case program[i] == 6:
			regB, err = bdv(regA, comboOperand)
			if err != nil {
				return err
			}
		case program[i] == 7:
			regC, err = cdv(regA, comboOperand)
			if err != nil {
				return err
			}
		}
		i += 2
	}

	var strValues []string
	for _, val := range output {
		strValues = append(strValues, strconv.Itoa(val)) // Convert int to string
	}

	// Join the slice of strings with commas
	part1 := strings.Join(strValues, ",")

	fmt.Println("Program:", program)
	fmt.Println("Registers:", regA, regB, regC)
	fmt.Println("Output: ", output, part1)

	return nil
}

func adv(regA, comboOperand int) (int, error) {
	result := int(float64(regA) / (math.Pow(2, float64(comboOperand))))

	return result, nil

}

func bdv(regA, comboOperand int) (int, error) {
	result := int(float64(regA) / (math.Pow(2, float64(comboOperand))))

	return result, nil

}

func cdv(regA, comboOperand int) (int, error) {
	result := int(float64(regA) / (math.Pow(2, float64(comboOperand))))

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

func out(comboOperand int) int {
	result := comboOperand % 8
	return result
}

func bxc(regB, regC int) (int, error) {
	result := regB ^ regC
	return result, nil
}

func jnz(regA int) bool {
	if regA != 0 {
		return true
	}
	return false
}

func findComboOperand(program, regA, regB, regC int) (int, error) {
	if program >= 0 && program <= 3 {
		return program, nil
	}

	switch {
	case program == 4:
		return regA, nil
	case program == 5:
		return regB, nil
	case program == 6:
		return regC, nil
	case program == 7:
		return -1, errors.New("Combo operand 7 is reserved and will not appear in valid programs.")
	}
	return -1, errors.New("Invalid program")

}
