package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day13.input", "Relative file path to use as input.")
	flag.Parse()
	// fmt.Println("Running Part 1:")
	// if err := Part1(*inputFile); err != nil {
	// 	fmt.Println("Error in Part 2:", err)
	// 	return
	// }
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
	var result float64
	contents := string(bytes)
	var aX, aY, bX, bY, X, Y float64

	lines := strings.Split(contents, "\n")
	count := 0

	// Iterate through lines in blocks of 3
	for i := 0; i < len(lines); i++ {
		parse := false
		line := strings.TrimSpace(lines[i]) // Remove spaces

		// Skip empty
		if line == "" {
			continue
		}

		switch count % 3 {
		case 0: // Button A line
			if n, err := fmt.Sscanf(line, "Button A: X+%f, Y+%f", &aX, &bX); err == nil && n == 2 {

			}
		case 1: // Button B line
			if n, err := fmt.Sscanf(line, "Button B: X+%f, Y+%f", &aY, &bY); err == nil && n == 2 {

			}
		case 2: // Prize line
			if n, err := fmt.Sscanf(line, "Prize: X=%f, Y=%f", &X, &Y); err == nil && n == 2 {

				parse = true

			}
		}
		// Increment count after processing
		count++

		if parse {
			//Solve after processing the A/B/Prize Lines
			fmt.Println("solving system of equations for :", aX, bX, aY, bY, X, Y)
			A := (X*bY - Y*aY) / (aX*bY - aY*bX) //Cramer's Rule, thanks Kristy! https://www.cuemath.com/algebra/cramers-rule/
			B := (Y*aX - X*bX) / (aX*bY - aY*bX)
			fmt.Println("A: ", A)
			fmt.Println("B: ", B)
			if A >= 0 && A <= 100 && B >= 0 && B <= 100 && A == float64(int64(A)) && B == float64(int64(B)) {
				result += A*3 + B
			}

		}
	}

	// Display result (if needed)
	fmt.Println("Result:", result)
	return nil
}

func Part2(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	var result float64
	contents := string(bytes)
	var aX, aY, bX, bY, X, Y float64
	offset := float64(10000000000000) //Part 2 offset, added to X & Y in the system of equations

	lines := strings.Split(contents, "\n")
	count := 0

	// Iterate through lines in blocks of 3
	for i := 0; i < len(lines); i++ {
		parse := false
		line := strings.TrimSpace(lines[i]) // Remove spaces

		// Skip empty
		if line == "" {
			continue
		}

		switch count % 3 {
		case 0: // Button A line
			if n, err := fmt.Sscanf(line, "Button A: X+%f, Y+%f", &aX, &bX); err == nil && n == 2 {

			}
		case 1: // Button B line
			if n, err := fmt.Sscanf(line, "Button B: X+%f, Y+%f", &aY, &bY); err == nil && n == 2 {

			}
		case 2: // Prize line
			if n, err := fmt.Sscanf(line, "Prize: X=%f, Y=%f", &X, &Y); err == nil && n == 2 {
				X += offset //Part 2 offset, added to X & Y in the system of equations
				Y += offset

				parse = true

			}
		}
		// Increment count after processing
		count++

		if parse {
			//Solve after processing the A/B/Prize Lines
			fmt.Println("solving system of equations for :", aX, bX, aY, bY, X, Y)
			A := (X*bY - Y*aY) / (aX*bY - aY*bX) //Cramer's Rule, thanks Kristy! https://www.cuemath.com/algebra/cramers-rule/
			B := (Y*aX - X*bX) / (aX*bY - aY*bX)
			fmt.Println("A: ", A)
			fmt.Println("B: ", B)
			//Part 2 change, remove the limit to 100 presses per A/B
			if A >= 0 && B >= 0 && A == float64(int64(A)) && B == float64(int64(B)) {
				result += A*3 + B
			}

		}
	}

	// Display result (if needed)
	fmt.Println("Result:", result)
	return nil
}
