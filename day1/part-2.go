package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	total := 2020

	sort.Ints(lines)

	for i := 0; i < len(lines)-2; i++ {
		if i == 0 || (i > 0) && lines[i] != lines[i+1] {
			low := i + 1
			high := len(lines) - 1
			rest := total - lines[i]
			for low < high {
				sum := lines[low] + lines[high]
				// found
				if sum == rest {
					fmt.Println(lines[low] * lines[high] * lines[i])
					return
				} else if sum > rest {
					high--
				} else {
					low++
				}
			}
		}
	}
}
