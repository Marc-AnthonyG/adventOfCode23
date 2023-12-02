package day

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() int {
	currentDir, _ := os.Getwd()
	file, err := os.Open(currentDir + "/input/input_day2.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}

	scanner := bufio.NewScanner(file)
	var globalSum int = 0

	for scanner.Scan() {
		currentLine := scanner.Text()
		gameIsPossible := true

		setOfDice := strings.Split(currentLine[7:], ";")

		for i := 0; i < len(setOfDice) && gameIsPossible; i++ {
			diceType := strings.Split(setOfDice[i], ",")
			for j := 0; j < len(diceType) && gameIsPossible; j++ {
				splitAgain := strings.Split(diceType[j], " ")
				if !checkIfRight(splitAgain[2], splitAgain[1]){
				    gameIsPossible = false
				}
			}
		}


		if gameIsPossible {
			gameIdStr:= strings.Split(currentLine, " ")[1]
			gameId,_ := strconv.Atoi(gameIdStr[:len(gameIdStr)-1])
			globalSum += gameId
		}
	}

	return globalSum
}

func Day2_2(input string) int {
	lines := strings.Split(input, "\n")
	var globalSum int = 0

	for  _,line := range lines[:len(lines)-1] {
		setOfDice := strings.Split(line[7:], ";")
		colorsMinimun := make(map[string]int)

		for i := 0; i < len(setOfDice); i++ {
			diceType := strings.Split(setOfDice[i], ",")
			for j := 0; j < len(diceType); j++ {
				splitAgain := strings.Split(diceType[j], " ")
				key := splitAgain[2]
				value,_ := strconv.Atoi(splitAgain[1])
				mapValue, keyExist := colorsMinimun[key]
				if !keyExist || mapValue < value {
					colorsMinimun[key] = value
				}
			}
		}

		power := 1
		for _,value := range colorsMinimun {
		    power *= value    
		}

		globalSum += power
	}

	return globalSum
}

func checkIfRight(color string, numberStr string) bool{
    number, _ := strconv.Atoi(numberStr)
    if color == "green" && number > 13 {
    	return false
    } else if color == "red" && number > 12 {
	return false
    } else if color == "blue" && number > 14 {
	return false
    }

    return true 
}

