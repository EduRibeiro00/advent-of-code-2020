package main

import (
	"bufio"
	"fmt"
	"os"
)

func binarySearch(str string, ceil int, leftChar rune, rightChar rune) int {
	low := 0
	high := ceil

	for _, elem := range str {
		mid := low + (high-low)/2
		if elem == leftChar {
			high = mid
		} else if elem == rightChar {
			low = mid + 1
		}
	}

	return low
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	highest := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := binarySearch(line[:7], 127, 'F', 'B')
		column := binarySearch(line[7:], 7, 'L', 'R')
		id := row*8 + column
		if highest < id {
			highest = id
		}
	}

	fmt.Println(highest)
}
