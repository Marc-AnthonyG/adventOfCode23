package day

import (
	"adventOfCode2023/day"
	"testing"
)

var inputDay4 string = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"

func Test_givenDay4Part1_whenBasicInput_thenAnwserIs13(t *testing.T) {
	// given
	expected := 13
	actual := day.Day4Part1(inputDay4)

	if actual != expected {
		t.Errorf("TestDay3Part1() returned %d, expected %d", actual, expected)
	}
}

func Test_givenDay4Part2_whenBasicInput_thenAnwserIs30(t *testing.T) {
	// given
	expected := 30
	actual := day.Day4Part2(inputDay4)

	if actual != expected {
		t.Errorf("TestDay3Part2() returned %d, expected %d", actual, expected)
	}
}



