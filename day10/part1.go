package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

	var num1Diff, num3Diff, prev int

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	for _, adapter := range adapters {
		diff := adapter - prev
		if diff == 1 {
			num1Diff++
		}
		if diff == 3 {
			num3Diff++
		}
		prev = adapter
	}

	res := num1Diff * num3Diff
	fmt.Println(res)
}
