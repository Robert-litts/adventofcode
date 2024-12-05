package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile = flag.String("inputFile", "../input/day5.example", "Relative file path to use as input.")
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

	var column1 []int
	var column2 []int
	var update_list [][]int

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		rules := strings.Split(line, "|")   //Split at "|" to create the rules
		updates := strings.Split(line, ",") //Split at ","" to loop over page number updates

		// Convert the first part to an integer and append to column1
		if len(rules) >= 2 {
			firstColumn, err := strconv.Atoi(rules[0])
			if err != nil {
				fmt.Println("Error converting first column value:", err)
				continue
			}
			column1 = append(column1, firstColumn)

			// Convert the second part to an integer and append to column2
			secondColumn, err := strconv.Atoi(rules[1])
			if err != nil {
				fmt.Println("Error converting second column value:", err)
				continue
			}
			column2 = append(column2, secondColumn)
		}

		if len(updates) >= 2 {
			//Create update_list
			var current_row []int
			for _, item := range updates {
				num, err := strconv.Atoi(item)
				if err != nil {
					fmt.Println("Error converting update value:", err)
					continue
				}
				current_row = append(current_row, num)
			}
			update_list = append(update_list, current_row)
		}
	}

	//Iterate thru update list, look for violations by checking each rule
	for _, each := range update_list {
		violations := 0
		//Iterate thru each rule, find index where a rule matches
		for y, rule := range column1 {
			index1 := -1 //initialize index to a value that is not possible
			index2 := -1
			//As we go thru each rule, check if it matches any value in the row
			for z, item := range each {
				//If the first rule matches the value in the update list, get the index
				if rule == item {
					index1 = z
				}
				//If the second rule matches the value in the update list, get the index
				if column2[y] == item {
					index2 = z
				}

			}
			//Now check if we have stored the index of any items matching the current rule
			//If index of the item from the first column > the index of the second, violation & break

			if index1 != -1 && index2 != -1 && index1 > index2 {
				violations++
				break
			}
		}
		if violations == 0 {
			score += each[len(each)/2] //Add the middle number to the score if all rules passed!
		}
	}

	fmt.Println("Score Part 1: ", score)
	return nil
}
