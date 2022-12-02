package main

import (
	"bufio"
	"fmt"
	"log"
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

	totals := createTotals(fileLines)
	max, location := findMax(totals)
	fmt.Println(max, location)
	fmt.Println(sumMaxValues(totals, 3))

}

func createTotals(fileLines []string) []int {
	var totals []int
	total := 0
	for _, line := range fileLines {
		if len(line) == 0 {
			totals = append(totals, total)
			total = 0
		} else {
			intLine, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += intLine
		}
	}
	// Collect the last total if the file doesn't end with a blank line
	totals = append(totals, total)
	return totals
}

func findMax(slice []int) (int, int) {
	max := 0
	location := 0
	for i, x := range slice {
		if x > max {
			max = x
			location = i
		}
	}
	return max, location
}

func removeElements(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func sumMaxValues(slice []int, top int) int {
	grand_total := 0
	for i := 0; i < top; i++ {
		max, location := findMax(slice)
		grand_total += max
		fmt.Println(max)
		slice = removeElements(slice, location)
	}
	return grand_total
}
