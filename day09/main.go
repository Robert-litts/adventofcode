package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day9.example", "Relative file path to use as input.")
	flag.Parse()
	fmt.Println("Running Part 1:")
	if err := Part1(*inputFile); err != nil {
		fmt.Println("Error in Part 1:", err)
		return
	}
	// fmt.Println("Running Part 2:")
	// if err := Part2(*inputFile); err != nil {
	// 	fmt.Println("Error in Part 2:", err)
	// 	return
	// }
}

func Part1(inputFile string) error {
	score := 0
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	contents := string(bytes)

	var fileBlocks []int
	var freeSpace []int
	var freeSpaceIndex []int
	var diskMap []string
	var digits []string

	//lines := strings.Split(contents, "\n")
	file := true
	for i := 0; i < len(contents); i++ {
		storeInt, err := strconv.Atoi(string(contents[i]))
		if err != nil {
			fmt.Println("Error converting first column value:", err)
			continue
		}

		//fmt.Println("Character: ", storeInt)
		if file {

			fileBlocks = append(fileBlocks, storeInt)
			file = false
		} else if !file {
			freeSpace = append(freeSpace, storeInt)
			file = true
		}

	}

	//fmt.Println("Fileblocks: ", fileBlocks)
	//fmt.Println("Freespace: ", freeSpace)

	for i := 0; i < len(fileBlocks); i++ {
		for j := 0; j < fileBlocks[i]; j++ {
			diskMap = append(diskMap, strconv.Itoa(i))
		}

		if i < len(freeSpace) {
			for k := 0; k < freeSpace[i]; k++ {
				diskMap = append(diskMap, ".")
				freeSpaceIndex = append(freeSpaceIndex, len(diskMap)-1)
			}

		}

	}
	//fmt.Println("Disk Map, Initial: ", diskMap)
	//fmt.Println("Free Space Index: ", freeSpaceIndex)

	for j := len(diskMap) - 1; j >= 0; j-- {
		// Skip already processed or free space slots
		if diskMap[j] != "." {
			digits = append(digits, diskMap[j])
		}
	}

	numDigits := len(digits)

	for i := 0; i < len(freeSpaceIndex); i++ {
		//fmt.Println("Current DiskMap: ", diskMap)
		// Start j from the last index of the diskMap
		for j := len(diskMap) - 1; j >= numDigits; j-- {
			// Skip free space slots
			if diskMap[j] == "." {
				continue
			}

			// Swap values when diskMap[j] is not a free space and freeSpaceIndex[i] is valid
			diskMap[freeSpaceIndex[i]] = diskMap[j]
			diskMap[j] = "."

			// Break after the swap is done for this free space
			break
		}
	}

	//fmt.Println("Disk Map, Final: ", diskMap)

	//Calculate checksum
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != "." {
			int, err := strconv.Atoi(string(diskMap[i]))
			if err != nil {
				fmt.Println("Error converting first column value:", err)
				continue
			}
			score += int * i

		}
	}

	//fmt.Println("Updated Score: ", new_score)

	fmt.Println("Score Part 1: ", score)
	return nil
}
