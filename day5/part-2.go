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

func calcID(line string) int {
	row := binarySearch(line[:7], 127, 'F', 'B')
	column := binarySearch(line[7:], 7, 'L', 'R')
	return row*8 + column
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lowestID := 0
	highestID := 1023
	foundIDs := make(map[int]struct{})

	for scanner.Scan() {
		id := calcID(scanner.Text())
		foundIDs[id] = struct{}{}
	}

	for i := lowestID; i <= highestID; i++ {
		if _, ok := foundIDs[i]; !ok {
			_, okBefore := foundIDs[i-1]
			_, okAfter := foundIDs[i+1]
			if okBefore && okAfter {
				fmt.Println(i)
				return
			}
		}
	}

	fmt.Println("Did not find my seat")
}
