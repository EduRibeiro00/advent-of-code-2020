package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return lines
}

func main() {
	lines := readFromFile("input.txt")
	count := 0
	re := regexp.MustCompile("[0-9]+")

	for _, line := range lines {
		data := strings.Split(line, ":")
		if len(data) != 2 {
			fmt.Printf("Error extracting information from a line")
			return
		}

		values := re.FindAllString(data[0], 2)
		if len(values) != 2 {
			fmt.Printf("Error extracting the range of a variable")
			return
		}
		// create variables
		low, err := strconv.Atoi(values[0])
		check(err)
		high, err := strconv.Atoi(values[1])
		check(err)
		c := data[0][len(data[0])-1]
		password := data[1][1:]

		countOccur := 0
		for _, char := range password {
			if byte(char) == c {
				countOccur++
			}
		}

		if low <= countOccur && countOccur <= high {
			count++
		}
	}

	fmt.Printf("%v", count)
}
