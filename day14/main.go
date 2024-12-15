package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day14.input", "Relative file path to use as input.")
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
	width := 101
	midWidth := math.Ceil(float64(width)/2) - 1

	height := 103
	midHeight := math.Ceil(float64(height)/2) - 1

	//for _, line := range strings.Split(string(bytes), "\n") {

	contents := string(bytes)
	safety := 0
	//lowest := 0
	quadTL := 0
	quadTR := 0
	quadBL := 0
	quadBR := 0
	minSafety := math.MaxFloat64 // Start with a very high value to ensure any safety value will be smaller
	//lowestIndex := -1
	nonSafe := 0

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

	//fmt.Println("Initial Matrix")
	//matrix := makeMatrix(width, height, &robots)
	// for _, char := range matrix {
	// 	fmt.Println(char)
	// }
	for i := 0; i < 500; i++ {
		nonSafe = 0

		//safety := 0
		for bots := range len(robots.Robots) {
			newX := robots.Robots[bots].X + robots.Robots[bots].vX
			newY := robots.Robots[bots].Y + robots.Robots[bots].vY

			//fmt.Printf("new X: %d, new Y: %d, height: %d, width: %d \n", newX, newY, height, width)

			// Wrap the X position within the width
			newX = (newX + width) % width

			// Wrap the Y position within the height
			newY = (newY + height) % height

			// Update the robot's coordinates
			robots.Robots[bots].X = newX
			robots.Robots[bots].Y = newY
			//fmt.Printf("X: %d, Y: %d, Midwidth: %f, Midheight: %f \n", robots.Robots[bots].X, robots.Robots[bots].Y, midWidth, midHeight)
			if robots.Robots[bots].X >= 0 && robots.Robots[bots].Y >= 0 {
				if robots.Robots[bots].X < int(midWidth) && robots.Robots[bots].Y < int(midHeight) {
					quadTL++
					//fmt.Printf("TL Added, Robot %d: X=%d, Y=%d\n", bots, robots.Robots[bots].X, robots.Robots[bots].Y)

				}
				if robots.Robots[bots].X > int(midWidth) && robots.Robots[bots].Y < int(midHeight) {
					quadTR++
					//fmt.Printf("TR Added, Robot %d: X=%d, Y=%d\n", bots, robots.Robots[bots].X, robots.Robots[bots].Y)

				}
				if robots.Robots[bots].X > int(midWidth) && robots.Robots[bots].Y > int(midHeight) {
					quadBR++
					//fmt.Printf("BR, Robot %d: X=%d, Y=%d\n", bots, robots.Robots[bots].X, robots.Robots[bots].Y)

				}
				if robots.Robots[bots].X < int(midWidth) && robots.Robots[bots].Y > int(midHeight) {
					quadBL++
					//fmt.Printf("BL, Robot %d: X=%d, Y=%d\n", bots, robots.Robots[bots].X, robots.Robots[bots].Y)

				}
				if robots.Robots[bots].X == int(midWidth) || robots.Robots[bots].Y == int(midHeight) {
					nonSafe++
				}

			}

		}
		fmt.Printf("Matrix after %d seconds \n", i)
		//matrix = makeMatrix(width, height, &robots)
		// for _, char := range matrix {
		// 	fmt.Println(char)

		// }
		safety = quadTL * quadTR * quadBL * quadBR
		// if float64(safety) < minSafety {
		// 	minSafety = float64(safety)
		// 	//lowestIndex = i // Update the index where the lowest safety occurs
		// 	fmt.Println("Safety decreased at second:", i)
		// 	fmt.Println("Safety:", safety)
		// 	fmt.Println("Lowest:", minSafety)
		// }

		// fmt.Println("Quadrant TL: ", quadTL)
		// fmt.Println("Quadrant TR: ", quadTR)
		// fmt.Println("Quadrant BL: ", quadBL)
		// fmt.Println("Quadrant BR: ", quadBR)
		fmt.Println("Non Safe: ", nonSafe)
		if nonSafe > 20 {
			matrix := makeMatrix(width, height, &robots)
			for _, char := range matrix {
				fmt.Println(char)
			}

		}

		if float64(safety) < minSafety {
			minSafety = float64(safety)

			fmt.Println("Safety: ", safety)
			fmt.Println("Lowest:", minSafety)

			quadBL = 0
			quadBR = 0
			quadTL = 0
			quadTR = 0

		}
	}
	// fmt.Println("Quadrant TL: ", quadTL)
	// fmt.Println("Quadrant TR: ", quadTR)
	// fmt.Println("Quadrant BL: ", quadBL)
	// fmt.Println("Quadrant BR: ", quadBR)
	// safety = quadTL * quadTR * quadBL * quadBR

	fmt.Println("Safety: ", safety)
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
