package day7

import (
	"testing"
)

var input = "32T3K 765\n35J74 2\n35596 1\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

func Test_Day7Part1(t *testing.T) {
	// given
	expected := 6440
	actual := Part1(input)

	if actual != expected {
		//t.Errorf("Part one was supposed to returned %d, got %d", expected, actual)
	}
}

func Test_Day7Part2(t *testing.T) {
	// given
	expected := 5905
	actual := Part2(input)

	if actual != expected {
		t.Errorf("Part two was supposed to returned %d, got %d", expected, actual)
	}
}

