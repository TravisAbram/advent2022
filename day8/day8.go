package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
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

	var tallest int
	
	// N E S W
	for i := 1; i < len(fileLines)-1; i++ {
		for j := 1; j < len(fileLines[i])-1; j++ {
			// value, _ := strconv.Atoi(string(fileLines[i][j]))
			// fmt.Println(value)
			var taller int
			for n := 0; n < i; n++ {
				if fileLines[n][j] >= fileLines[i][j] {
					// fmt.Println("North")
					taller++
					break
				}
			}
			for e := j+1; e < len(fileLines[i]); e++ {
				if fileLines[i][e] >= fileLines[i][j] {
					// fmt.Println("East")
					taller++
					break
				}
			}
			for s := i+1; s < len(fileLines); s++ {
				if fileLines[s][j] >= fileLines[i][j] {
					// fmt.Println("South")
					taller++
					break
				}
			}
			for w := 0; w < j; w++ {
				if fileLines[i][w] >= fileLines[i][j] {
					// fmt.Println("West")
					taller++
					break
				}
			}
			
			if taller != 4 {
				tallest++
			}
		}
	}
	
	fmt.Println("Part 1:", tallest + (2*len(fileLines)) + (2*(len(fileLines[0])-2)))
}
