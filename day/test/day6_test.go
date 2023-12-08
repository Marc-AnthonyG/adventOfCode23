package day

import (
	"adventOfCode2023/day"
	"testing"
)

var inputDay6 string = "Time:      7  15   30\nDistance:  9  40  200"

func Test_Day6Part1(t *testing.T) {
	// given
	expected := 288
	actual := day.Day6Part1(inputDay6)

	if actual != expected {
		t.Errorf("TestDay6Part1() returned %d, expected %d", actual, expected)
	}
}

func Test_Day6Part2(t *testing.T) {
	expected := 71503
	actual := day.Day6Part2(inputDay6)

	if actual != expected {
		t.Errorf("TestDay6Part2() returned %d, expected %d", actual, expected)
	}
}

