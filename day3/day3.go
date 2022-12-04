package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	first, second := splitCompartments("abcdeb")
	match := findMatch(first, second)
	fmt.Println(match)
	fmt.Println(getCharacterScore(match))
}

func splitCompartments(rucksack string)(first, second string){
	first, second = rucksack[:len(rucksack)/2], rucksack[(len(rucksack)/2):]
	return
}

func findMatch(first, second string) rune {
	for _, x := range first {
		for _, y := range second {
			if x == y {
				return strconv.(x)
			}
		}
	}
	return "Not found"
}

func getCharacterScore(character string) int {
	score, _ := strconv.Atoi(character)
	return score
}