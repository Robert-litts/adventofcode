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
	readLock := false
	readKey := false
	var lockkey string
	contents := string(bytes)
	lines := strings.Split(contents, "\n\n") //split on double newline to get each unique key or lock

	for set, line := range lines { // this is each set of lock or key values, need to determine if key or lock
		readKey, readLock = false, false
		nums := strings.Split(line, "\n")
		for idx, k := range nums { // this is each individual line within a lock or key.
			if idx == 0 {
				lockkey = lockOrKey(k)
				if lockkey == "lock" {
					readLock = true
					readKey = false
				} else if lockkey == "key" {
					readLock = false
					readKey = true
				}
			}
			//Assign values to the 5 positions in the lock or key based on "#"
			if idx > 0 && idx < 6 && (readKey || readLock) {
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
				fmt.Println("Lock: ", locks[lock], " Key: ", keys[key], " is valid")
			}
		}
	}

	fmt.Println("Valid: ", valid)

	return nil
}

// Determine if we are viewing a lock or key based on first line
func lockOrKey(input string) string {
	//Check if the input is a lock or a key
	var lockorkey string
	if input == "#####" {
		lockorkey = "lock"
	} else if input == "....." {
		lockorkey = "key"
	}
	return lockorkey
}

// Compare the keys and locks, if the sum of the keys and locks is greater than 5, then it is not valid
func compareKeys(keys, locks [5]int) bool {
	//Compare the keys and locks

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		lock := locks[i]
		if (key + lock) > 5 {
			return false
		}

	}
	return true
}
