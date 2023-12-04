package day

import (
	"fmt"
	"strings"
)

type Card struct {
	accepting  map[string]bool
	totalCopie int
}

func NewCard() Card {
	return Card{
		accepting:  make(map[string]bool),
		totalCopie: 0,
	}
}

func (card *Card) getValue(cardValues []string) int {
	value := 0

	for _, valueInCard := range cardValues {
		_, ok := card.accepting[valueInCard]

		if ok {
			card.totalCopie = 1 + card.totalCopie
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}

	return value
}

func (card *Card) addAccepting(acceptedValue []string) {
	for _, value := range acceptedValue {
		card.accepting[value] = true
	}
}

func Day4(input string) {
	fmt.Println("Day 4")
	fmt.Println("Part 1", Day4Part1(input))
	fmt.Println("Part 2", Day4Part2(input))
}

func Day4Part2(input string) int {
	lines := strings.Split(input, "\n")
	cards := make([]Card, 0)

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		if strings.Trim(lines[lineIndex], " ") == "" {
			continue
		}
		onlyNumber := strings.Split(lines[lineIndex], ":")
		lineOfNumber := strings.Split(onlyNumber[1], "|")
		card := NewCard()
		card.addAccepting(strings.Fields(lineOfNumber[0]))
		card.getValue(strings.Fields(lineOfNumber[1]))
		cards = append(cards, card)
	}

	globalValue := 0

	for index := range cards {
		globalValue += exploreCard(cards, index)+ 1
	}

	return globalValue
}

func exploreCard(cards []Card, currentIndex int) int {
	value := cards[currentIndex].totalCopie

	for i:= 1; i <= cards[currentIndex].totalCopie; i++{
		if i+currentIndex >= len(cards) {
			continue
		} else {
			value += exploreCard(cards, i+currentIndex)
		}
	}

	return value
}

func Day4Part1(input string) int {
	lines := strings.Split(input, "\n")
	globalValue := 0

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		if strings.Trim(lines[lineIndex], " ") == "" {
			continue
		}
		onlyNumber := strings.Split(lines[lineIndex], ":")
		lineOfNumber := strings.Split(onlyNumber[1], "|")
		card := NewCard()
		card.addAccepting(strings.Fields(lineOfNumber[0]))

		globalValue += card.getValue(strings.Fields(lineOfNumber[1]))

	}

	return globalValue
}
