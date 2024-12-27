package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day20.input", "Relative file path to use as input.")
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
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, start, end, err := makeMatrix(bytes)
	if err != nil {
		return err
	}
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Println(start, end)
	visited := make(map[Coordinate]bool)
	directions := []Coordinate{
		{Y: 0, X: 1},  // Right
		{Y: 1, X: 0},  // Down
		{Y: 0, X: -1}, // Left
		{Y: -1, X: 0}, // Up
	}
	cheats := 0

	//fmt.Println(matrix)
	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println("Start: ", start)
	fmt.Println("End: ", end)

	//Get the baseline path + steps to end
	steps, basePath := BFS(matrix, start, end, visited)
	fmt.Println("Steps to reach the end:", steps)
	fmt.Println("Path to reach the end:", basePath)

	for p, coord := range basePath {
		//p represents the current number of steps from the start
		//fmt.Println(coord)
		//check surrounding cells for walls, check two cells in each dir

		for _, d := range directions {
			visited = make(map[Coordinate]bool) //reset visited grid

			neighbor1 := Coordinate{
				Y: coord.Y + d.Y,
				X: coord.X + d.X,
			}
			neighbor2 := Coordinate{
				Y: coord.Y + d.Y*2,
				X: coord.X + d.X*2,
			}

			// Check if the neighbors are within the grid bounds and first is a wall, and second is a . and second is within the original basePath
			if isWithinGrid(neighbor1, rows, cols) && isWithinGrid(neighbor2, rows, cols) && matrix[neighbor1.Y][neighbor1.X] == "#" && NeighborInBasePath(neighbor2, basePath) && (matrix[neighbor2.Y][neighbor2.X] == "." || matrix[neighbor2.Y][neighbor2.X] == "E") {
				//fmt.Println("Found a wall at:", neighbor1)

				// check if the neighbor2 is within the basePath
				for i, coord := range basePath {
					if coord == neighbor2 {
						//We now know that this cheat puts us back on the base path
						//i is the index of neighbor2 in the base path, need to calculate how many steps left to the end & compare
						//We have already gone "p" steps along the track, and 1 step to the wall + 1 step past
						//We are now at "i" steps along the track, so we have "steps - i" steps left to the end
						timeToEnd := p + 2 + (steps - i)

						timeSaved := steps - timeToEnd
						if timeSaved >= 100 {
							fmt.Println("Time Saved: ", timeSaved)
							cheats++

						}

					}
				}

			}
		}
	}
	fmt.Println("Total cheats over 100: ", cheats)
	return nil
}

// Struct to represent a coordinate on the grid. X and Y are integers.
type Coordinate struct {
	X int
	Y int
}

type State struct {
	coordinate Coordinate
	steps      int
	path       []Coordinate
}

// Define the Queue struct to hold Coordinates
type Queue struct {
	List []State
}

// Function to add a coordinate to the queue
func (q *Queue) Enqueue(state State) {
	q.List = append(q.List, state)
}

// Function to remove a coordinate from the queue
func (q *Queue) Dequeue() State {
	if q.isEmpty() {
		fmt.Println("Queue is empty.")
		return State{} // Return an empty Coordinate if the queue is empty
	}
	state := q.List[0]
	q.List = q.List[1:]

	return state
}

// Function to check if the queue is empty
func (q *Queue) isEmpty() bool {
	return len(q.List) == 0
}

func makeMatrix(bytes []byte) ([][]string, Coordinate, Coordinate, error) {
	var matrix [][]string
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var start, end Coordinate

	for rowIndex, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]string, len(line))
		for colIndex, c := range line {
			row[colIndex] = string(c)

			if string(c) == "S" {
				start = Coordinate{Y: rowIndex, X: colIndex}
			}
			if string(c) == "E" {
				end = Coordinate{Y: rowIndex, X: colIndex}
			}
		}
		matrix = append(matrix, row)
	}
	return matrix, start, end, nil
}

func BFS(matrix [][]string, start Coordinate, end Coordinate, visited map[Coordinate]bool) (int, []Coordinate) {
	rows := len(matrix)
	cols := len(matrix[0])

	//fmt.Printf("Start character: %s\n", matrix[start.Y][start.X])

	startState := State{coordinate: start, steps: 0, path: []Coordinate{start}}

	directions := []Coordinate{
		{Y: 0, X: 1},  // Right
		{Y: 1, X: 0},  // Down
		{Y: 0, X: -1}, // Left
		{Y: -1, X: 0}, // Up
	}

	queue := Queue{}
	queue.Enqueue(startState)
	visited[start] = true

	for !queue.isEmpty() {
		current := queue.Dequeue()

		if current.coordinate == end {
			// fmt.Println("BFS Complete, Ending at: ", current.coordinate)
			// fmt.Println("Path: ", current.path)
			return current.steps, current.path
		}

		for _, d := range directions {
			neighbor := Coordinate{
				Y: current.coordinate.Y + d.Y,
				X: current.coordinate.X + d.X,
			}

			if neighbor.Y >= 0 && neighbor.Y < rows &&
				neighbor.X >= 0 && neighbor.X < cols {
				if !visited[neighbor] && matrix[neighbor.Y][neighbor.X] != "#" {
					visited[neighbor] = true
					neighborState := State{
						coordinate: neighbor,
						steps:      current.steps + 1,
						path:       append(current.path, neighbor),
					}
					queue.Enqueue(neighborState)
				}
			}
		}
	}
	return -1, nil
}

func isWithinGrid(coord Coordinate, rows, cols int) bool {
	return coord.Y >= 0 && coord.Y < rows && coord.X >= 0 && coord.X < cols
}

func NeighborInBasePath(neighbor2 Coordinate, basePath []Coordinate) bool {
	for _, coord := range basePath {
		if coord == neighbor2 {
			return true
		}
	}
	return false
}

//Significantly slower, alternative method for part 1, use BFS from each neighbor location.
// currentSteps, currentPath := BFS(matrix, neighbor2, end, visited)
// 				//fmt.Println("Removed wall, result: ", currentSteps)
// 				//check if steps is not -1, path is not nil, and path is shorter than the base path
// 				//Step count would be steps up to the cheat, p, 2 steps for the cheat and then steps from the cheat to the end
// 				if currentSteps != -1 && currentPath != nil && currentSteps+p+2 < steps {
// 					//fmt.Println("Current Path shorter than base path")
// 					//fmt.Println("Current total steps: ", currentSteps)
// 					timeSaved := steps - (currentSteps + p + 2)
// 					if timeSaved >= 100 {
// 						//fmt.Println("Time Saved: ", timeSaved)
// 						cheats++

// 					}
// 				}
// 				//add the wall back
// 				matrix[neighbor1.Y][neighbor1.X] = "#"
