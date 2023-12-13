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

	for _, line := range lines {
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
		monceau := make([][]int, 1)
		monceau[0] = append(monceau[0], entry[len(entry)-1])
		
		currentLevel := 1
		foundOnlyZero := false
		for !foundOnlyZero {
			for i := 0; i <= currentLevel; i++ {
				if i == 0 {
					monceau[0] = append(monceau[0], entry[len(entry)-1-currentLevel])
				} else if i < currentLevel {
					newNumber := monceau[i-1][len(monceau[i-1])-2] - monceau[i-1][len(monceau[i-1])-1]
					monceau[i] = append(monceau[i], newNumber)
				} else {
					newNumber := monceau[i-1][len(monceau[i-1])-2] - monceau[i-1][len(monceau[i-1])-1]
					foundOnlyZero = newNumber == 0
					monceau = append(monceau, []int{newNumber})
				}
			}
			currentLevel++
		}

		for i := currentLevel-1; i >= 0; i-- {
			if i == 0 {
				globalAnwser += monceau[0][0]
			} else {
				extropaltedNumber := monceau[i][0] + monceau[i-1][0]
				monceau[i-1] = append([]int{extropaltedNumber}, monceau[i-1]...)
			}
		}
	}

	return globalAnwser
}

func Part2(input string) int {
	inputs := parseInput(input)

	globalAnwser := 0 

	for _, entry := range inputs {
		monceau := make([][]int, 1)
		monceau[0] = append(monceau[0], entry[0])
		
		currentLevel := 1
		foundOnlyZeroOnce := false
		foundOnlyZeroTwice := false
		for !foundOnlyZeroOnce || !foundOnlyZeroTwice {
			for i := 0; i <= currentLevel; i++ {
				if i == 0 {
					monceau[0] = append(monceau[0], entry[currentLevel])
				} else if i < currentLevel {
					newNumber := monceau[i-1][len(monceau[i-1])-1] - monceau[i-1][len(monceau[i-1])-2]
					monceau[i] = append(monceau[i], newNumber)
				} else {
					newNumber := monceau[i-1][1] - monceau[i-1][0]
					foundOnlyZeroTwice = newNumber == 0 && foundOnlyZeroOnce
					foundOnlyZeroOnce = newNumber == 0 || foundOnlyZeroOnce
					monceau = append(monceau, []int{newNumber})
				}
			}
			currentLevel++
		}

		fmt.Println(monceau)
		for i := currentLevel-1; i >= 0; i-- {
			if i == 0 {
				globalAnwser += monceau[0][0]
			} else {
				extropaltedNumber := monceau[i][0] - monceau[i-1][0]
				monceau[i-1] = append([]int{-extropaltedNumber}, monceau[i-1]...)
			}
		}
	}

	return globalAnwser
}

