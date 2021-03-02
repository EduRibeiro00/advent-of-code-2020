package main

import (
	"bufio"
	"fmt"
	"os"
)

func fillMapWithLine(line string, answers map[string]int) {
	for _, char := range line {
		val, ok := answers[string(char)]
		if ok {
			answers[string(char)] = val + 1
		} else {
			answers[string(char)] = 1
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	peopleInGroup := 0
	sameGroupAnswer := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for _, v := range sameGroupAnswer {
				if v == peopleInGroup {
					sum++
				}
			}

			sameGroupAnswer = make(map[string]int)
			peopleInGroup = 0
			continue
		}

		peopleInGroup++
		fillMapWithLine(line, sameGroupAnswer)
	}

	for _, v := range sameGroupAnswer {
		if v == peopleInGroup {
			sum++
		}
	}

	fmt.Println(sum)
}
