package day

import (
	"fmt"
	"strconv"
	"strings"
)

type ListOfMap struct {
	mapOfThing []MapTTT
}

type MapTTT struct {
	toRange   int
	fromRange int
	rangeLen  int
}

func (mapT *MapTTT) mapFromValue(number int) int {
	diff := number - mapT.fromRange
	return mapT.toRange + diff
}

func (listOfMap *ListOfMap) append(mapT MapTTT) {
	listOfMap.mapOfThing = append(listOfMap.mapOfThing, mapT)
}

func newListOfMap() ListOfMap {
	return ListOfMap{make([]MapTTT, 0)}
}

func newMapTTT(argument []string) MapTTT {
	toRange, _ := strconv.Atoi(argument[0])
	fromRange, _ := strconv.Atoi(argument[1])
	rangeLen, _ := strconv.Atoi(argument[2])
	return MapTTT{
		toRange:   toRange,
		fromRange: fromRange,
		rangeLen:  rangeLen,
	}
}

func Day5(input string) {
	fmt.Println("Day 5")
	fmt.Println("Part 1", Day5Part1(input))
	fmt.Println("Part 2", Day5Part2(input))
}

func Day5Part2(input string) int {
	lines := strings.Split(input, "\n")
	seeds := strings.Fields(lines[0])[1:]

	mapOfMap := parseInput(input)

	seedsPair := make([][]int, 0)

	for i := 0; i < len(seeds)-1; i += 2 {
		startingSeed, _ := strconv.Atoi(seeds[i])
		ranged, _ := strconv.Atoi(seeds[i+1])

		pair := []int{startingSeed, ranged}
		seedsPair = append(seedsPair, pair)
	}

	bestLocation := -1
	for _, seedPair := range seedsPair {
		for curSeed := seedPair[0]; curSeed < seedPair[0] + seedPair[1]; curSeed ++ {
			seedLocation := tryToEvaluateLocation(int(curSeed), mapOfMap)
			if bestLocation == -1 || bestLocation >= seedLocation {
				bestLocation = seedLocation
			}
		}
	}

	return bestLocation
}

func Day5Part1(input string) int {
	lines := strings.Split(input, "\n")
	seeds := strings.Fields(lines[0])[1:]

	mapOfMap := parseInput(input)

	bestLocation := -1
	for _, seed := range seeds {
		seedNumber, _ := strconv.Atoi(seed)
		seedLocation := tryToEvaluateLocation(seedNumber, mapOfMap)
		if bestLocation == -1 || bestLocation >= seedLocation {
			bestLocation = seedLocation
		}
	}

	return bestLocation
}

func tryToEvaluateLocation(seed int, listOfMaps []ListOfMap) int {
	fromVal := seed

	for _, listOfMap := range listOfMaps {
		toVal := fromVal
		for _, mapTTT := range listOfMap.mapOfThing {
			if mapTTT.content(fromVal) {
				toVal = mapTTT.mapFromValue(fromVal)
				break
			} else if mapTTT.hasNoImpact(fromVal) {
				continue
			} else if mapTTT.toRange < fromVal {
				toVal += mapTTT.rangeLen
			} else {
				toVal = mapTTT.rangeLen
			}
		}
		fromVal = toVal
	}

	return fromVal
}

func (mapT *MapTTT) content(number int) bool {
	return number >= mapT.fromRange && number < mapT.fromRange+mapT.rangeLen
}

func (mapT MapTTT) hasNoImpact(number int) bool {
	everyThingIsAbove := number < mapT.fromRange && number < mapT.toRange
	everyThingIsUnder := number > mapT.fromRange+mapT.rangeLen && number > mapT.toRange+mapT.rangeLen

	return everyThingIsAbove || everyThingIsUnder
}

func parseInput(input string) []ListOfMap {
	paragraphs := strings.Split(input, "\n\n")[1:]
	listOfMaps := make([]ListOfMap, 0)

	for _, paragraph := range paragraphs {
		if strings.Trim(paragraph, " ") == "" {
			continue
		}
		mapStrs := strings.Split(paragraph, "\n")[1:]
		mapForParagraph := newListOfMap()
		for _, mapStr := range mapStrs {
			if strings.Trim(mapStr, " ") == "" {
				continue
			}
			mapForParagraph.append(newMapTTT(strings.Fields(mapStr)))
		}
		listOfMaps = append(listOfMaps, mapForParagraph)
	}

	return listOfMaps
}
