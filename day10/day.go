package day10

import (
	"fmt"
	"strings"
)

func Run(input string) {
	fmt.Println("Day 10")
	fmt.Println("Part 1", Part1(input))
	fmt.Println("Part 2", Part2(input))
}

func parseInput(input string) ([]string, int, int) {
	lines := strings.Split(input, "\n")
	list := make([]string, 0)
	indexOfSLine := -1
	indexOfSCol := -1

	for lineIndex, line := range lines {
		if line == "" {
			continue
		}
		list = append(list, line)

		indexOfS := strings.Index(line, "S")
		if indexOfS != -1 {
			indexOfSLine = lineIndex
			indexOfSCol = indexOfS
		}
	}

	return list, indexOfSLine, indexOfSCol
}

var directions = []string{"F"}

func Part1(input string) int {
	labyrinth, lineIndex, indexOfSCol := parseInput(input)


	for _, direction := range directions {
		isLoop, steps := checkIfLoop(labyrinth, direction, lineIndex, indexOfSCol)
		if isLoop {
			if steps%2 == 0 {
				return steps/2
			} else {
				return steps/2 + 1
			}
		}
	}

	return -1
}

func checkIfLoop(labyrinth []string, startDir string, startLine int, startCol int) (bool, int) {
	steps := 0
	curCol := startCol
	curLine := -1
	lastCol := startCol
	lastLine := startLine

	if startDir == "|" || startDir == "L" || startDir == "J" {
		curLine = startLine - 1
	} else if startDir == "-"{
		curCol = startCol + 1
	} else {
		curLine = startLine + 1
	}

	labyrinth[startLine] = labyrinth[startLine][:startCol] + startDir + labyrinth[startLine][startCol+1:]
	for _, line := range labyrinth {
		fmt.Println(line)
	}

	for true {
		steps += 1
		tempCol := curCol
		tempLine := curLine

		if curLine < 0 || curLine >= len(labyrinth) || curCol < 0 || curCol >= len(labyrinth[curLine]) {
			return false, -1
		}

		curSymbol := string(labyrinth[curLine][curCol])
		fmt.Println(curSymbol)

		if curSymbol == "." {
			return false, 0
		}

		if curSymbol == "|" {
			if lastCol != curCol {
				return false, -1
			}

			if lastLine < curLine {
				curLine += 1
			} else {
				curLine -= 1
			}
		}

		if curSymbol == "-" {
			if lastLine != curLine {
				return false, -1
			}

			if lastCol < curCol {
				curCol += 1
			} else {
				curCol -= 1
			}
		}

		if curSymbol == "L" {
			if lastLine == curLine+1 || lastCol == curCol-1 {
				return false, -1
			}

			if lastCol == curCol {
				curCol += 1
			} else {
				curLine -= 1
			}
		}

		if curSymbol == "J" {
			if lastLine == curLine+1 || lastCol == curCol+1 {
				return false, -1
			}

			if lastCol == curCol {
				curCol -= 1
			} else {
				curLine -= 1
			}
		}

		if curSymbol == "7" {
			if lastLine == curLine-1 || lastCol == curCol+1 {
				return false, -1
			}

			if lastCol == curCol {
				curCol -= 1
			} else {
				curLine += 1
			}
		}

		if curSymbol == "F" {
			if lastLine == curLine-1 || lastCol == curCol-1 {
				return false, -1
			}

			if lastCol == curCol {
				curCol += 1
			} else {
				curLine += 1
			}
		}

		if curCol == startCol && curLine == startLine {
			return true, steps
		}

		lastCol = tempCol
		lastLine = tempLine
	}

	return false, -1
}

func Part2(input string) int {
	return 0
}
