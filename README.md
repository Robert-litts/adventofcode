# [Advent of Code](https://adventofcode.com) 🎄

### Overview & Solutions 

Solutions to [Advent of Code](https://adventofcode.com) challenges, a yearly coding event started in 2015 where a new puzzle is released each day from December 1st to December 25th. Solutions are categorized by year and then by day. Check out a behind the scenes look at AoC by the creator, Eric Wastl, during his [2024 keynote at CPP North](https://www.youtube.com/watch?v=uZ8DcbhojOw) for AoC's 10th anniversary.

Puzzles can be solved in any language, follow a holiday-themed storyline, and cover a wide variety of programming topics including, but not limited to: 
- **Algorithms & Graph Theory**  
  - Pathfinding algorithms (e.g., Dijkstra, breadth-first search, depth-first search, A*)
  - Dynamic programming
  - Greedy algorithms
  - Graph traversal and searching

- **Bitwise Operations**  
  - Bit shifts
  - Bitwise masks
  - Bitwise operations (AND, OR, XOR)

- **Data Structures**  
  - Arrays and lists
  - Stacks and queues
  - Hash maps and dictionaries
  - Heaps (min/max heaps)
  - Trees (binary trees, balanced trees, etc.)

- **Cryptography**  
  - Encoding and decoding
  - Hashing algorithms (e.g., MD5, SHA)
  - Encryption algorithms (e.g., Caesar cipher, RSA)

- **Combinatorics**  
  - Permutations and combinations
  - Subsets and subset sums
  - Counting distinct possibilities and arrangements

- **Game Theory**  
  - Two-player games
  - Turn-based mechanics
  - Determining optimal strategies

- **Number Theory**  
  - Finding primes (e.g., Sieve of Eratosthenes)
  - Calculating greatest common divisors (GCD)
  - Solving Diophantine equations
  - Modular arithmetic

- **Computational Geometry**  
  - Working with grids and coordinate systems
  - Calculating intersections (lines, circles, polygons)
  - Simulating physical systems (e.g., collision detection, convex hull)

- **Recursion**  
  - Traversing tree-like structures
  - Divide-and-conquer algorithms
  - Backtracking algorithms (e.g., solving puzzles, generating permutations)

- **File & Data Parsing**  
  - Handling large input datasets
  - Converting data into usable formats (e.g., CSV, JSON, XML)
  - Data extraction and transformation

### Automatically Download Input (2024 Edition in Go)
To automatically download your input, create a `.env` file in the `input` directory, add your session cookie and then run input/get_input.go with the day flag set to the input you'd like to grab:

```bash
cd 2024
touch .env
echo 'SESSION_COOKIE="YOUR SESSION COOKIE"' > .env
go run input/get_input.go --day 1
```

