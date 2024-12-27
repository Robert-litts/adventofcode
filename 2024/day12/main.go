package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day12.example", "Relative file path to use as input.")
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

	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, unique, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	//fmt.Println(matrix)
	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("Unique: \n", unique)

	fmt.Println("\nUnique Characters and their Coordinates:")
	for char, coords := range unique {
		fmt.Printf("Character '%s': %v\n", char, coords)
		fmt.Println(coords[0].X, coords[0].Y)
	}

	visited := make(map[Coordinate]bool)
	price := 0

	// //Conduct breadth first search from each of the unique character on the map
	for char, coords := range unique {

		fmt.Printf("Processing character: %s\n", char)
		//For each character, search all coordinates where it appears
		for _, coord := range coords {

			//if we have already visited the coordinate previously (via another search), skip it.
			if visited[coord] {
				continue
			}
			//If we have never been to this coordinate, conduct a BFS and calculate the price

			price += BFS(matrix, coord, visited)
		}

		fmt.Printf("Visited %d nodes this run \n", len(visited))
		fmt.Printf("Area: %d\n", price)

	}

	fmt.Println("Final Area: ", price)
	return nil
}

func Part2(inputFile string) error {

	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, unique, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	//fmt.Println(matrix)
	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("Unique: \n", unique)

	fmt.Println("\nUnique Characters and their Coordinates:")
	for char, coords := range unique {
		fmt.Printf("Character '%s': %v\n", char, coords)
		fmt.Println(coords[0].X, coords[0].Y)
	}

	visited := make(map[Coordinate]bool)
	price := 0
	corners := 0

	// //Conduct breadth first search from each of the unique character on the map
	for char, coords := range unique {

		fmt.Printf("Processing character: %s\n", char)
		//For each character, search all coordinates where it appears
		for _, coord := range coords {

			//if we have already visited the coordinate previously (via another search), skip it.
			if visited[coord] {
				continue
			}
			//If we have never been to this coordinate, conduct a BFS and calculate the price

			price_new, corners_new := BFS2(matrix, coord, visited)
			price += price_new
			corners += corners_new
		}

		fmt.Printf("Visited %d nodes this run \n", len(visited))
		fmt.Printf("Area: %d\n", price)

	}

	fmt.Println("Final Area: ", price)
	fmt.Println("Corners: ", corners)
	return nil
}

///////////////////////////////////////////////////////////////////////////////

// makeMatrix takes the byte slice, splits it into lines and converts it into a 2D integer matrix.
func makeMatrix(bytes []byte) ([][]string, map[string][]Coordinate, error) {
	var matrix [][]string
	uniqueChar := make(map[string][]Coordinate)

	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	for rowIndex, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Row to hold the strings
		row := make([]string, len(line))

		for colIndex, c := range line {
			char := string(c)
			row[colIndex] = char

			coord := Coordinate{X: colIndex, Y: rowIndex}
			uniqueChar[char] = append(uniqueChar[char], coord)

		}
		matrix = append(matrix, row)
	}

	return matrix, uniqueChar, nil
}

type Summit struct {
	X, Y int
}

// Struct to represent a visited list of coordinates. Points is a slice of Coordinate structs.
type Visited struct {
	Points []Coordinate
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Coordinate struct {
	X int
	Y int
}

// Method to move the coordinate by a delta value. Delta is added to current coordinates.
func (c *Coordinate) Move(delta Coordinate) {
	c.X += delta.X
	c.Y += delta.Y
}

// Define the Queue struct to hold Coordinates
type Queue struct {
	List []Coordinate
}

// Function to add a coordinate to the queue
func (q *Queue) Enqueue(coordinate Coordinate) {
	q.List = append(q.List, coordinate)
}

// Function to remove a coordinate from the queue
func (q *Queue) Dequeue() Coordinate {
	if q.isEmpty() {
		fmt.Println("Queue is empty.")
		return Coordinate{} // Return an empty Coordinate if the queue is empty
	}
	coordinate := q.List[0]
	q.List = q.List[1:]

	return coordinate
}

// Function to check if the queue is empty
func (q *Queue) isEmpty() bool {
	return len(q.List) == 0
}

// BFS using Queue struct for managing the BFS
func BFS(matrix [][]string, start Coordinate, visited map[Coordinate]bool) int {
	rows := len(matrix)
	cols := len(matrix[0])
	area := 1
	perimeter := 0

	// Directions for moving up, down, left, and right
	directions := []Coordinate{
		{X: 0, Y: 1},  // Right
		{X: 1, Y: 0},  // Down
		{X: 0, Y: -1}, // Left
		{X: -1, Y: 0}, // Up
	}

	// Initialize the queue and enqueue the starting point
	queue := Queue{}
	queue.Enqueue(start)
	visited[start] = true

	// Perform BFS
	for !queue.isEmpty() {
		// Dequeue the first element
		current := queue.Dequeue()

		// Explore all neighbors
		for _, d := range directions {
			neighbor := Coordinate{X: current.X + d.X, Y: current.Y + d.Y}

			// Check if the neighbor is within bounds
			if neighbor.X >= 0 && neighbor.X < cols && neighbor.Y >= 0 && neighbor.Y < rows {
				// Check if the neighbor has been visited
				if !visited[neighbor] && matrix[current.Y][current.X] == matrix[neighbor.Y][neighbor.X] {
					// Mark the neighbor as visited and add to the queue
					visited[neighbor] = true
					queue.Enqueue(neighbor)
					area++ //Increment the area for every unique cell visited on this search
				} else if matrix[current.Y][current.X] != matrix[neighbor.Y][neighbor.X] {
					// If the neighbor has a different value, it's part of the perimeter
					perimeter++
				}
			} else {
				//if not in bounds, it means it is part of the perimeter
				perimeter++
			}
		}
	}

	// Calculate and return the price (perimeter * area)
	price := perimeter * area
	return price
}

// func BFS2(matrix [][]string, start Coordinate, visited map[Coordinate]bool) int {
// 	rows := len(matrix)
// 	cols := len(matrix[0])
// 	area := 1
// 	perimeter := 0
// 	edge := 0

// 	// Directions for moving up, down, left, and right
// directions := []Coordinate{
// 	{X: 0, Y: 1},  // Right
// 	{X: 1, Y: 0},  // Down
// 	{X: 0, Y: -1}, // Left
// 	{X: -1, Y: 0}, // Up
// }

// 	// Initialize the queue and enqueue the starting point
// 	queue := Queue{}
// 	queue.Enqueue(start)
// 	visited[start] = true

// 	// Perform BFS
// 	for !queue.isEmpty() {
// 		// Dequeue the first element
// 		current := queue.Dequeue()

// 		// Explore all neighbors
// 		for _, d := range directions {
// 			neighbor := Coordinate{X: current.X + d.X, Y: current.Y + d.Y}
// 			rotate90Neighbor := Coordinate{X: current.X + d.Y, Y: current.Y - d.X}
// 			diagonal := Coordinate{X: current.X + d.X + d.Y, Y: current.Y + d.Y - d.X}

// 			// Check if the neighbor is within bounds
// 			if neighbor.X >= 0 && neighbor.X < cols && neighbor.Y >= 0 && neighbor.Y < rows {
// 				// Check if the neighbor has been visited
// 				if !visited[neighbor] && matrix[current.Y][current.X] == matrix[neighbor.Y][neighbor.X] {
// 					// Mark the neighbor as visited and add to the queue
// 					visited[neighbor] = true
// 					queue.Enqueue(neighbor)
// 					area++ //Increment the area for every unique cell visited on this search
// 					fmt.Println("Neighbor found, moving in Direction: ", d.X, )
// 				} else if matrix[current.Y][current.X] != matrix[neighbor.Y][neighbor.X] {
// 					// If the neighbor has a different value, it's part of the perimeter
// 					//Check for corner: rotate 90deg, if neighbor & not in diagonal (sum of orig + rotated direc), then not corner
// 					//rotate 90deg

// 					if rotate90Neighbor.X >= 0 && rotate90Neighbor.X < cols && rotate90Neighbor.Y >= 0 && rotate90Neighbor.Y < rows &&
// 						matrix[current.Y][current.X] == matrix[rotate90Neighbor.Y][rotate90Neighbor.X] &&
// 						diagonal.X >= 0 && diagonal.X < cols && diagonal.Y >= 0 && diagonal.Y < rows &&
// 						matrix[current.Y][current.X] != matrix[diagonal.Y][diagonal.X] {
// 					} else {
// 						edge++
// 					}

// 					perimeter++

// 				} else {
// 					// if matrix[current.Y][current.X] == matrix[rotate90Neighbor.Y][rotate90Neighbor.X] && matrix[current.Y][current.X] != matrix[diagonal.Y][diagonal.X] {
// 					// 	continue
// 					// } else {
// 					// 	edge++
// 					// }
// 					//if not in bounds, it means it is part of the perimeter
// 					perimeter++
// 				}
// 			}

// 		}
// 	}
// 	fmt.Println("Edges: ", edge)
// 	// Calculate and return the price (perimeter * area)
// 	price := perimeter * area
// 	return price
// }

func BFS2(matrix [][]string, start Coordinate, visited map[Coordinate]bool) (int, int) {
	rows := len(matrix)
	cols := len(matrix[0])
	area := 1
	perimeter := 0
	corners := 0

	// Directions for moving up, down, left, and right
	directions := []Coordinate{
		{X: 0, Y: 1},  // Right
		{X: 1, Y: 0},  // Down
		{X: 0, Y: -1}, // Left
		{X: -1, Y: 0}, // Up
	}
	right := directions[0]
	down := directions[1]
	left := directions[2]
	up := directions[3]

	// Initialize the queue and enqueue the starting point
	queue := Queue{}
	queue.Enqueue(start)
	visited[start] = true
	//var lastDirection Coordinate
	rightBoolean := false
	downBoolean := false
	leftBoolean := false
	upBoolean := false

	// Perform BFS
	for !queue.isEmpty() {

		// Dequeue the first element
		current := queue.Dequeue()
		//var lastDirection Coordinate

		// Explore all neighbors
		for _, d := range directions {
			neighbor := Coordinate{X: current.X + d.X, Y: current.Y + d.Y}

			// Check if the neighbor is within bounds
			if neighbor.X >= 0 && neighbor.X < cols && neighbor.Y >= 0 && neighbor.Y < rows {
				if matrix[current.Y][current.X] == matrix[neighbor.Y][neighbor.X] {
					// Mark the neighbor as visited and add to the queue
					if !visited[neighbor] {
						visited[neighbor] = true
						queue.Enqueue(neighbor)
						area++
						//lastDirection = d
						fmt.Println("Current line: ", current, matrix[current.Y][current.X])
						fmt.Println("Neighbor: ", neighbor, matrix[neighbor.Y][neighbor.X])
						fmt.Println("Last Direction: ", d)
						fmt.Println("Neighbor found, moving in Direction: ", d.X, d.Y)
						// Check if the current edge forms a corner
						// if lastDirection.X != d.X || lastDirection.Y != d.Y {
						// 	corners++

						// 	fmt.Println("Corner found at:", current, "with direction change from:", lastDirection, "to:", d)
						// }
						if d == right {
							rightBoolean = true
						} else if d == down {
							downBoolean = true
						} else if d == left {
							leftBoolean = true
						} else if d == up {
							upBoolean = true
						}

					}

				} else {
					// Calculate perimeter
					perimeter++
					if d == right {
						rightBoolean = true
					} else if d == down {
						downBoolean = true
					} else if d == left {
						leftBoolean = true
					} else if d == up {
						upBoolean = true
					}

					// Check if the current edge forms a corner
					// if lastDirection.X != d.X || lastDirection.Y != d.Y {
					// 	corners++
					// 	fmt.Println("Corner found at:", current, "with direction change from:", lastDirection, "to:", d)
					// }
				}
			}

		}

		if !rightBoolean && !downBoolean && !leftBoolean && !upBoolean {
			corners += 2
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}
		if !rightBoolean && !downBoolean && !leftBoolean && upBoolean {
			corners += 2
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}
		if !rightBoolean && !upBoolean && !leftBoolean && downBoolean {
			corners += 2
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}
		if !downBoolean && !upBoolean && !leftBoolean && rightBoolean {
			corners += 2
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}
		if !downBoolean && !upBoolean && !rightBoolean && leftBoolean {
			corners += 2
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}

		if !rightBoolean && !downBoolean && upBoolean && leftBoolean {
			corners++
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}
		if !rightBoolean && !upBoolean && downBoolean && leftBoolean {
			corners++
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)
		}

		if !leftBoolean && !downBoolean && upBoolean && rightBoolean {
			corners++
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}

		if !leftBoolean && !upBoolean && downBoolean && rightBoolean {
			corners++
			fmt.Printf("Corner found at %d, location, %d corners total: \n", current, corners)

		}
	}

	fmt.Println("Corners found: ", corners)
	price := perimeter * area
	return price, corners
}
