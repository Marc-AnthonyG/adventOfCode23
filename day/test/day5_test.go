package day

import (
	"adventOfCode2023/day"
	"testing"
)

var inputDay5 string = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"

func Test_givenDay5Part1_whenBasicInput_thenAnwserIs13(t *testing.T) {
	// given
	expected := 35
	actual := day.Day5Part1(inputDay5)

	if actual != expected {
		t.Errorf("TestDay5Part1() returned %d, expected %d", actual, expected)
	}
}

func Test_givenDay5Part2_whenBasicInput_thenAnwserIs30(t *testing.T) {
	expected := 46
	actual := day.Day5Part2(inputDay5)

	if actual != expected {
		t.Errorf("TestDay5Part2() returned %d, expected %d", actual, expected)
	}
}

