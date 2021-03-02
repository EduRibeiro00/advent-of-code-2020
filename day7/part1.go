package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename io.Reader) map[string][]string {
	scanner := bufio.NewScanner(filename)
	m := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " bags contain ")
		containerBag := words[0]
		restOfLine := words[1]
		if restOfLine == "no other bags." {
			continue
		}

		containedBags := strings.Split(restOfLine, ", ")
		for _, subLine := range containedBags {
			words := strings.SplitN(subLine, " ", 2)
			_, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}
			words = strings.Split(words[1], " bag")
			containedBag := words[0]

			if arr, ok := m[containedBag]; ok {
				m[containedBag] = append(arr, containerBag)
			} else {
				m[containedBag] = []string{containerBag}
			}
		}
	}

	return m
}

func getBagsThatCanHold(m map[string][]string, bag string) map[string]struct{} {
	bags := make(map[string]struct{})
	return getBagsThatCanHoldAux(m, bag, bags)
}

func getBagsThatCanHoldAux(m map[string][]string, bag string, bags map[string]struct{}) map[string]struct{} {
	if containerBags, ok := m[bag]; ok {
		for _, containerBag := range containerBags {
			if _, ok := bags[containerBag]; !ok {
				bags[containerBag] = struct{}{}
			}
			bags = getBagsThatCanHoldAux(m, containerBag, bags)
		}
	}

	return bags
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(len(getBagsThatCanHold(parseInput(file), "shiny gold")))
}
