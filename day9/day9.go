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

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	
	rope := make([]Segment, 2)
	
	for _, line := range fileLines {
		direction, count := parseInstructions(line)
		fmt.Println(direction, count, rope[0])

		for i :=0; i < count; i++ {
			rope[0] = moveSegment(rope[0], direction)
		}
	}
}

func moveSegment(s Segment, direction string)(Segment) {
	switch direction {
	case "U":
		s.up()
	case "D":
		s.down()
	case "L":
		s.left()
	case "R":
		s.right()
	}
	return s
}

func parseInstructions(line string)(direction string, count int) {
	split := strings.Split(line, " ")
	direction = split[0]
	count, _ = strconv.Atoi(split[1])
	return
}

type Segment struct {
	x int
	y int
}

type Motion interface {
	left()
	right()
	up()
	down()
}

func (s *Segment) left() {
	s.x--
}

func (s *Segment) right() {
	s.x++
}

func (s *Segment) up() {
	s.y--
}

func (s *Segment) down() {
	s.y++
}