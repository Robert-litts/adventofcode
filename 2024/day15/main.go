package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct to represent a visited list of coordinates. Points is a slice of Coordinate structs.
type Visited struct {
	Points []Coordinate
}

// Method to add a point to the visited list. If the point already exists, it does nothing.
func (v *Visited) AddPoint(newLoc Coordinate) {
	found := false
	for _, point := range v.Points {
		if point.X == newLoc.X && point.Y == newLoc.Y {
			found = true
			break
		}
	}
	if !found {
		v.Points = append(v.Points, newLoc)
	}
}

// Method to add a point to the visited list. If the point already exists, it does nothing.
func (v *Visited) CheckDir(newLoc Coordinate) bool {
	found := false
	for _, point := range v.Points {
		if point == newLoc {
			found = true
			break
		}
	}
	if !found {
		v.Points = append(v.Points, newLoc)
	}
	return found
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Coordinate struct {
	X   int
	Y   int
	Dir string
}

// Method to move the coordinate by a delta value. Delta is added to current coordinates.
func (c *Coordinate) Move(delta Coordinate) {
	c.X += delta.X
	c.Y += delta.Y
}

// Method to add two coordinates
func (c *Coordinate) Add(delta Coordinate) Coordinate {
	return Coordinate{X: c.X + delta.X, Y: c.Y + delta.Y, Dir: delta.Dir}
}

// Helper function to validate a move, returns a tuple of booleans indicating if the move is in bounds and if it's blocked.
func isValidMove(loc Coordinate, matrix [][]string, maxX, maxY int) (bool, bool) {
	// Check if the location is inbounds
	inBounds := loc.X >= 0 && loc.Y >= 0 && loc.X < maxX && loc.Y < maxY

	// Check if the position is blocked ("#")
	isBlocked := false
	if inBounds {
		isBlocked = matrix[loc.Y][loc.X] == "#" || matrix[loc.Y][loc.X] == "O"
	}

	return inBounds, isBlocked
}

func main() {
	var inputFile = flag.String("inputFile", "../input/day15.input", "Relative file path to use as input.")
	flag.Parse()
	start := time.Now()
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
	duration := time.Since(start)
	fmt.Printf("Execution Time: %s\n", duration)
}

func Part1(inputFile string) error {
	score := 0
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, moves, guardLoc, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	fmt.Println("Guard Starting at :", guardLoc)
	fmt.Println("Moves: ", moves)
	// inside := true
	// dirUp := true
	// dirDown := false
	// dirLeft := false
	// dirRight := false

	var visited Visited
	newLoc := guardLoc
	visited.Points = append(visited.Points, guardLoc) // Add the starting point to visited

	moveUp := Coordinate{X: 0, Y: -1}
	moveDown := Coordinate{X: 0, Y: 1}
	moveLeft := Coordinate{X: -1, Y: 0}
	moveRight := Coordinate{X: 1, Y: 0}
	movesMap := make(map[string]Coordinate)
	movesMap["^"] = moveUp
	movesMap["v"] = moveDown
	movesMap["<"] = moveLeft
	movesMap[">"] = moveRight

	// Determine the number of rows and columns
	y := len(matrix) // number of rows, "Height", "Y"
	//x := len(matrix[0]) // number of columns (assuming all rows are the same length), "Width"

	//Print the matrix for debugging purposes (uncomment to see the map)
	for i := 0; i < y; i++ {
		fmt.Println(matrix[i])
	}
	foundCrate := false
	//newMove := false
	for i := 0; i < len(moves); i++ {
		fmt.Println("Current Move: ", moves[i])

		if moves[i] == "^" {
			newLoc = newLoc.Add(moveUp)
			visited.Points = append(visited.Points, newLoc)
		} else if moves[i] == "v" {
			newLoc = newLoc.Add(moveDown)
			visited.Points = append(visited.Points, newLoc)
		} else if moves[i] == "<" {
			newLoc = newLoc.Add(moveLeft)
			visited.Points = append(visited.Points, newLoc)
		} else if moves[i] == ">" {
			newLoc = newLoc.Add(moveRight)
			visited.Points = append(visited.Points, newLoc)
		}
		fmt.Println("New location: ", newLoc)
		for {
			// Check if the location is inbounds and not blocked
			inBounds, isBlocked := isValidMove(newLoc, matrix, y, len(matrix[0]))

			//First Case: Free Space -- move guard only
			if inBounds && !isBlocked {
				fmt.Printf("Moving %s to: %d, %d\n", moves[i], newLoc.X, newLoc.Y)
				matrix[guardLoc.Y][guardLoc.X] = "."        //Make Guard's previous position a "."
				guardLoc = guardLoc.Add(movesMap[moves[i]]) //move guard to next position

				matrix[guardLoc.Y][guardLoc.X] = "@" // Move guard on the matrix
				//Move crate to newLoc
				if foundCrate {
					matrix[newLoc.Y][newLoc.X] = "O" //Move crate to new open position
					foundCrate = false

				}
				newLoc = guardLoc //reset to guard's location

				visited.Points = append(visited.Points, newLoc)
				//newMove = true //get new move

				break

				//if blocked, could be either # (wall) or O Crate
			}

			//Second Case: Wall, get new move
			if inBounds && isBlocked && matrix[newLoc.Y][newLoc.X] == "#" {
				fmt.Println("Blocked by wall, getting new move")
				newLoc = guardLoc //Reached a wall, reset newLoc back to GuardLoc
				break
			}

			//Third case - Crate, check next position in same direction
			if inBounds && isBlocked && matrix[newLoc.Y][newLoc.X] == "O" { //reached a crate, need to check if we can move the crate
				fmt.Printf("Space blocked from moving %s from %d, %d to: %d, %d, there is a %s in the way\n",
					moves[i], guardLoc.X, guardLoc.Y, newLoc.X, newLoc.Y, matrix[newLoc.Y][newLoc.X]) // Check if we can move the crate to the new location
				// If yes, move the crate and update the guard's position
				// If no, get new move and update guardLoc
				newLoc = newLoc.Add(movesMap[moves[i]]) // Check the next spot in the same direction as the guard's movement
				fmt.Println("BLOCKED by CRATE, Checking new Location: ", newLoc)
				foundCrate = true

				continue
			}

		}
		//For debugging, print every updated matrix
		// y = len(matrix)
		// for j := 0; j < y; j++ {
		// 	fmt.Println(matrix[j])
		// }
	}

	gps := 0
	for rowIndex, row := range matrix {
		for colIndex, cell := range row {
			for _, char := range cell {
				if string(char) == "O" {
					gps += rowIndex*100 + colIndex //Calculate GPS position
				}
			}
		}
	}
	fmt.Printf("Final Guard Location: %d, %d\n", guardLoc.X, guardLoc.Y)
	fmt.Println("Score:", score)
	fmt.Println("GPS: ", gps)
	return nil
}

// makeMatrix takes the byte slice, splits it into lines and converts it into a matrix of strings.
func makeMatrix(bytes []byte) ([][]string, []string, Coordinate, error) {
	var matrix [][]string
	movesList := []string{}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var guard_loc Coordinate
	//var moves bool = false
	movesLine := 0

	for k, line := range lines {
		numCols := len(line)
		if line == "" {
			//moves = true
			movesLine = k + 1
			break
		}
		row := make([]string, numCols)
		for i, c := range line {
			row[i] = string(c)

			if row[i] == "@" {
				guard_loc = Coordinate{X: i, Y: k} // Store the guard location
			} // Convert each rune to a string
		}
		matrix = append(matrix, row)
	}

	for i := movesLine; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}

		for _, m := range line {
			movesList = append(movesList, string(m))
		}
	}

	return matrix, movesList, guard_loc, nil
}
