package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	lines := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return lines
}

func main() {
	lines := readFromFile("input.txt")
	m := make(map[int]int)
	total := 2020

	for _, number := range lines {
		if val, ok := m[total-number]; ok {
			fmt.Println(number * val)
		}
		m[number] = number
	}
}
