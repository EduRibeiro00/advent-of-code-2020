package utils

import (
	"bufio"
	"os"
)

func readFromFile(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return lines
}
