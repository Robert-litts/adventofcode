package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day14.example", "Relative file path to use as input.")
	flag.Parse()
	fmt.Println("Running Part 1:")
	if err := Part1(*inputFile); err != nil {
		fmt.Println("Error in Part 2:", err)
		return
	}
	// fmt.Println("Running Part 2:")
	// if err := Part2(*inputFile); err != nil {
	// 	fmt.Println("Error in Part 2:", err)
	// 	return
	// }
}

func Part1(inputFile string) error {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	//result := 0
	var vX, vY, x, y int
	robots := Robots{}
	width := 11
	height := 7

	//for _, line := range strings.Split(string(bytes), "\n") {

	contents := string(bytes)

	lines := strings.Split(contents, "\n")

	// Iterate through lines in blocks of 3
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i]) // Remove spaces

		// Skip empty
		if line == "" {
			continue
		}

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vX, &vY)
		robot := Robot{X: x, Y: y, vX: vX, vY: vY}
		robots.Robots = append(robots.Robots, robot)

		// Display result (if needed)
		//fmt.Println("Result:", result)
	}
	// for bots := range len(robots.Robots) {
	// 	fmt.Println(robots.Robots[bots].vX, robots.Robots[bots].vY)
	// }

	fmt.Println("Initial Matrix")
	matrix := makeMatrix(width, height, &robots)
	for _, char := range matrix {
		fmt.Println(char)
	}
	for i := 0; i < 6; i++ {
		for bots := range len(robots.Robots) {
			fmt.Printf("Moving robot: %d X, %d Y, %d vX, %d vY \n", robots.Robots[bots].X, robots.Robots[bots].Y, robots.Robots[bots].vX, robots.Robots[bots].vY)
			newX := robots.Robots[bots].X + robots.Robots[bots].vX
			newY := robots.Robots[bots].Y + robots.Robots[bots].vY
			if newX >= width {
				newX = width - newX
			}
			if newX < 0 {
				newX = newX + width
			}

			if newY >= height {
				newY = height - newY
			}
			if newY < 0 {
				newY = newY + height
			}
			robots.Robots[bots].X = newX
			robots.Robots[bots].Y = newY

		}
		fmt.Printf("Matrix after %d seconds \n", i)
		matrix = makeMatrix(width, height, &robots)
		for _, char := range matrix {
			fmt.Println(char)

		}

	}

	return nil

}

// func Part2(inputFile string) error {
// 	bytes, err := os.ReadFile(inputFile)
// 	if err != nil {
// 		return err
// 	}
// 	var result float64
// 	contents := string(bytes)
// 	var aX, aY, bX, bY, X, Y float64
// 	offset := float64(10000000000000) //Part 2 offset, added to X & Y in the system of equations

// 	lines := strings.Split(contents, "\n")
// 	count := 0

// 	// Iterate through lines in blocks of 3
// 	for i := 0; i < len(lines); i++ {
// 		parse := false
// 		line := strings.TrimSpace(lines[i]) // Remove spaces

// 		// Skip empty
// 		if line == "" {
// 			continue
// 		}

// 		switch count % 3 {
// 		case 0: // Button A line
// 			if n, err := fmt.Sscanf(line, "Button A: X+%f, Y+%f", &aX, &bX); err == nil && n == 2 {

// 			}
// 		case 1: // Button B line
// 			if n, err := fmt.Sscanf(line, "Button B: X+%f, Y+%f", &aY, &bY); err == nil && n == 2 {

// 			}
// 		case 2: // Prize line
// 			if n, err := fmt.Sscanf(line, "Prize: X=%f, Y=%f", &X, &Y); err == nil && n == 2 {
// 				X += offset //Part 2 offset, added to X & Y in the system of equations
// 				Y += offset

// 				parse = true

// 			}
// 		}
// 		// Increment count after processing
// 		count++

// 		if parse {
// 			//Solve after processing the A/B/Prize Lines
// 			fmt.Println("solving system of equations for :", aX, bX, aY, bY, X, Y)
// 			A := (X*bY - Y*aY) / (aX*bY - aY*bX) //Cramer's Rule, thanks Kristy! https://www.cuemath.com/algebra/cramers-rule/
// 			B := (Y*aX - X*bX) / (aX*bY - aY*bX)
// 			fmt.Println("A: ", A)
// 			fmt.Println("B: ", B)
// 			//Part 2 change, remove the limit to 100 presses per A/B
// 			if A >= 0 && B >= 0 && A == float64(int64(A)) && B == float64(int64(B)) {
// 				result += A*3 + B
// 			}

// 		}
// 	}

// 	// Display result (if needed)
// 	fmt.Println("Result:", result)
// 	return nil
// }

func makeMatrix(width, height int, robots *Robots) [][]string {
	var matrix [][]string
	//uniqueChar := make(map[string][]Coordinate)

	//contents := string(bytes)

	for i := 0; i < height; i++ {

		// Row to hold the strings
		row := make([]string, width)

		for i := 0; i < width; i++ {

			row[i] = "."

		}
		matrix = append(matrix, row)
	}

	for bots := range len(robots.Robots) {
		x := robots.Robots[bots].X
		y := robots.Robots[bots].Y

		if matrix[y][x] != "." {
			val, err := strconv.Atoi(matrix[y][x])
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
			}
			val++
			matrix[y][x] = strconv.Itoa(val)
			continue
		}

		if matrix[y][x] == "." {
			matrix[y][x] = "1"
			continue
		}
	}

	return matrix
}

// Struct to represent a visited list of coordinates. Points is a slice of Coordinate structs.
type Robots struct {
	Robots []Robot
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Robot struct {
	X  int
	Y  int
	vX int
	vY int
}
