package day10

import (
	"fmt"
	"math"
	"strings"
)

func Run(input string) {
	fmt.Println("Day 10")
	fmt.Println("Part 1", Part1(input))
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

var directions = []string{"|", "-", "L", "J", "7", "F"}

func Part1(input string) int {
	labyrinth, lineIndex, indexOfSCol := parseInput(input)


	for _, direction := range directions {
		isLoop, steps, vectices := checkIfLoop(labyrinth, direction, lineIndex, indexOfSCol)
		fmt.Println(shoelaceArea(vectices))
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

func shoelaceArea(vertices [][]float64) float64 {
	n := len(vertices)
	total := 0.0

	for i := 0; i < n-1; i++ {
		total += (vertices[i][0] * vertices[i+1][1]) - (vertices[i+1][0] * vertices[i][1])
	}

	total += (vertices[n-1][0] * vertices[0][1]) - (vertices[0][0] * vertices[n-1][1])

	return 0.5 * math.Abs(total)
}

func checkIfLoop(labyrinth []string, startDir string, startLine int, startCol int) (bool, int, [][]float64) {
	steps := 0
	curCol := startCol
	curLine := -1
	lastCol := startCol
	lastLine := startLine
	vectices := make([][]float64, 0)
	
	if startDir == "|" || startDir == "L" || startDir == "J" {
		curLine = startLine - 1
	} else if startDir == "-"{
		curCol = startCol + 1
	} else {
		curLine = startLine + 1
	}

	labyrinth[startLine] = labyrinth[startLine][:startCol] + startDir + labyrinth[startLine][startCol+1:]

	for true {
		steps += 1
		tempCol := curCol
		tempLine := curLine

		if curLine < 0 || curLine >= len(labyrinth) || curCol < 0 || curCol >= len(labyrinth[curLine]) {
			return false, -1, nil
		}

		curSymbol := string(labyrinth[curLine][curCol])

		if curSymbol == "." {
			return false, 0, nil
		}

		if curSymbol == "|" {
			if lastCol != curCol {
				return false, -1, nil
			}

			if lastLine < curLine {
				curLine += 1
			} else {
				curLine -= 1
			}
		}

		if curSymbol == "-" {
			if lastLine != curLine {
				return false, -1, nil
			}

			if lastCol < curCol {
				curCol += 1
			} else {
				curCol -= 1
			}
		}

		if curSymbol == "L" {
			if lastLine == curLine+1 || lastCol == curCol-1 {
				return false, -1, nil
			}

			vectices = append(vectices, []float64{float64(lastCol), float64(lastLine)})

			if lastCol == curCol {
				curCol += 1
			} else {
				curLine -= 1
			}
		}

		if curSymbol == "J" {
			if lastLine == curLine+1 || lastCol == curCol+1 {
				return false, -1, nil
			}

			vectices = append(vectices, []float64{float64(lastCol), float64(lastLine)})
			if lastCol == curCol {
				curCol -= 1
			} else {
				curLine -= 1
			}
		}

		if curSymbol == "7" {
			if lastLine == curLine-1 || lastCol == curCol+1 {
				return false, -1, nil
			}

			vectices = append(vectices, []float64{float64(lastCol), float64(lastLine)})
			if lastCol == curCol {
				curCol -= 1
			} else {
				curLine += 1
			}
		}

		if curSymbol == "F" {
			if lastLine == curLine-1 || lastCol == curCol-1 {
				return false, -1, nil
			}

			vectices = append(vectices, []float64{float64(lastCol), float64(lastLine)})
			if lastCol == curCol {
				curCol += 1
			} else {
				curLine += 1
			}
		}

		if curCol == startCol && curLine == startLine {
			return true, steps, vectices
		}

		lastCol = tempCol
		lastLine = tempLine
	}

	return false, -1, nil
}
