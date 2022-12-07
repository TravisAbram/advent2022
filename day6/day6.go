package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(readFile)-4; i++ {
		set := map[byte]int{}
		for _, char := range readFile[i:i+4]{
			set[char] += 1
		}
		if len(set) == 4 {
			fmt.Println("Part 1:", i+4, set)
			goto PART2
		}
	}

	PART2:
	for i := 0; i < len(readFile)-14; i++ {
		set := map[byte]int{}
		for _, char := range readFile[i:i+14]{
			set[char] += 1
		}
		if len(set) == 14 {
			fmt.Println("Part 2:", i+14, set)
			return
		}
	}
}