package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type bagEntry struct {
	bagName string
	quant   int
}

func parseInput(filename io.Reader) map[string][]bagEntry {
	scanner := bufio.NewScanner(filename)
	m := make(map[string][]bagEntry)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " bags contain ")
		containerBag := words[0]
		restOfLine := words[1]
		if restOfLine == "no other bags." {
			continue
		}

		arr := []bagEntry{}

		containedBags := strings.Split(restOfLine, ", ")
		for _, subLine := range containedBags {
			words := strings.SplitN(subLine, " ", 2)
			value, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}
			words = strings.Split(words[1], " bag")
			containedBag := words[0]

			arr = append(arr, bagEntry{bagName: containedBag, quant: value})
		}

		m[containerBag] = arr
	}

	return m
}

func getBagsInside(m map[string][]bagEntry, bag string) int {
	sum := 0
	if containerBags, ok := m[bag]; ok {
		for _, containerBagEntry := range containerBags {
			containerBag := containerBagEntry.bagName
			value := containerBagEntry.quant
			sum = sum + value + value*getBagsInside(m, containerBag)
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(getBagsInside(parseInput(file), "shiny gold"))
}
