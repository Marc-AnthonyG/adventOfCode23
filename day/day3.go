package day

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day3(input string) {
	fmt.Println("Day 3")
	fmt.Println("Part 1", Day3Part1(input))
	fmt.Println("Part 2", Day3Part2(input))
}

func Day3Part2(input string) int {
	lines := strings.Split(input, "\n")

	globalGearRatio := 0

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		for charIndex := 0; charIndex < len(line); charIndex++ {
			if string(line[charIndex]) == "*" {
				number1, number2 := getNumbersNearby(lines, charIndex, lineIndex)
				globalGearRatio += number1 * number2
			}
		}
	}

	return globalGearRatio
}

func getNumbersNearby(lines []string, charIndex int, lineIndex int) (int, int) {
	number1 := 0
	number2 := 0

	for i := lineIndex - 1; i <= lineIndex+1; i++ {
		if i >= 0 && i < len(lines) {
			line := lines[i]
			for j := charIndex - 1; j <= charIndex+1; j++ {
				if j >= 0 && j < len(line) {
					if unicode.IsDigit(rune(line[j])) {
						newNumber, lastIndex := getNumberAtIndex(line, j)
						if number1 == 0 {
							number1 = newNumber
						} else if number2 == 0 {
							number2 = newNumber
						} else {
							return 0, 0
						}
						j = lastIndex
					}
				}
			}
		}
	}

	return number1, number2
}

func getNumberAtIndex(line string, index int) (int,int) {
	beginingIndex := index
	lastIndex := index+1
	for beginingIndex > 0 && unicode.IsDigit(rune(line[beginingIndex-1])) {
		beginingIndex--
	}
	for lastIndex < len(line) && unicode.IsDigit(rune(line[lastIndex])) {
		lastIndex++
	}
	number, _ := strconv.Atoi(line[beginingIndex:lastIndex])
	return number, lastIndex
}



func Day3Part1(input string) int {
	lines := strings.Split(input, "\n")

	globalValue := 0

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		for charIndex := 0; charIndex < len(line); charIndex++ {
			if unicode.IsDigit(rune(line[charIndex])) {
				lastDegitIndex := getLastIndexOfNumber(line, charIndex)
				if hasASymbolNearby(lines, charIndex, lastDegitIndex, lineIndex) {
					number, _ := strconv.Atoi(line[charIndex:lastDegitIndex+1])
					globalValue += number
				}
				charIndex = lastDegitIndex
			}
		}
	}

	return globalValue
}

func hasASymbolNearby(lines []string, charIndex int, lastDegitIndex int, lineIndex int) bool {
	for i := lineIndex - 1; i <= lineIndex+1; i++ {
		if i >= 0 && i < len(lines) {
			line := lines[i]
			for j := charIndex - 1; j <= lastDegitIndex+1; j++ {
				if j >= 0 && j < len(line) {
					if string(line[j]) != "." && !unicode.IsDigit(rune(line[j])) {
						return true
					}
				}
			}
		}
	}

	return false
}

func getLastIndexOfNumber(line string, beginingIndex int) int {
	index := beginingIndex
	for ; index < len(line) && unicode.IsDigit(rune(line[index])); index++ {
	}
	return index-1
}
