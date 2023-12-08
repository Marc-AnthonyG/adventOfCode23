package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	time     int
	distance int
}

func Day6(input string) {
	fmt.Println("Day 6")
	fmt.Println("Part 1", Day6Part1(input))
	fmt.Println("Part 2", Day6Part2(input))
}

func parseInputDay6(input string) []Game {
	lines := strings.Split(input, "\n")
	timeStrs := strings.Fields(strings.Split(lines[0], ":")[1])
	distanceStrs := strings.Fields(strings.Split(lines[1], ":")[1])

	games := make([]Game, 0)
	for i := range timeStrs {
		time, error := strconv.Atoi(timeStrs[i])
		must(error, "time must be integer")
		distance, error := strconv.Atoi(distanceStrs[i])
		must(error, "time must be integer")
		games = append(games, Game{time, distance})
	}

	return games
}

func Day6Part2(input string) int {
	lines := strings.Split(input, "\n")
	timeStrs := strings.Split(lines[0], ":")[1]
	distanceStrs := strings.Split(lines[1], ":")[1]
	
	distance, _ := strconv.Atoi(strings.ReplaceAll(distanceStrs, " ", ""))
	time, _ := strconv.Atoi(strings.ReplaceAll(timeStrs, " ", ""))


	minimun := binaryMinimunApply(distance, time)
	return time - (2 * minimun) + 1
}

func Day6Part1(input string) int {
	games := parseInputDay6(input)

	total := 1

	for _, game := range games {
		minimun := binaryMinimunApply(game.distance, game.time)
		total *= game.time - (2 * minimun) + 1
	}

	return total
}

func binaryMinimunApply(distance int, time int) int {
	min := 0 
	max := time/2

	for min < max {
		mid := (min + max) / 2
		expressionResult := (time - mid) * mid

		if expressionResult <= distance {
			min = mid + 1
		} else {
			max = mid
		}
	}

	return min
}

func must(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %v", msg, err)
	}
}
