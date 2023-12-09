package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(input string) {
	fmt.Println("Day 7")
	fmt.Println("Part 1", Part1(input))
	fmt.Println("Part 2", Part2(input))
}

type Hand struct {
	card string
	bid int
}

func parseInput(input string) []Hand{
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0)
	
	for _, line := range lines{
		handStr := strings.Fields(line)
		bid, _ := strconv.Atoi(handStr[1])

		hands = append(hands, Hand{handStr[0], bid})
	}
	
	return hands
}

func Part2(input string) int {
	return -1
}

func Part1(input string) int {
	hands := parseInput(input)

	for _, hand := range hands{
		fmt.Println(hand)
	}
	

	return -1
}

