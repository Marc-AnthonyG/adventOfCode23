package day10

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := ".....\n.S-7.\n.|.|.\n.L-J.\n....."
	got := Part1(input)
	if got != 4 {
		t.Errorf("Part1(input) = %d, want %d", got, 4)
	}
}

func TestPart1Harder(t *testing.T) {
	input := "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"
	got := Part1(input)
	if got != 8 {
		t.Errorf("Part1(input) = %d, want %d", got, 8)
	}
}

