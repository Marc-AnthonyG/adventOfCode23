package day

import (
	"adventOfCode2023/day"
	"testing"
)

var input string = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

func Test_givenDay3Part1_whenBasicInput_thenAnwserIs4361(t *testing.T) {
	// given
	expected := 4361
	actual := day.Day3Part1(input)

	if actual != expected {
		t.Errorf("TestDay3Part1() returned %d, expected %d", actual, expected)
	}
}

func Test_givenDay3Part2_whenBasicInput_thenAnwserIs467835(t *testing.T) {
	// given
	expected := 467835
	actual := day.Day3Part2(input)

	if actual != expected {
		t.Errorf("TestDay3Part2() returned %d, expected %d", actual, expected)
	}
}



