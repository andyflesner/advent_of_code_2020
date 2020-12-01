package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//expenseList := []int{1721, 979, 366, 299, 675, 1456}
	expenseList := getInput("inputs/day1.input")

	// Day 1 - Part 1
	match1, match2 := matchTwo(expenseList)
	if match1 < 0 || match2 < 0 {
		fmt.Println("Unable to find two entries that add up to 2020. Exiting.")
		os.Exit(1)
	}
	fmt.Printf("Day 1 - Part 1 - Match1: %d\n", expenseList[match1])
	fmt.Printf("Day 1 - Part 1 - Match2: %d\n", expenseList[match2])
	answerPart1 := expenseList[match1] * expenseList[match2]
	fmt.Printf("Day 1 - Part 1 - Answer: %d\n", answerPart1)

	// Day 1 - Part 2
	match1, match2, match3 := matchThree(expenseList)
	if match1 < 0 || match2 < 0 || match3 < 0 {
		fmt.Println("Unable to find three entries that add up to 2020. Exiting.")
		os.Exit(1)
	}
	fmt.Printf("Day 1 - Part 2 - Match 1: %d\n", expenseList[match1])
	fmt.Printf("Day 1 - Part 2 - Match 2: %d\n", expenseList[match2])
	fmt.Printf("Day 1 - Part 2 - Match 3: %d\n", expenseList[match3])
	answerPart2 := expenseList[match1] * expenseList[match2] * expenseList[match3]
	fmt.Printf("Day 1 - Part 2 - Answer: %d\n", answerPart2)
}

func twoAddTo2020(in1 int, in2 int) (match bool) {
	if in1+in2 == 2020 {
		return true
	}
	return false
}

func threeAddTo2020(in1 int, in2 int, in3 int) (match bool) {
	if in1+in2+in3 == 2020 {
		return true
	}
	return false
}

func matchTwo(expenseList []int) (int, int) {
	for i := 0; i < len(expenseList); i++ {
		for j := 0; j < len(expenseList); j++ {
			if i != j && twoAddTo2020(expenseList[i], expenseList[j]) {
				return i, j
			}
		}
	}
	return -1, -1
}

func matchThree(expenseList []int) (int, int, int) {
	for i := 0; i < len(expenseList); i++ {
		for j := 0; j < len(expenseList); j++ {
			for k := 0; k < len(expenseList); k++ {
				if i != j && j != k && i != k && threeAddTo2020(expenseList[i], expenseList[j], expenseList[k]) {
					return i, j, k
				}
			}
		}
	}
	return -1, -1, -1
}

func getInput(path string) []int {
	var contents []int
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		contents = append(contents, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error occurred during input processing: %d", err)
	}
	return contents
}
