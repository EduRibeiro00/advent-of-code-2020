package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func twoSum(n int, numbers []int) bool {
	for i, elem1 := range numbers {
		for j, elem2 := range numbers {
			if i == j {
				continue
			}
			if elem1+elem2 == n {
				return true
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	N := 25

	cnt := 0
	availablePrevious := []int{}

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if cnt < N {
			availablePrevious = append(availablePrevious, i)
		} else {
			if twoSum(i, availablePrevious) {
				availablePrevious = append(availablePrevious, i)
				availablePrevious = availablePrevious[1:]
			} else {
				fmt.Println(i)
				return
			}
		}
		cnt++
	}
}
