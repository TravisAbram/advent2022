package main

import (
	"bufio"
	"fmt"
	"math"
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
	
	rope := make([]Segment, 10)
	
	tail := map[Segment]int{}
	
	for _, line := range fileLines {
		direction, count := parseInstructions(line)

		for i :=0; i < count; i++ {
			// fmt.Println(direction, count, i, rope[0], rope[1])
			rope[0] = moveSegment(rope[0], direction)
			for j := 1; j < len(rope); j++ {
				rope[j] = follow(rope[j-1], rope[j])
			}
			tail[rope[len(rope)-1]] += 1
		}
	}
	fmt.Println("Part 2: ", len(tail))
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

func follow(head, tail Segment)(Segment){
	xabs := math.Abs(float64(head.x)-float64(tail.x))
	yabs := math.Abs(float64(head.y)-float64(tail.y))
	if xabs < 2 && yabs < 2 {
		return tail
	}

	x_motion := head.x - tail.x
	y_motion := head.y - tail.y
	
	if x_motion > 0 {
		tail.right()
	}
	if x_motion < 0 {
		tail.left()
	}
	if y_motion > 0 {
		tail.down()
	}
	if y_motion < 0 {
		tail.up()
	}
	return tail
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