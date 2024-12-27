package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day25.input", "Relative file path to use as input.")
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

	keys := make(map[int][5]int)
	locks := make(map[int][5]int)
	valid := 0
	contents := string(bytes)
	lines := strings.Split(contents, "\n\n") //split on double newline to get each unique key or lock

	for set, line := range lines { // this is each set of lock or key values, need to determine if key or lock
		var readKey, readLock bool
		nums := strings.Split(line, "\n")
		lockkey := lockOrKey(nums[0])

		if lockkey == "lock" {
			readLock = true
		} else if lockkey == "key" {
			readKey = true
		}
		for idx, k := range nums { // this is each individual line within a lock or key.
			if idx > 0 && idx < 6 {
				for p, each := range k {
					if readLock && string(each) == "#" {
						val := locks[set]
						val[p]++
						locks[set] = val
					}
					if readKey && string(each) == "#" {
						val := keys[set]
						val[p]++
						keys[set] = val
					}
				}

			}

		}

	}

	for lock := range locks {
		for key := range keys {
			if compareKeys(locks[lock], keys[key]) {
				valid++
			}
		}
	}

	fmt.Println("Valid: ", valid)

	return nil
}

// Determine if we are viewing a lock or key based on first line
func lockOrKey(input string) string {
	lockOrKeyMap := map[string]string{
		"#####": "lock",
		".....": "key",
	}
	return lockOrKeyMap[input]
}

// Compare the keys and locks, if the sum of the keys and locks is greater than 5, then it is not valid
func compareKeys(keys, locks [5]int) bool {
	//Compare the keys and locks

	for i := 0; i < len(keys); i++ {
		if (keys[i] + locks[i]) > 5 {
			return false
		}
	}
	return true
}
