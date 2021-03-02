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

func main() {
	const slopeRight = 3
	// const slopeDown = 1

	lines := readFromFile("input.txt")
	trees := 0

	linePos := 0
	for _, line := range lines {
		if line[linePos] == '#' {
			trees++
		}
		linePos = (linePos + slopeRight) % len(line)
	}

	fmt.Println("Number of trees: ", trees)
}
