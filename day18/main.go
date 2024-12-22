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
	var inputFile = flag.String("inputFile", "../input/day18.example", "Relative file path to use as input.")
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
	// makeMatrix function will create the matrix from the input
	matrix, err := makeMatrix(7)
	var corruption []Coordinate
	if err != nil {
		return err
	}
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	contents := string(bytes)
	var x, y int

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		position := strings.Split(line, ",") //Split at "|" to create the rules
		x, err = strconv.Atoi(strings.TrimSpace(position[0]))
		if err != nil {
			fmt.Println("Error converting first column value:", err)
			continue
		}
		y, err = strconv.Atoi(strings.TrimSpace(position[1]))
		if err != nil {
			fmt.Println("Error converting first column value:", err)
			continue
		}

		corruption = append(corruption, Coordinate{X: x, Y: y})
	}

	for k := 0; k < 12; k++ {
		coord := corruption[k] //Get the next coordinate from the list of corruptions

		matrix[coord.Y][coord.X] = "#" //Mark the corruption on the grid
	}

	//fmt.Println("X, Y: ", x, y)

	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println(corruption)

	// fmt.Println("Start: ", start)
	// fmt.Println("End: ", end)

	return nil
}

func makeMatrix(size int) ([][]string, error) {
	var matrix [][]string

	for j := 0; j < size; j++ {
		// Row to hold the integers
		row := make([]string, size)

		for i := 0; i < size; i++ {
			row[i] = "."
		}

		// Append the row to the matrix
		matrix = append(matrix, row)
	}

	// Return the created matrix and the list of zero positions
	return matrix, nil
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Coordinate struct {
	X int
	Y int
}
