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
	bid  int
}

func parseInput(input string) []Hand {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
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
	result := make([][]Hand, 7)

	for _, hand := range hands {
		test := getCardFormated(hand.card)
		isPair := false
		isThree := false
		isFour := false
		isFive := false
		hasTwoPair := false

		for _, value := range test {
			if value == 2 {
				hasTwoPair = isPair
				isPair = true
			} else if value == 3 {
				isThree = true
			} else if value == 4 {
				isFour = true
			} else if value == 5 {
				isFive = true
			}
		}

		if isFive {
			result[0] = appendInOrder(result[0], hand)
		} else if isFour {
			result[1] = appendInOrder(result[1], hand)
		} else if isThree && isPair {
			result[2] = appendInOrder(result[2], hand)
		} else if isThree {
			result[3] = appendInOrder(result[3], hand)
		} else if isPair && hasTwoPair {
			result[4] = appendInOrder(result[4], hand)
		}else if isPair {
			result[5] = appendInOrder(result[5], hand)
		} else {
			result[6] = appendInOrder(result[6], hand)
		}
	}

	globalResult := 0
	conteur := len(hands)
	fmt.Println(conteur)

	for _, value := range result {
		fmt.Println(value)
		for _, hand := range value {
			globalResult += hand.bid * conteur
			conteur--
		}
	}

	return globalResult
}

func appendInOrder(hands []Hand, hand Hand) []Hand {
	min := 0
	max := len(hands)

	for min < max {
		mid := (min + max) / 2

		for i := 0; i < len(hand.card); i++ {
			if hands[mid].card[i] == hand.card[i] {
				continue
			} else if compare(rune(hand.card[i]), rune(hands[mid].card[i])) {
				max = mid
				break
			} else {
				min = mid + 1
				break
			}
		}
	}
	
	return insert(hands, min, hand)
}

func insert(a []Hand, index int, value Hand) []Hand {
    if len(a) == index { // nil or empty slice or after last element
        return append(a, value)
    }
    a = append(a[:index+1], a[index:]...) // index < len(a)
    a[index] = value
    return a
}

func compare(r1 rune, r2 rune) bool {
	if r1 == 'A' {
		return true
	} else if r1 == 'K' {
		return r2 != 'A'
	} else if r1 == 'Q' {
		return r2 != 'A' && r2 != 'K'
	} else if r1 == 'J' {
		return r2 != 'A' && r2 != 'K' && r2 != 'Q'
	} else if r1 == 'T' {
		return r2 != 'A' && r2 != 'K' && r2 != 'Q' && r2 != 'J'
	}
	

	return r1 > r2
}

func getCardFormated(cards string) map[rune]int {
	result := map[rune]int{}

	for _, card := range cards {
		if _, ok := result[card]; ok {
			result[card]++
		} else {
			result[card] = 1
		}
	}

	return result
}
