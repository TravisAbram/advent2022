package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var fileLines []string

	for fileScanner.Scan(){
		fileLines = append(fileLines, fileScanner.Text())
	}
	fmt.Println("Part 1: ", calculateMatch(fileLines))
}

func calculateMatch(rounds []string) int {
	var total int
	for _, round := range rounds {
		total += calculateScore(round)
	}
	return total
}

func calculateScore(move string) int {
	var theirs, mine string = string(move[0]), string(move[2])
	result := winLoseDraw(theirs, mine)
	shape := shapeScore(mine)
	return result + shape
}

func winLoseDraw(theirs, mine string) int {
	switch round := theirs + mine; round {
	case "AX", "BY", "CZ": // all draws
		return 3
	case "AY", "BZ", "CX": // all wins
		return 6
	default:
		return 0
	}
}

func shapeScore(mine string) int {
	switch mine {
	case "X":
		return 1
	case "Y":
		return 2
	default:
		return 3
	}
}