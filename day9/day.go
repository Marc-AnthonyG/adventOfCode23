package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(input string) {
	fmt.Println("Day 9")
	fmt.Println("Part 1", Part1(input))
	fmt.Println("Part 2", Part2(input))
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	list := make([][]int, 0)

	for _, line := range lines[1:] {
		if line == "" { continue }
		numberStrs := strings.Fields(line)
		newList := make([]int, 0)
		for _, numberStr := range numberStrs {
			number, _ := strconv.Atoi(numberStr)
			newList = append(newList, number)
		}
		list = append(list, newList)
	}

	return list 
}


func Part1(input string) int {
	inputs := parseInput(input)

	globalAnwser := 0 

	for _, entry := range inputs {
		monceau := make([]int, 0)
		monceau = append(monceau, entry[len(entry)-2])
		print(monceau)
	}

	return globalAnwser
}

func Part2(input string) int {
	return 0
}

