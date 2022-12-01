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

	group_elves(fileLines)

}

func group_elves(fileLines []string){
	for _, line := range fileLines {
		if len(line) == 0 {
			fmt.Println("blank")
		} else {
			intLine, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(intLine)
		}
	}
}