package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day8.example", "Relative file path to use as input.")
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

	// makeMatrix function will create the matrix from the input
	matrix, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	// Determine the number of rows and columns
	rows := len(matrix)       // number of rows
	columns := len(matrix[0]) // number of columns (assuming all rows are the same length)
	antennaMap := AntennaMap{}
	antennaType := []string{}
	antinodeMap := AntinodeMap{}

	// Parse the matrix to find X and check surrounding cells for "MAS" in all directions
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] != "." {
				found := false
				for _, v := range antennaType {
					if v == string(matrix[i][j]) {
						found = true
						break
					}
				}
				//Add unique types to the antennaType list
				if !found {
					antennaType = append(antennaType, string(matrix[i][j]))
				}

				//Map all the antenna locations
				//fmt.Printf("Antenna Found: %s at %d, %d \n", matrix[i][j], i, j)
				antenna := Antenna{
					X:    i,
					Y:    j,
					Type: string(matrix[i][j]),
				}
				antennaMap.Antennas = append(antennaMap.Antennas, antenna)
			}
		}
	}

	//fmt.Println("Antenna Map: ", antennaMap)

	//Loop through all antennas by type
	for i := 0; i < len(antennaType); i++ {
		//fmt.Println("Antenna Type: ", antennaType[i])
		for j := 0; j < len(antennaMap.Antennas); j++ {
			if antennaMap.Antennas[j].Type == antennaType[i] {
				//fmt.Println("AntennaFound: ", antennaMap.Antennas[j])
				//Find distance between antenna and all other antennas of the same type
				for k := 0; k < len(antennaMap.Antennas); k++ {
					if antennaMap.Antennas[k].Type == antennaType[i] && k != j {
						//fmt.Println("Finding antinodes between: ", antennaMap.Antennas[j].X, " ", antennaMap.Antennas[j].Y, " and ", antennaMap.Antennas[k].X, antennaMap.Antennas[k].Y)
						//fmt.Println("Finding antinodes between: ", antennaMap.Antennas[j], " and ", antennaMap.Antennas[k])
						x3, y3, x4, y4 := findAntiNodes(antennaMap.Antennas[j].X, antennaMap.Antennas[k].X, antennaMap.Antennas[j].Y, antennaMap.Antennas[k].Y)
						inBoundsX := x3 >= 0 && x3 < columns && y3 >= 0 && y3 < rows
						inBoundsY := x4 >= 0 && x4 < columns && y4 >= 0 && y4 < rows

						//Calculate distance
						if inBoundsX {
							//fmt.Printf("New Antinodes, (%d, %d) \n", x3, y3)
							antinode := Antinode{
								X: x3,
								Y: y3,
							}
							antinodeMap.AddPoint(antinode)

						}
						if inBoundsY {
							//fmt.Printf("New Antinodes, (%d, %d) \n", x4, y4)
							antinode := Antinode{
								X: x4,
								Y: y4,
							}

							antinodeMap.AddPoint(antinode)
						}

					}
				}

			}
		}
	}
	//fmt.Println("Antinodes: ", antinodeMap.Antinodes)
	fmt.Println("Antinodes Total: ", len(antinodeMap.Antinodes))
	return nil
}

func Part2(inputFile string) error {
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
	antennaMap := AntennaMap{}
	antennaType := []string{}
	antinodeMap := AntinodeMap{}

	// Parse the matrix to find X and check surrounding cells for "MAS" in all directions
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] != "." {
				found := false
				for _, v := range antennaType {
					if v == string(matrix[i][j]) {
						found = true
						break
					}
				}
				//Add unique types to the antennaType list
				if !found {
					antennaType = append(antennaType, string(matrix[i][j]))
				}

				//Map all the antenna locations
				//fmt.Printf("Antenna Found: %s at %d, %d \n", matrix[i][j], i, j)
				antenna := Antenna{
					X:    i,
					Y:    j,
					Type: string(matrix[i][j]),
				}
				antennaMap.Antennas = append(antennaMap.Antennas, antenna)
			}
		}
	}

	//fmt.Println("Antenna Map: ", antennaMap)

	//Loop through all antennas by type
	for i := 0; i < len(antennaType); i++ {
		//fmt.Println("Antenna Type: ", antennaType[i])
		for j := 0; j < len(antennaMap.Antennas); j++ {
			if antennaMap.Antennas[j].Type == antennaType[i] {
				//fmt.Println("AntennaFound: ", antennaMap.Antennas[j])
				//Find distance between antenna and all other antennas of the same type
				for k := 0; k < len(antennaMap.Antennas); k++ {
					if antennaMap.Antennas[k].Type == antennaType[i] && k != j {
						//fmt.Println("Finding antinodes between: ", antennaMap.Antennas[j].X, " ", antennaMap.Antennas[j].Y, " and ", antennaMap.Antennas[k].X, antennaMap.Antennas[k].Y)
						findAntiNodesPart2(antennaMap.Antennas[j].X, antennaMap.Antennas[k].X, antennaMap.Antennas[j].Y, antennaMap.Antennas[k].Y, columns, rows, &antinodeMap)

					}

				}
			}
		}
	}
	//fmt.Println("Antinodes: ", antinodeMap.Antinodes)
	fmt.Println("Antinodes Total: ", len(antinodeMap.Antinodes))
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

// Struct to represent an antenna on the grid. X and Y are integers.
type Antenna struct {
	X    int
	Y    int
	Type string
}

// Struct to represent an antinode.
type Antinode struct {
	X int
	Y int
}

// Struct to represent a visited list of antennas.
type AntennaMap struct {
	Antennas []Antenna
}

// Struct to represent a visited list of antennas.
type AntinodeMap struct {
	Antinodes []Antinode
}

// Method to add a point to the antenna map. If the point already exists, it does nothing.
func (a *AntennaMap) AddPoint(Loc Antenna) {
	found := false
	for _, point := range a.Antennas {
		if point == Loc {
			found = true
			break
		}
	}
	if !found {
		a.Antennas = append(a.Antennas, Loc)
	}
}

func (a *AntinodeMap) AddPoint(Antinode Antinode) {
	found := false
	for _, point := range a.Antinodes {
		if point == Antinode {
			found = true
			break
		}
	}
	if !found {
		a.Antinodes = append(a.Antinodes, Antinode)
	}
}

func findAntiNodes(x1, x2, y1, y2 int) (x3, y3, x4, y4 int) {
	x3 = 2*x1 - x2
	y3 = 2*y1 - y2
	x4 = 2*x2 - x1
	y4 = 2*y2 - y1

	return x3, y3, x4, y4
}

func findAntiNodesPart2(x1, x2, y1, y2, columns, rows int, antinodeMap *AntinodeMap) {
	dx := x2 - x1
	dy := y2 - y1
	x3 := 2*x1 - x2
	y3 := 2*y1 - y2
	x4 := 2*x2 - x1
	y4 := 2*y2 - y1

	// Add the first antinode at (x3, y3)
	if x3 >= 0 && x3 < columns && y3 >= 0 && y3 < rows {
		antinode := Antinode{X: x3, Y: y3}
		antinodeMap.AddPoint(antinode)
	}

	// Add the second antinode at (x4, y4)
	if x4 >= 0 && x4 < columns && y4 >= 0 && y4 < rows {
		antinode := Antinode{X: x4, Y: y4}
		antinodeMap.AddPoint(antinode)
	}

	// Check for repeating patterns starting from (x3, y3)
	for {
		x3 += dx
		y3 += dy
		if x3 < 0 || x3 >= columns || y3 < 0 || y3 >= rows {
			break
		}
		antinodeMap.AddPoint(Antinode{X: x3, Y: y3})
	}

	// Check for repeating patterns starting from (x4, y4)
	for {
		x4 -= dx
		y4 -= dy
		if x4 < 0 || x4 >= columns || y4 < 0 || y4 >= rows {
			break
		}
		antinodeMap.AddPoint(Antinode{X: x4, Y: y4})
	}

}
