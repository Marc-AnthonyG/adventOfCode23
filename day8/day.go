package day8

import (
	"fmt"
	"regexp"
	"strings"
)

func Run(input string) {
	fmt.Println("Day 8")
	fmt.Println("Part 1", Part1(input))
	fmt.Println("Part 2", Part2(input))
}

type Tuple struct {
	right string
	left  string
}

func Part1(input string) int {
	direction, points := parseInput(input)

	currentPoint := "AAA"
	numberOfSteps := 0
	for currentPoint != "ZZZ" {
		currentDiction := direction[numberOfSteps%len(direction)]
		if currentDiction == 'R' {
			currentPoint = points[currentPoint].right
		} else {
			currentPoint = points[currentPoint].left
		}
		numberOfSteps++
	}

	return numberOfSteps
}

func parseInput(input string) (string, map[string]Tuple) {
	lines := strings.Split(input, "\n")
	direction := lines[0]

	points := make(map[string]Tuple)
	pattern := `(\w{3}) = \((\w{3}), (\w{3})\)`
	re := regexp.MustCompile(pattern)

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		points[matches[1]] = Tuple{right: matches[3], left: matches[2]}
	}

	return direction, points
}

func Part2(input string) int {
	direction, points := parseInput(input)

	startingPoints := make([]string, 0)

	for key := range points {
		if key[2] == 'A' {
			startingPoints = append(startingPoints, key)
		}
	}

	numberOfSteps := 0
	solutionFound := false

	for !solutionFound {
		solutionFound = true
		for index := range startingPoints {
			currentDiction := direction[numberOfSteps%len(direction)]
			if currentDiction == 'R' {
				nextPoint := points[startingPoints[index]].right
				solutionFound = solutionFound && nextPoint[2] == 'Z'
				startingPoints[index] = nextPoint
			} else {
				nextPoint := points[startingPoints[index]].left
				solutionFound = solutionFound && nextPoint[2] == 'Z'
				startingPoints[index] = nextPoint
			}
		}
		numberOfSteps++
		fmt.Println(startingPoints)
	}

	return numberOfSteps
}
