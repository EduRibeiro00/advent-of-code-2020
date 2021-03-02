package main

import (
	"bufio"
	"os"
	"strconv"
)

func max(arr []int) int {
	max := 0
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	adapters := []int{}

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, i)
	}

	adapters = append(adapters, max(adapters)+3)

	// maybe DP, comecar com o valor do telemovel
}
