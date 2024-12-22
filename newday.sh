#!/bin/bash
#This script generates a new Advent of Code folder and pulls the input for the day
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

#Set the day variable from the command line argument
day=$1
dir="day$day"

# Check if the directory exists & create folder
if [ -d "$dir" ]; then
    echo "Directory $dir already exists."
    exit 1
else
    echo "Creating new dir for day $day"
    mkdir -p "$dir"
fi

#Navigate to the new folder
cd "day$day" || exit

#Create the new Go file
touch "main.go"

#Write starter go file contents
cat > main.go << EOF

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day17.example", "Relative file path to use as input.")
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
	//scoreP1 := 0
	//pq := &Queue{}
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, start, end, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println("Start: ", start)
	fmt.Println("End: ", end)

	return nil
}

func makeMatrix(bytes []byte) ([][]string, Coordinate, Coordinate, error) {
	var matrix [][]string
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var start, end Coordinate
	//var nodeCount int

	for rowIndex, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Row to hold the integers
		row := make([]string, len(line))

		for colIndex, c := range line {

			// Assign the integer to the row
			row[colIndex] = string(c)

			// Check if the number is 0 and record the position
			if string(c) == "S" {
				start = Coordinate{X: rowIndex, Y: colIndex}
				fmt.Println("Starting Position: ", start)

			}

			if string(c) == "E" {
				end = Coordinate{X: rowIndex, Y: colIndex}
				fmt.Println("Ending Position: ", end)
			}

		}

		// Append the row to the matrix
		matrix = append(matrix, row)
	}

	// Return the created matrix and the list of zero positions
	return matrix, start, end, nil
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Coordinate struct {
	X int
	Y int
}
EOF

#Pull the input for the day
cd .. && ./input/input --day 17

#Create blank example file
touch ./input/day$day.example

echo "Created folder and input for day $day"