package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day4.example", "Relative file path to use as input.")
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
	score := 0
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	// Determine the number of rows and columns
	rows := len(matrix)       // number of rows
	columns := len(matrix[0]) // number of columns (assuming all rows are the same length)
	fmt.Println("Rows: ", rows)
	fmt.Println("Columns: ", columns)

	// Parse the matrix to find X and check surrounding cells for "MAS" in all directions
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == "X" {
				// Check surrounding cells for "MAS" in all directions

				// Upwards (i-1, i-2, i-3)
				if i-3 >= 0 && matrix[i-1][j] == "M" && matrix[i-2][j] == "A" && matrix[i-3][j] == "S" {
					score++
				}
				// Downwards (i+1, i+2, i+3)
				if i+3 < rows && matrix[i+1][j] == "M" && matrix[i+2][j] == "A" && matrix[i+3][j] == "S" {
					score++
				}
				// Left (j-1, j-2, j-3)
				if j-3 >= 0 && matrix[i][j-1] == "M" && matrix[i][j-2] == "A" && matrix[i][j-3] == "S" {
					score++
				}
				// Right (j+1, j+2, j+3)
				if j+3 < columns && matrix[i][j+1] == "M" && matrix[i][j+2] == "A" && matrix[i][j+3] == "S" {
					score++
				}
				// Diagonal: Top-left to bottom-right (i-1, j-1) -> (i-2, j-2) -> (i-3, j-3)
				if i-3 >= 0 && j-3 >= 0 && matrix[i-1][j-1] == "M" && matrix[i-2][j-2] == "A" && matrix[i-3][j-3] == "S" {
					score++
				}
				// Diagonal: Bottom-right to top-left (i+1, j+1) -> (i+2, j+2) -> (i+3, j+3)
				if i+3 < rows && j+3 < columns && matrix[i+1][j+1] == "M" && matrix[i+2][j+2] == "A" && matrix[i+3][j+3] == "S" {
					score++
				}
				// Diagonal: Top-right to bottom-left (i-1, j+1) -> (i-2, j+2) -> (i-3, j+3)
				if i-3 >= 0 && j+3 < columns && matrix[i-1][j+1] == "M" && matrix[i-2][j+2] == "A" && matrix[i-3][j+3] == "S" {
					score++
				}
				// Diagonal: Bottom-left to top-right (i+1, j-1) -> (i+2, j-2) -> (i+3, j-3)
				if i+3 < rows && j-3 >= 0 && matrix[i+1][j-1] == "M" && matrix[i+2][j-2] == "A" && matrix[i+3][j-3] == "S" {
					score++
				}
			}
		}
	}

	fmt.Println("Score Part 1: ", score)
	return nil
}

// makeMatrix takes the byte slice, splits it into lines and converts it into a matrix of strings.
func makeMatrix(bytes []byte) ([][]string, error) {
	var matrix [][]string
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		numCols := len(line)

		row := make([]string, numCols)
		for i, c := range line {
			row[i] = string(c) // Convert each rune to a string
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}
