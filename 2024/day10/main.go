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
	var inputFile = flag.String("inputFile", "../input/day10.example1", "Relative file path to use as input.")
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
	scoreP1 := 0
	scoreP2 := 0
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// makeMatrix function will create the matrix from the input
	matrix, zeros, err := makeMatrix(bytes)
	if err != nil {
		return err
	}

	//fmt.Println(matrix)
	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	//Conduct breadth first search from each of the 0 positions on the map
	for rows := range len(zeros) {
		summitsP1 := BFS1(matrix, zeros[rows].row, zeros[rows].col)
		summitsP2 := BFS2(matrix, zeros[rows].row, zeros[rows].col)

		//fmt.Println("Summits from current run: ", summits)
		scoreP1 += summitsP1
		scoreP2 += summitsP2

	}

	//fmt.Println("Summits Visited: ", len(summits))
	fmt.Println("Score Part 1: ", scoreP1)
	fmt.Println("Score Part 2: ", scoreP2)
	return nil
}

///////////////////////////////////////////////////////////////////////////////

// makeMatrix takes the byte slice, splits it into lines and converts it into a 2D integer matrix.
// Also returns 0 positions in the matrix.
func makeMatrix(bytes []byte) ([][]int, []struct{ row, col int }, error) {
	var matrix [][]int
	var zeroPositions []struct{ row, col int }
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	for rowIndex, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Row to hold the integers
		row := make([]int, len(line))

		for colIndex, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, nil, fmt.Errorf("error converting value '%s' to int: %v", string(c), err)
			}

			// Assign the integer to the row
			row[colIndex] = num

			// Check if the number is 0 and record the position
			if num == 0 {
				zeroPositions = append(zeroPositions, struct{ row, col int }{rowIndex, colIndex})
			}
		}

		// Append the row to the matrix
		matrix = append(matrix, row)
	}

	// Return the created matrix and the list of zero positions
	return matrix, zeroPositions, nil
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

// Part 1 Breadth First Search, tracking only unique paths
func BFS1(graph [][]int, r, c int) int {
	// Initializing the map that will keep track if the node is visited
	visited := make(map[Coordinate]bool)

	// Creating a Queue variable to store nodes to be processed
	var bfsQueue Queue

	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
	var summits []Summit

	// Marking current node as visited
	//isvisited[node] = true

	// Adding the current node to the queue
	bfsQueue.Enqueue(Coordinate{X: r, Y: c})
	visited[Coordinate{X: r, Y: c}] = true

	// Running a loop until the queue becomes empty
	for !bfsQueue.isEmpty() {
		currNode := bfsQueue.Dequeue()
		//fmt.Println(currNode)
		r := currNode.X
		c := currNode.Y
		//fmt.Println("Current Node: ", r, c)
		if graph[r][c] == 9 {
			//fmt.Println("Summit Found at :", r, c)
			summits = append(summits, Summit{X: r, Y: c})
		}

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if 0 <= nr && nr < len(graph) && 0 <= nc && nc < len(graph[0]) && graph[r][c]+1 == graph[nr][nc] && !visited[Coordinate{X: nr, Y: nc}] {
				visited[Coordinate{X: nr, Y: nc}] = true
				bfsQueue.Enqueue(Coordinate{X: nr, Y: nc})

			}
		}

		// Remove the current node from the queue after visiting it
	}
	//fmt.Println("Summits Found: ", summits)
	return len(summits)

}

// Part 2 BFS where the only difference is we do not track unique paths
func BFS2(graph [][]int, r, c int) int {
	// Initializing the map that will keep track if the node is visited
	visited := make(map[Coordinate]bool)

	// Creating a Queue variable to store nodes to be processed
	var bfsQueue Queue

	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
	var summits []Summit

	// Marking current node as visited
	//isvisited[node] = true

	// Adding the current node to the queue
	bfsQueue.Enqueue(Coordinate{X: r, Y: c})
	visited[Coordinate{X: r, Y: c}] = true

	// Running a loop until the queue becomes empty
	for !bfsQueue.isEmpty() {
		currNode := bfsQueue.Dequeue()
		//fmt.Println(currNode)
		r := currNode.X
		c := currNode.Y
		//fmt.Println("Current Node: ", r, c)
		if graph[r][c] == 9 {
			//fmt.Println("Summit Found at :", r, c)
			summits = append(summits, Summit{X: r, Y: c})
		}

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if 0 <= nr && nr < len(graph) && 0 <= nc && nc < len(graph[0]) && graph[r][c]+1 == graph[nr][nc] {
				visited[Coordinate{X: nr, Y: nc}] = true
				bfsQueue.Enqueue(Coordinate{X: nr, Y: nc})

			}
		}

		// Remove the current node from the queue after visiting it
	}
	//fmt.Println("Summits Found: ", summits)
	return len(summits)

}
