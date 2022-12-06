package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var stackRows, instructions []string
	var line int

	for fileScanner.Scan() {
		line++
		if line <= 8 {
			stackRows = append(stackRows, fileScanner.Text())
		} else if line >= 11 {
			instructions = append(instructions, fileScanner.Text())
		}
	}
	
	stacks := buildStacks(stackRows)
	stacksP2 := buildStacks(stackRows)
	// [[N D M Q B P Z] [C L Z Q M D H V] [Q H R D V F Z G] [H G D F N] [N F Q] [D Q V Z F B T] [Q M T Z D V S H] [M G F P N Q] [B W R M]]
	for _, row := range instructions {
		move, from, to := parseInstructions(row)
		from--
		to--
		for i := 0; i < move; i++ {
			target := stacks[from][len(stacks[from])-1]
			stacks[to] = append(stacks[to], target)
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}
	
	fmt.Println("Part 1:", getTopValues(stacks))
	
	for _, row := range instructions {
		move, from, to := parseInstructions(row)
		from--
		to--
		stacksP2[to] = append(stacksP2[to], stacksP2[from][len(stacksP2[from])-move:]...)
		stacksP2[from] = stacksP2[from][:len(stacksP2[from])-move]
	}
	
	fmt.Println("Part 2:", getTopValues(stacksP2))
}

func getTopValues(stacks [][]string) []string {
	values := make([]string, len(stacks))
	for i := 0; i < len(stacks); i++ {
		values[i] = stacks[i][len(stacks[i])-1]
	}
	return values
}

func parseInstructions(instruction string)(move, from, to int){
	instructions := strings.Split(instruction, " ")
	move, _ = strconv.Atoi(instructions[1])
	from, _ = strconv.Atoi(instructions[3])
	to, _ = strconv.Atoi(instructions[5])
	return
}

func buildStacks(stackRows []string)([][]string) {
	stacks := make([][]string, len(stackRows)+1)
	
	for i := 0; i <= 8; i++ {
		var length int = 8
		var position int = 1
		for _, line := range stackRows {
			character := string(line[(i*4)+1])
			if character == " " {
				length--
			}
		}
		stacks[i] = make([]string, length)
		for _, line := range stackRows {
			character := string(line[(i*4)+1])
			if character == " " {
				continue
			}
			stacks[i][length-position] = character
			position++
		}
	}
	return stacks
}