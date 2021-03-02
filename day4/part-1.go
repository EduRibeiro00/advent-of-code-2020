package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Empty struct{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	validPassports := 0
	requiredFields := map[string]struct{}{"byr": Empty, "iyr": Empty, "eyr": Empty, "hgt": Empty, "hcl": Empty, "ecl": Empty, "pid": Empty}
	scannedFields := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		// reached blank line; another passport incoming. Clear scannedFields
		if line == "" {
			if len(scannedFields) == len(requiredFields) {
				validPassports++
			}

			scannedFields = scannedFields[:0]
			continue
		}

		passportFields := strings.Fields(line)
		for _, passportField := range passportFields {
			field := strings.Split(passportField, ":")[0]

			// check if field is required
			if _, ok := requiredFields[field]; ok {
				scannedFields = append(scannedFields, field)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// for the last passport
	if len(scannedFields) == len(requiredFields) {
		validPassports++
	}

	fmt.Println("Valid passports: ", validPassports)
}
