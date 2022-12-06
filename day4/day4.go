package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Part 1:", sumOfContainments(fileLines))
	fmt.Println("Part 2:", sumOfOverlaps(fileLines))
}

func sumOfContainments(fileLines []string) int {
	var total int
	for _, pair := range fileLines {
		a, b, x, y := splitAssignments(pair)
		total += checkContainment(a, b, x, y)
	}
	return total
}

func sumOfOverlaps(fileLines []string) int {
	var total int
	for _, pair := range fileLines {
		a, b, x, y := splitAssignments(pair)
		total += checkOverlap(a, b, x, y)
	}
	return total
}


func splitAssignments(assignments string)(a, b, x, y int){
	var first, second []string
	pairs := strings.Split(assignments, ",")
	for i, value := range pairs {
		if i == 0 {
			first = strings.Split(value, "-")
		} else {
			second = strings.Split(value, "-")
		}
	}
	a, _ = strconv.Atoi(first[0])
	b, _ = strconv.Atoi(first[1])
	x, _ = strconv.Atoi(second[0])
	y, _ = strconv.Atoi(second[1])
	return
}

func checkContainment(a, b, x, y int) int {
	if a == x {
		return 1
	} else if a < x {
		if b >= y {
			return 1
		}
	} else {
		if b <= y {
			return 1
		}
	}
	return 0
}

func checkOverlap(a, b, x, y int) int {
	if a == x || b == y {
		return 1
	} else if a < x {
		if b >= x {
			return 1
		}
	} else {
		if y >= a {
			return 1
		}
	}
	return 0
}