package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
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

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	lines = lines[1:] // skip first line, it is the devil

	var pwd string = "/"
	tree := map[string]int{}

	for _, line := range lines {
		switch line[:4] {
		case "$ cd":
			if line[5:] == ".." {
				dirSize := tree[pwd]
				pwd = path.Dir(pwd)
				tree[pwd] += dirSize
			} else {
				if pwd != "/" {
					pwd = pwd + "/" + line[5:]
				} else {
					pwd = "/" + line[5:]
				}
			}
		case "$ ls":
			continue
		default:
			split := strings.Split(line, " ")
			size, _ := strconv.Atoi(split[0])
			tree[pwd] += size
		}
	}
	
	var part1total int
	for _, v := range tree {
		if v <= 100000 {
			part1total += v
		}
	}
	fmt.Println("Part 1:", part1total)
	
	var remaining int = 70000000 - tree["/"]
	var needed int = 30000000 - remaining
	var answer int = 30000000
	for _, v := range tree {
		if v >= needed {
			if v < answer {
				answer = v
			}
		}
	}
	fmt.Println("Part 2:", answer)
}
