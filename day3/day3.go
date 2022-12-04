package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	
	fmt.Println("Part 1:", sumPriorities(fileLines))
	fmt.Println("Part 2:", sumBadges(fileLines))

}

func sumPriorities(sacks []string) int {
	var total int
	for _, rucksack := range sacks {
		first, second := splitCompartments(rucksack)
		match := findMatch(first, second)
		total += getCharacterScore(match)
	}
	return total
}

func sumBadges (sacks []string) int {
	var total int
	for i := 0; i < len(sacks)/3; i++ {
		first, second, third := sacks[i*3], sacks[(i*3)+1], sacks[(i*3)+2]
		match := findMatch3(first, second, third)
		total += getCharacterScore(match)
	}
	return total
}

func splitCompartments(rucksack string) (first, second string) {
	first, second = rucksack[:len(rucksack)/2], rucksack[(len(rucksack)/2):]
	return
}

func findMatch(first, second string) string {
	for _, x := range first {
		for _, y := range second {
			if x == y {
				return string(x)
			}
		}
	}
	return "Not found"
}

func findMatch3(first, second, third string) string {
	for _, x := range first {
		for _, y := range second {
			if x == y {
				for _, z := range third {
					if x == z {
						return string(x)
					}
				}
			}
		}
	}
	return "Not found"
}

func getCharacterScore(character string) int {
	priority := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	return bytes.IndexAny(priority, character) + 1
}
