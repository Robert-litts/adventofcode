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
	var inputFile = flag.String("inputFile", "../input/day18.input", "Relative file path to use as input.")
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
	size := 70
	matrix, err := makeMatrix(size + 1)
	var corruption []Coordinate
	visited := make(map[Coordinate]bool)
	start := Coordinate{X: 0, Y: 0}
	end := Coordinate{X: size, Y: size}
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

	for x, coord := range corruption {
		visited = make(map[Coordinate]bool)
		matrix[coord.Y][coord.X] = "#" //Mark the corruption on the grid
		steps := BFS(matrix, start, end, visited)
		//Part 1, After 1024 bytes have fallen
		if x == 1024 {
			fmt.Println("Part 1 answer: ", steps)
			continue
		}
		//Part 2, Find first coordinate where no path can be found (BFS returns -1)
		if steps == -1 {
			fmt.Println("No path found, final corruption: ", coord)
			return nil
		}
	}
	// fmt.Println("Grid:")
	// for _, row := range matrix {
	// 	fmt.Println(row)
	// }

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

type State struct {
	coordinate Coordinate
	steps      int
}

// BFS using Queue struct for managing the BFS
func BFS(matrix [][]string, start Coordinate, end Coordinate, visited map[Coordinate]bool) int {
	rows := len(matrix)
	cols := len(matrix[0])
	startState := State{coordinate: start, steps: 0}

	// Directions for moving up, down, left, and right
	directions := []Coordinate{
		{X: 0, Y: 1},  // Right
		{X: 1, Y: 0},  // Down
		{X: 0, Y: -1}, // Left
		{X: -1, Y: 0}, // Up
	}

	// Initialize the queue and enqueue the starting point
	queue := Queue{}
	queue.Enqueue(startState)
	visited[start] = true

	// Perform BFS
	for !queue.isEmpty() {
		// Dequeue the first element
		current := queue.Dequeue()
		if current.coordinate.X == end.X && current.coordinate.Y == end.Y {
			// If the neighbor has a different value, it's part of the perimeter

			//fmt.Println("Ending at: ", current.coordinate.X, current.coordinate.Y)
			return current.steps
		}

		// Explore all neighbors
		for _, d := range directions {
			neighbor := Coordinate{X: current.coordinate.X + d.X, Y: current.coordinate.Y + d.Y}
			neighborState := State{coordinate: neighbor, steps: current.steps + 1}

			// Check if the neighbor is within bounds
			if neighbor.X >= 0 && neighbor.X < cols && neighbor.Y >= 0 && neighbor.Y < rows {
				// Check if the neighbor has been visited
				if !visited[neighbor] && matrix[current.coordinate.Y][current.coordinate.X] == "." {
					// Mark the neighbor as visited and add to the queue
					visited[neighbor] = true
					queue.Enqueue(neighborState)

				}
			}
		}
	}

	return -1
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
