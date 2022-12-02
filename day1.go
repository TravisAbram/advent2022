package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
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

	totals := create_totals(fileLines)
	max, location := find_max(totals)
	fmt.Println(max, location)
	fmt.Println(sum_max_values(totals, 3))

}

func create_totals(fileLines []string)([]int){
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

func find_max(slice []int)(int, int){
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

func remove_element(slice []int, index int)[]int{
	return append(slice[:index], slice[index+1:]...)
}

func sum_max_values(slice []int, top int) int {
	grand_total := 0
	for i := 0; i < top; i++{
		max, location := find_max(slice)
		grand_total += max
		fmt.Println(max)
		slice = remove_element(slice, location)
	}
	return grand_total
}