package day9

import (
	"testing"
)

func Test_Day9Part1(t *testing.T) {
	var input = "0 3 6 9 12 15"

	actual := Part1(input)

	if actual != 18 {
		t.Errorf("Part one was supposed to returned %d, got %d", 18, actual)
	}
}

func Test_Day9Part1largerInput(t *testing.T) {
	var input2 = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	actual2 := Part1(input2)

	if actual2 != 114 {
		t.Errorf("Part one was supposed to returned %d, got %d", 114, actual2)
	}
}


func Test_Day9Part2(t *testing.T) {
	var input = "0 3 6 9 12 15"

	actual := Part2(input)

	if actual != -3 {
		t.Errorf("Part one was supposed to returned %d, got %d", -3, actual)
	}
}

func Test_Day9Part2Diff(t *testing.T) {
	var input = "10 13 16  21 30 45"

	actual := Part2(input)

	if actual != 5 {
		t.Errorf("Part one was supposed to returned %d, got %d", 5, actual)
	}
}
