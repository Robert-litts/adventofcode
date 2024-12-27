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
		isBlocked = matrix[loc.Y][loc.X] == "#"
	}

	return inBounds, isBlocked
}

func main() {
	var inputFile = flag.String("inputFile", "../input/day6.input", "Relative file path to use as input.")
	flag.Parse()
	start := time.Now()
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
	matrix, guardLoc, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	fmt.Println("Guard Starting at :", guardLoc)
	inside := true
	dirUp := true
	dirDown := false
	dirLeft := false
	dirRight := false

	var visited Visited
	visited.Points = append(visited.Points, guardLoc) // Add the starting point to visited

	moveUp := Coordinate{X: 0, Y: -1}
	moveDown := Coordinate{X: 0, Y: 1}
	moveLeft := Coordinate{X: -1, Y: 0}
	moveRight := Coordinate{X: 1, Y: 0}

	// Determine the number of rows and columns
	y := len(matrix)    // number of rows, "Height", "Y"
	x := len(matrix[0]) // number of columns (assuming all rows are the same length), "Width"

	// Print the matrix for debugging purposes (uncomment to see the map)
	// for i := 0; i < y; i++ {
	// 	fmt.Println(matrix[i])
	// }

	for {
		if inside {

			// Movement logic
			if dirUp {
				newLoc := guardLoc.Add(moveUp)

				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving up to: %d, %d\n", newLoc.X, newLoc.Y)
					//fmt.Println("Next Spot Contains : ", matrix[newLoc.X][newLoc.Y])
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirUp = false
					dirRight = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")

					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			} else if dirDown {
				newLoc := guardLoc.Add(moveDown)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving down to: %d, %d\n", newLoc.X, newLoc.Y)
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirDown = false
					dirLeft = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")
					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			} else if dirLeft {
				newLoc := guardLoc.Add(moveLeft)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving left to: %d, %d\n", newLoc.X, newLoc.Y)
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirLeft = false
					dirUp = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")
					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			} else if dirRight {
				newLoc := guardLoc.Add(moveRight)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving right to: %d, %d\n", newLoc.X, newLoc.Y)
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirRight = false
					dirDown = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")
					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			}

		} else {
			break
		}
	}

	score = len(visited.Points)
	fmt.Println(visited.Points)
	fmt.Printf("Final Guard Location: %d, %d\n", guardLoc.X, guardLoc.Y)
	fmt.Println("Score:", score)
	return nil
}

///////////////////////////////////////////////////////////////////////////////

func Part2(inputFile string) error {
	score := 0
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, guardLoc, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	start := guardLoc

	fmt.Println("Guard Starting at :", guardLoc)
	inside := true
	dirUp := true
	dirDown := false
	dirLeft := false
	dirRight := false

	var visited Visited
	visited.Points = append(visited.Points, guardLoc) // Add the starting point to visited

	moveUp := Coordinate{X: 0, Y: -1, Dir: "U"}
	moveDown := Coordinate{X: 0, Y: 1, Dir: "D"}
	moveLeft := Coordinate{X: -1, Y: 0, Dir: "L"}
	moveRight := Coordinate{X: 1, Y: 0, Dir: "R"}

	// Determine the number of rows and columns
	y := len(matrix)    // number of rows, "Height", "Y"
	x := len(matrix[0]) // number of columns (assuming all rows are the same length), "Width"

	// Print the matrix for debugging purposes (uncomment to see the map)
	// for i := 0; i < y; i++ {
	// 	fmt.Println(matrix[i])
	// }

	for {
		if inside {

			// Movement logic
			if dirUp {
				newLoc := guardLoc.Add(moveUp)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving up to: %d, %d\n", newLoc.X, newLoc.Y)
					//fmt.Println("Next Spot Contains : ", matrix[newLoc.X][newLoc.Y])
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirUp = false
					dirRight = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")

					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			} else if dirDown {
				newLoc := guardLoc.Add(moveDown)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving down to: %d, %d\n", newLoc.X, newLoc.Y)
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirDown = false
					dirLeft = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")
					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			} else if dirLeft {
				newLoc := guardLoc.Add(moveLeft)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving left to: %d, %d\n", newLoc.X, newLoc.Y)
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirLeft = false
					dirUp = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")
					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			} else if dirRight {
				newLoc := guardLoc.Add(moveRight)
				inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
				if inBounds && !isBlocked {
					//fmt.Printf("Moving right to: %d, %d\n", newLoc.X, newLoc.Y)
					guardLoc = newLoc
					visited.AddPoint(newLoc)
				} else if isBlocked {
					//fmt.Println("Blocked by an obstacle.")
					dirRight = false
					dirDown = true
				} else if !inBounds {
					fmt.Println("Move is out of bounds.")
					dirUp = false
					dirLeft = false
					dirRight = false
					dirLeft = false
					inside = false
					break
				}
			}

		} else {
			break
		}
	}

	//Loop through all the points the guard will visit
	for _, point := range visited.Points {
		// if point.X == newLoc.X && point.Y == newLoc.Y {
		// 	found = true
		// 	break
		// }

		var visitedLoop Visited

		//skip the starting location
		if point.X == start.X && point.Y == start.Y {
			continue
		}
		//Add obstruction to matrix
		matrix[point.Y][point.X] = "#"
		guardLoc = start //reset guard to the starting point
		//reset all logic
		inside := true
		dirUp := true
		dirDown := false
		dirLeft := false
		dirRight := false

		//Rerun the matrix with the new obstruction
		for {
			if inside {

				// Movement logic
				if dirUp {
					newLoc := guardLoc.Add(moveUp)

					inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
					if inBounds && !isBlocked {
						//fmt.Printf("Moving up to: %d, %d\n", newLoc.X, newLoc.Y)
						//fmt.Println("Next Spot Contains : ", matrix[newLoc.X][newLoc.Y])
						guardLoc = newLoc
						loop := visitedLoop.CheckDir(newLoc)
						if loop {
							//fmt.Println("Loop detected.")
							score++
							dirUp = false
							dirLeft = false
							dirRight = false
							dirLeft = false
							inside = false
							break
						}
					} else if isBlocked {
						//fmt.Println("Blocked by an obstacle.")
						dirUp = false
						dirRight = true
					} else if !inBounds {
						//fmt.Println("Move is out of bounds.")

						dirUp = false
						dirLeft = false
						dirRight = false
						dirLeft = false
						inside = false
						break
					}
				} else if dirDown {
					newLoc := guardLoc.Add(moveDown)
					inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
					if inBounds && !isBlocked {
						//fmt.Printf("Moving down to: %d, %d\n", newLoc.X, newLoc.Y)
						guardLoc = newLoc
						loop := visitedLoop.CheckDir(newLoc)
						if loop {
							//fmt.Println("Loop detected.")
							score++
							dirUp = false
							dirLeft = false
							dirRight = false
							dirLeft = false
							inside = false
							break
						}
					} else if isBlocked {
						//fmt.Println("Blocked by an obstacle.")
						dirDown = false
						dirLeft = true
					} else if !inBounds {
						//fmt.Println("Move is out of bounds.")
						dirUp = false
						dirLeft = false
						dirRight = false
						dirLeft = false
						inside = false
						break
					}
				} else if dirLeft {
					newLoc := guardLoc.Add(moveLeft)
					inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
					if inBounds && !isBlocked {
						//fmt.Printf("Moving left to: %d, %d\n", newLoc.X, newLoc.Y)
						guardLoc = newLoc
						loop := visitedLoop.CheckDir(newLoc)
						if loop {
							//fmt.Println("Loop detected.")
							score++
							dirUp = false
							dirLeft = false
							dirRight = false
							dirLeft = false
							inside = false
							break
						}
					} else if isBlocked {
						//fmt.Println("Blocked by an obstacle.")
						dirLeft = false
						dirUp = true
					} else if !inBounds {
						//fmt.Println("Move is out of bounds.")
						dirUp = false
						dirLeft = false
						dirRight = false
						dirLeft = false
						inside = false
						break
					}
				} else if dirRight {
					newLoc := guardLoc.Add(moveRight)
					inBounds, isBlocked := isValidMove(newLoc, matrix, x, y)
					if inBounds && !isBlocked {
						//fmt.Printf("Moving right to: %d, %d\n", newLoc.X, newLoc.Y)
						guardLoc = newLoc
						loop := visitedLoop.CheckDir(newLoc)
						if loop {
							//fmt.Println("Loop detected.")
							score++
							dirUp = false
							dirLeft = false
							dirRight = false
							dirLeft = false
							inside = false
							break
						}
					} else if isBlocked {
						//fmt.Println("Blocked by an obstacle.")
						dirRight = false
						dirDown = true
					} else if !inBounds {
						//fmt.Println("Move is out of bounds.")
						dirUp = false
						dirLeft = false
						dirRight = false
						dirLeft = false
						inside = false
						break
					}
				}

			} else {
				break
			}
		}
		//Reset point to dot
		//fmt.Println("Resetting Matrix")
		matrix[point.Y][point.X] = "."

	}

	//score = len(visited.Points)
	//mt.Println(visited.Points)
	fmt.Printf("Final Guard Location: %d, %d\n", guardLoc.X, guardLoc.Y)
	fmt.Println("Score P2: ", score)
	return nil
}

// makeMatrix takes the byte slice, splits it into lines and converts it into a matrix of strings.
func makeMatrix(bytes []byte) ([][]string, Coordinate, error) {
	var matrix [][]string
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var guard_loc Coordinate

	for k, line := range lines {
		numCols := len(line)

		row := make([]string, numCols)
		for i, c := range line {
			row[i] = string(c)
			if row[i] == "^" {
				guard_loc = Coordinate{X: i, Y: k} // Store the guard location
			} // Convert each rune to a string
		}
		matrix = append(matrix, row)
	}

	return matrix, guard_loc, nil
}
