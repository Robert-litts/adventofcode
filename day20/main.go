package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day20.example", "Relative file path to use as input.")
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
	fmt.Println(start, end, matrix[start.Y][start.X])
	fmt.Println("Matrix: ", matrix[3][1])
	visited := make(map[Coordinate]bool)

	//fmt.Println(matrix)
	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println("Start: ", start)
	fmt.Println("End: ", end)

	steps := BFS(matrix, start, end, visited)
	fmt.Println("Steps to reach the end:", steps)

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

// Then in your BFS function, when accessing the matrix:
func BFS(matrix [][]string, start Coordinate, end Coordinate, visited map[Coordinate]bool) int {
	rows := len(matrix)
	cols := len(matrix[0])

	fmt.Printf("Start character: %s\n", matrix[start.Y][start.X])

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
			fmt.Println("BFS Complete, Ending at: ", current.coordinate)
			fmt.Println("Path: ", current.path)
			return current.steps
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
	return -1
}
