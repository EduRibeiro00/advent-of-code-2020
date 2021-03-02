package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func countTreesInSlope(lines []string, slopeRight int, slopeDown int) int {
	trees := 0

	linePos := 0
	for idx, line := range lines {
		if idx%slopeDown != 0 {
			continue
		}

		if line[linePos] == '#' {
			trees++
		}
		linePos = (linePos + slopeRight) % len(line)
	}

	return trees
}

func main() {
	lines := readFromFile("input.txt")

	totalTrees := countTreesInSlope(lines, 1, 1)
	totalTrees *= countTreesInSlope(lines, 3, 1)
	totalTrees *= countTreesInSlope(lines, 5, 1)
	totalTrees *= countTreesInSlope(lines, 7, 1)
	totalTrees *= countTreesInSlope(lines, 1, 2)

	fmt.Println("Number of trees: ", totalTrees)
}
