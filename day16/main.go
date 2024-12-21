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
	var inputFile = flag.String("inputFile", "../input/day16.example2", "Relative file path to use as input.")
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
	//fmt.Println("Node Count: ", len(nodeMap.List))

	// //Conduct breadth first search from each of the 0 positions on the map
	// for rows := range len(zeros) {
	// 	summitsP1 := BFS1(matrix, zeros[rows].row, zeros[rows].col)
	// 	summitsP2 := BFS2(matrix, zeros[rows].row, zeros[rows].col)

	// 	//fmt.Println("Summits from current run: ", summits)
	// 	scoreP1 += summitsP1
	// 	scoreP2 += summitsP2

	// }

	//distance := Dijkstra(matrix, start, end)
	cost, steps := Dijkstra2(matrix, start, end)

	//fmt.Println("Summits Visited: ", len(summits))
	//fmt.Println("Score Part 1: ", distance)
	fmt.Println("Cost for part 2: ", cost)
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
	pos  Coordinate
	cost int
	dir  string
}

type CoordParent struct {
	cost   int
	parent []Coordinate
}

type Reindeer struct {
	x, y    int
	cost    int
	dir     string
	prevDir string
	parent  Coordinate
	path    [][2]int
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

	startPos := Reindeer{x: start.X, y: start.Y, cost: 0, dir: directionsMap[directions[3]], prevDir: ""}
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
					x:       nr,
					y:       nc,
					cost:    totalCost,
					dir:     currDirection,
					prevDir: currNode.dir,
				}
				heap.Push(&pq, newReindeer)
			}
		}
	}
	return -1
}

// Dijkstra's algorithm implementation for finding the shortest path in a grid.
func Dijkstra2(graph [][]string, start Coordinate, end Coordinate) (int, int) {
	//visited := make(map[Coordinate]CoordParent) //store visited nodes at their lowest cost
	visited := make(map[State]int)
	//prevNode := Coordinate{}
	//reindeers := make(map[Coordinate]*Reindeer)
	bestScore := math.MaxInt32

	// Creating a Queue variable to store nodes to be processed
	var pq PriorityQueue
	heap.Init(&pq)

	// directions := [][2]int{
	// 	{-1, 0}, // Up
	// 	{1, 0},  // Down
	// 	{0, -1}, // Left
	// 	{0, 1},  // Right
	// }

	directions := []struct {
		dx, dy int
		name   string
	}{
		{-1, 0, "Up"},
		{1, 0, "Down"},
		{0, -1, "Left"},
		{0, 1, "Right"},
	}

	// directionsMap := map[[2]int]string{
	// 	{-1, 0}: "Up",
	// 	{1, 0}:  "Down",
	// 	{0, -1}: "Left",
	// 	{0, 1}:  "Right",
	// }
	// Pre-allocate the initial path slice - each position takes 2 integers
	initialPath := make([][2]int, 0, len(graph)*len(graph[0])/2)
	initialPath = append(initialPath, [2]int{start.X, start.Y})
	startState := State{pos: Coordinate{X: start.X, Y: start.Y}, cost: 0, dir: "Right"}
	startPos := Reindeer{x: start.X, y: start.Y, cost: 0, dir: "Right", prevDir: "", path: initialPath}
	//reindeers[start] = &startPos
	var tiles [][2]int //prevNode = Coordinate{X: start.X, Y: start.Y}
	//visited[Coordinate{X: start.X, Y: start.Y}] = 0
	fmt.Println("starting at: ", startPos)
	visited[startState] = 0
	//previous := make(map[State]State)

	// Adding the current node to the queue
	heap.Push(&pq, startPos)
	//var bestPath []Coordinate

	// Running a loop until the queue becomes empty
	for pq.Len() > 0 {
		currNode := heap.Pop(&pq).(Reindeer)
		if currNode.cost > bestScore {
			//currNode.path = nil
			return bestScore, len(tiles)
		}
		//fmt.Println("Current Node: ", currNode)
		//fmt.Println("Current Parent: ", currNode.parent)
		//fmt.Println(currNode)
		r := currNode.x
		c := currNode.y
		currState := State{Coordinate{r, c}, currNode.cost, currNode.dir}
		//fmt.Println("Current Node: ", r, c)

		// Skip if we have already visited this node from same dir with a lower cost
		if existingCost, exists := visited[currState]; exists && existingCost < currNode.cost {
			continue //if we already visited as a lower cost, skip
		}
		if graph[r][c] == "E" {
			//fmt.Println("Summit Found at :", r, c)
			visited[State{Coordinate{r, c}, currNode.cost, currNode.dir}] = currNode.cost
			//fmt.Println("End of Path Reached, Cost: ", currNode.cost)
			//fmt.Println("Visited nodes: ", visited)
			//fmt.Println("Visited count: ", len(visited))
			//fmt.Println("Reindeers: ", len(reindeers))
			// for coord := range reindeers {
			// 	fmt.Println("Reindeer at: ", coord)
			// 	fmt.Println("Reindeer parent: ", reindeers[coord].parent)
			// }
			//steps := countStepsToDestination(visited, Coordinate{r, c})
			//steps := len(visited)
			//fmt.Println("final Node path: ", currNode.path)

			if currNode.cost < bestScore {
				bestScore = currNode.cost
				fmt.Println("New minimum score: ", bestScore)
				fmt.Println("final Node path: ", currNode.path)
				//bestPath = reconstructPath(previous, currState)
				//fmt.Println("Best Path: ", bestPath)
				tiles = currNode.path
				continue

			}
			if currNode.cost == bestScore {
				fmt.Println("final Node path: ", currNode.path)
				//fmt.Println("Best Path: ", bestPath)
				tiles = mergeTupleTiles(tiles, currNode.path)
				continue

			}

			if currNode.cost > bestScore {

				return bestScore, len(tiles)

			}
		}

		// Mark the node as visited with the current cost
		visited[currState] = currNode.cost

		//basePath := make([]Coordinate, len(currNode.path), len(currNode.path)+1)
		//copy(basePath, currNode.path)
		//basePath := append(make([]Coordinate, len(currNode.path)), currNode.path...)
		// Pre-allocate the new path slice with exact capacity needed
		// basePath := make([][2]int, len(currNode.path), len(currNode.path)+1)
		// copy(basePath, currNode.path)
		for _, dir := range directions {
			currDirection := dir.name
			nr, nc := r+dir.dx, c+dir.dy
			if 0 <= nr && nr < len(graph) && 0 <= nc && nc < len(graph[0]) && graph[nr][nc] != "#" {
				// Calculate the new cost, based on the turning mechanism (step forward = 1, turn = 1000 + 1 for step)
				stepCost := 1
				turnCost := calcTurnCost(currNode.dir, currDirection)
				cost := currNode.cost + turnCost + stepCost

				// newPath := append([]Coordinate(nil), currNode.path...) // Create a new slice with a copy of the current path
				// newPath = append(newPath, Coordinate{X: nr, Y: nc})    // Append the new coordinate to the path

				newState := State{
					pos:  Coordinate{nr, nc},
					dir:  dir.name,
					cost: cost,
				}

				//basePath = append(basePath, Coordinate{X: nr, Y: nc})

				// Skip if we've seen this state with a better cost
				if existingCost, exists := visited[newState]; exists && existingCost <= cost {
					continue
				}

				//previous[newState] = currState
				//newPath := append(basePath, Coordinate{X: nr, Y: nc})
				newPath := append([][2]int(nil), currNode.path...) // Copy the existing path
				newPath = append(newPath, [2]int{nr, nc})          // Append the new coordinate
				//newPath := append(basePath, [2]int{nr, nc})

				newReindeer := Reindeer{
					x:       nr,
					y:       nc,
					cost:    cost,
					dir:     currDirection,
					prevDir: currNode.dir,
					//parent:  Coordinate{r, c},
					//path: append([]Coordinate(nil), currNode.path...),
					path: newPath,
				}
				//newReindeer.path = append(newReindeer.path, Coordinate{X: nr, Y: nc})
				heap.Push(&pq, newReindeer)
				//prevNode = Coordinate{r, c}
				//reindeers[Coordinate{nr, nc}] = newReindeer
			}
		}

	}
	fmt.Println("Total Tiles: ", len(tiles))
	fmt.Println("Tiles: ", tiles)
	fmt.Println("Visited: ", visited)
	return bestScore, len(tiles)
}

// func countStepsToDestination(visited map[Coordinate]CoordParent, end Coordinate) int {
// 	// Initialize a counter to keep track of the number of steps
// 	stepCount := 0

// 	// Backtrack from the destination node
// 	currNode := end

// 	// Loop until we reach a node with no parent (i.e., the starting node)
// 	for {
// 		// If there are no parents for the current node, we've reached the start
// 		if parentNode, exists := visited[currNode]; exists {
// 			// Increment step count
// 			stepCount++
// 			//fmt.Println("New Step: ", currNode, visited[currNode])

// 			// Move to one of the parent nodes (typically you will have only one parent)
// 			if len(parentNode.parent) > 0 {
// 				currNode = parentNode.parent[0] // Move to the first parent node
// 			} else {
// 				break // If no parent, we reached the start node
// 			}
// 		} else {
// 			// If there's no entry for this node, we're done (or there was an error)
// 			break
// 		}
// 	}

// 	return stepCount
// }

// // Helper function to find all possible paths from start to end
// func findAllPaths(parents map[Coordinate][]Coordinate, start, end Coordinate) [][]Coordinate {
// 	if start == end {
// 		return [][]Coordinate{{end}}
// 	}

// 	var paths [][]Coordinate
// 	parentCoords, exists := parents[end]
// 	if !exists {
// 		return paths
// 	}

// 	for _, parent := range parentCoords {
// 		subPaths := findAllPaths(parents, start, parent)
// 		for _, subPath := range subPaths {
// 			newPath := append(subPath, end)
// 			paths = append(paths, newPath)
// 		}
// 	}

// 	return paths
// }

// Helper function to merge two paths of coordinates without duplicates
func mergeTiles(existing, new []Coordinate) []Coordinate {
	seen := make(map[string]bool)
	result := make([]Coordinate, 0)

	// Add all existing tiles
	for _, coord := range existing {
		key := fmt.Sprintf("%d,%d", coord.X, coord.Y)
		if !seen[key] {
			seen[key] = true
			result = append(result, coord)
		}
	}

	// Add new tiles if they haven't been seen
	for _, coord := range new {
		key := fmt.Sprintf("%d,%d", coord.X, coord.Y)
		if !seen[key] {
			seen[key] = true
			result = append(result, coord)
		}
	}

	return result
}

// // Helper function to reconstruct the path
// func reconstructPath(previous map[State]State, endState State) []Coordinate {
// 	path := []Coordinate{endState.pos}
// 	current := endState

// 	// Traverse backwards through the previous states until we can't go further
// 	for {
// 		prev, exists := previous[current]
// 		if !exists {
// 			break
// 		}
// 		path = append([]Coordinate{prev.pos}, path...)
// 		current = prev
// 	}

// 	return path
// }

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
