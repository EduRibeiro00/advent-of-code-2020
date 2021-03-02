package main

import (
	"bufio"
	"fmt"
	"os"
)

func fillMapWithLine(line string, answers map[string]struct{}, curSum int) int {
	newSum := curSum
	for _, char := range line {
		if _, ok := answers[string(char)]; !ok {
			newSum++
		}
		answers[string(char)] = struct{}{}
	}

	return newSum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	sameGroupAnswer := make(map[string]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sameGroupAnswer = make(map[string]struct{})
			continue
		}

		sum = fillMapWithLine(line, sameGroupAnswer, sum)
	}

	fmt.Println(sum)
}
