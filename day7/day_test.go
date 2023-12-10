package day7

import (
	"testing"
)

var input = "32T3K 765\nAJ6J6 1\nQQ66A 1\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

func Test_Day7Part1(t *testing.T) {
	// given
	expected := 6440
	actual := Part1(input)

	if actual != expected {
		t.Errorf("Part one was supposed to returned %d, got %d", expected, actual)
	}
}


