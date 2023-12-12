package day9

import (
	"testing"
)

func Test_Day9Part1(t *testing.T) {
	// given
	var input = "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)\n"
	var input2 = "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)\n"
	actual := Part1(input)
	actual2 := Part1(input2)

	if actual != 2{
		t.Errorf("Part one was supposed to returned %d, got %d", 2, actual)
	}
	if actual2 != 6{
		t.Errorf("Part one was supposed to returned %d, got %d", 6, actual)
	}
}

func Test_Day9Part2(t *testing.T) {
	// given
	var input = "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)"
	actual := Part2(input)

	if actual != 6 {
		t.Errorf("Part two was supposed to returned %d, got %d", 6, actual)
	}
}

