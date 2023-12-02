package day

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Day1() int {
	currentDir, _ := os.Getwd()
	file, err := os.Open(currentDir + "/input/input_day1.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}

	scanner := bufio.NewScanner(file)
	var globalSum int = 0

	for scanner.Scan() {
		currentLine := scanner.Text()

		firstDegit, lastDegit := findFirstAndLastDegit(currentLine)
		stringNubmer := strconv.Itoa(firstDegit) + strconv.Itoa(lastDegit)
                current, errorParseInt := strconv.Atoi(stringNubmer)
		if errorParseInt != nil {
			return -1
		}

		globalSum += current
	}

	return globalSum
}

func findFirstAndLastDegit(string string) (int, int) {
	firstDegit := 0
	secondDigit := 0

	for i := 0; i < len(string); i++ {
		currentRune := rune(string[i])
		if unicode.IsDigit(currentRune) {
			if firstDegit == 0 {
				firstDegit = int(currentRune - '0')
			}
			secondDigit = int(currentRune - '0')
		} else {
			numberSpellOut := isNumberSpellOut(string, i)
			if numberSpellOut != -1 {
				if firstDegit == 0 {
					firstDegit = numberSpellOut
				}
				secondDigit = numberSpellOut
			}
		}
	}

	return firstDegit, secondDigit
}

func isNumberSpellOut(string string, index int) int {
	if haveXMoreChar(len(string), index, 3) {
		word := string[index : index+3]
		if word == "one" {
			return 1
		} else if word == "two" {
			return 2
		} else if word == "six" {
			return 6
		}
	}

	if haveXMoreChar(len(string), index, 4) {
		word := string[index : index+4]
		if word == "four" {
			return 4
		} else if word == "five" {
			return 5
		} else if word == "nine" {
			return 9
		}
	}

	if haveXMoreChar(len(string), index, 5) {
		word := string[index : index+5]
		if word == "three" {
			return 3
		} else if word == "seven" {
			return 7
		} else if word == "eight" {
			return 8
		}
	}

	return -1
}

func haveXMoreChar(lenght int, currentIndex int, lenghtToSpell int) bool {
	return lenght >= currentIndex+lenghtToSpell
}
