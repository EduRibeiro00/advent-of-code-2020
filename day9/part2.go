package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, e := range arr {
		if e < m {
			m = e
		}
	}
	return m
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, e := range arr {
		if e > m {
			m = e
		}
	}
	return m
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	TARGET := 15690279

	numbers := []int{}

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}

	var left, right, curSum int

	for right < len(numbers) {
		if curSum == TARGET && len(numbers[left:right]) > 1 {
			res := min(numbers[left:right]) + max(numbers[left:right])
			fmt.Println(res)
			return
		}

		if curSum < TARGET {
			right++
			curSum += numbers[right]
		}

		if curSum > TARGET {
			left++
			curSum -= numbers[left]
		}
	}

}
