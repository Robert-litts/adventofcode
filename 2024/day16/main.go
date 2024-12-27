package main

import (
	"container/heap"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day16.input", "Relative file path to use as input.")
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

	//fmt.Println(matrix)
	fmt.Println("Grid:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println("Start: ", start)
	fmt.Println("End: ", end)
	cost, steps := Dijkstra2(matrix, start, end)

	fmt.Println("Cost for part 1: ", cost)
	fmt.Println("Steps Part 2: ", steps)
	// fmt.Println("Total Steps: ", len(steps))
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

type State struct {
	pos Coordinate
	//cost int
	dir string
}

type CoordParent struct {
	cost   int
	parent []Coordinate
}

type Reindeer struct {
	x, y int
	cost int
	dir  string
	path [][2]int
}

func calcCost(curNodeDir string, currentDir string) int {
	if curNodeDir == currentDir {
		return 1
	}
	return 1001
}

func calcTurnCost(curNodeDir string, currentDir string) int {
	if curNodeDir == currentDir {
		return 0
	}
	return 1000
}

type PriorityQueue []Reindeer

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Reindeer)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// Dijkstra's algorithm implementation for finding the shortest path in a grid.
// Part 1 implementation
func Dijkstra(graph [][]string, start Coordinate, end Coordinate) int {
	visited := make(map[Coordinate]int) //store visited nodes at their lowest cost

	// Creating a Queue variable to store nodes to be processed
	var pq PriorityQueue
	heap.Init(&pq)

	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}

	directionsMap := map[[2]int]string{
		{-1, 0}: "Up",
		{1, 0}:  "Down",
		{0, -1}: "Left",
		{0, 1}:  "Right",
	}

	startPos := Reindeer{x: start.X, y: start.Y, cost: 0, dir: directionsMap[directions[3]]}
	fmt.Println("starting at: ", startPos)

	// Adding the current node to the queue
	heap.Push(&pq, &startPos)

	// Running a loop until the queue becomes empty
	for pq.Len() > 0 {
		currNode := heap.Pop(&pq).(*Reindeer)
		fmt.Println("Current Node: ", currNode)
		////fmt.Println(currNode)
		r := currNode.x
		c := currNode.y
		//fmt.Println("Current Node: ", r, c)
		if graph[r][c] == "E" {
			//fmt.Println("Summit Found at :", r, c)
			fmt.Println("End of Path Reached, Cost: ", currNode.cost)
			fmt.Println("Visited nodes: ", visited)
			fmt.Println("Visited count: ", len(visited))
			return currNode.cost
		}

		// Skip if we have already visited this node with a lower cost
		if existingCost, exists := visited[Coordinate{r, c}]; exists && existingCost <= currNode.cost {
			continue //if we already visited as a lower cost, skip
		}

		// Mark the node as visited with the current cost
		visited[Coordinate{r, c}] = currNode.cost

		for _, dir := range directions {
			currDirection := directionsMap[dir]
			nr, nc := r+dir[0], c+dir[1]
			cost := currNode.cost + 1 // Assuming each step costs 1 unit
			turnCost := calcCost(currNode.dir, currDirection)
			if 0 <= nr && nr < len(graph) && 0 <= nc && nc < len(graph[0]) && graph[nr][nc] != "#" {
				// Calculate the new cost, based on the turning mechanism (step forward = 1, turn = 1000 + 1 for step)
				totalCost := cost + turnCost

				newReindeer := &Reindeer{
					x:    nr,
					y:    nc,
					cost: totalCost,
					dir:  currDirection}
				heap.Push(&pq, newReindeer)
			}
		}
	}
	return -1
}

// Dijkstra's algorithm implementation for finding the shortest path in a grid.
// Part 2 implementation, tracks the path traveled for every node & then finds the unique set of steps at the end
func Dijkstra2(graph [][]string, start Coordinate, end Coordinate) (int, int) {
	visited := make(map[State]int)
	bestScore := math.MaxInt32
	// Creating a Queue variable to store nodes to be processed
	var pq PriorityQueue
	heap.Init(&pq)
	directions := []struct {
		dx, dy int
		name   string
	}{
		{-1, 0, "Up"},
		{1, 0, "Down"},
		{0, -1, "Left"},
		{0, 1, "Right"},
	}

	//Initial path stores the starting path for the graph
	initialPath := make([][2]int, 0, len(graph)*len(graph[0])/2)
	initialPath = append(initialPath, [2]int{start.X, start.Y})
	startState := State{pos: Coordinate{X: start.X, Y: start.Y}, dir: "Right"}
	startPos := Reindeer{x: start.X, y: start.Y, cost: 0, dir: "Right", path: initialPath}
	var tiles [][2]int
	fmt.Println("starting at: ", startPos)
	visited[startState] = 0

	// Adding the current node to the queue
	heap.Push(&pq, startPos)

	// Running a loop until the queue becomes empty
	for pq.Len() > 0 {
		currNode := heap.Pop(&pq).(Reindeer)
		if currNode.cost > bestScore {
			//currNode.path = nil
			return bestScore, len(tiles)
		}

		r := currNode.x
		c := currNode.y
		currState := State{Coordinate{r, c}, currNode.dir}
		//fmt.Println("Current Node: ", r, c)

		// Skip if we have already visited this node from same dir with a lower cost
		if existingCost, exists := visited[currState]; exists && existingCost < currNode.cost {
			continue //if we already visited as a lower cost, skip
		}
		if graph[r][c] == "E" {
			visited[State{Coordinate{r, c}, currNode.dir}] = currNode.cost

			if currNode.cost < bestScore {
				bestScore = currNode.cost
				fmt.Println("New minimum score: ", bestScore)
				// fmt.Println("final Node path: ", currNode.path)
				tiles = currNode.path

				continue

			}
			if currNode.cost == bestScore {
				fmt.Println("final Node path: ", currNode.path)
				tiles = mergeTupleTiles(tiles, currNode.path)
				continue

			}

			if currNode.cost > bestScore {

				return bestScore, len(tiles)

			}
		}

		// Mark the node as visited with the current cost
		visited[currState] = currNode.cost

		// Explore the neighbors
		for _, dir := range directions {
			currDirection := dir.name
			nr, nc := r+dir.dx, c+dir.dy
			if 0 <= nr && nr < len(graph) && 0 <= nc && nc < len(graph[0]) && graph[nr][nc] != "#" {
				// Calculate the new cost, based on the turning mechanism (step forward = 1, turn = 1000)
				stepCost := 1
				turnCost := calcTurnCost(currNode.dir, currDirection)
				cost := currNode.cost + turnCost + stepCost

				newState := State{
					pos: Coordinate{nr, nc},
					dir: dir.name,
				}
				// Skip if we've seen this state with a better cost
				if existingCost, exists := visited[newState]; exists && existingCost <= cost {
					continue
				}

				newPath := append([][2]int(nil), currNode.path...) // Copy the existing path
				newPath = append(newPath, [2]int{nr, nc})          // Append the new coordinate
				//newPath := append(basePath, [2]int{nr, nc})

				newReindeer := Reindeer{
					x:    nr,
					y:    nc,
					cost: cost,
					dir:  currDirection,
					path: newPath,
				}
				heap.Push(&pq, newReindeer)

			}
		}

	}

	return bestScore, len(tiles)
}

// Helper function to merge paths stored as [2]int tuples
func mergeTupleTiles(existing, new [][2]int) [][2]int {
	seen := make(map[[2]int]bool)

	// Add existing positions to seen map
	for _, pos := range existing {
		seen[pos] = true
	}

	// Add new positions if not already seen
	for _, pos := range new {
		if !seen[pos] {
			existing = append(existing, pos)
			seen[pos] = true
		}
	}

	return existing
}
