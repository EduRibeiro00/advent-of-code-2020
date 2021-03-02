package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Empty struct{}

func unpackString(str string) (field string, value string) {
	arr := strings.Split(str, ":")
	return arr[0], arr[1]
}

func checkInRange(value string, low int, high int, requiredLen int) bool {
	if len(value) != requiredLen {
		return false
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return i >= low && i <= high
}

func checkValidField(field string, value string) bool {
	switch field {
	case "byr":
		return checkInRange(value, 1920, 2002, 4)
	case "iyr":
		return checkInRange(value, 2010, 2020, 4)
	case "eyr":
		return checkInRange(value, 2020, 2030, 4)
	case "hgt":
		if match, err := regexp.MatchString(`^([0-9]{2}in)|([0-9]{3}cm)$`, value); match && err == nil {
			unit := value[len(value)-2:]
			if unit == "in" {
				return checkInRange(value[:len(value)-2], 59, 76, 2)
			}
			return checkInRange(value[:len(value)-2], 150, 193, 3)
		}
		return false
	case "hcl":
		if match, err := regexp.MatchString(`^#([0-9a-f]){6}$`, value); match && err == nil {
			return true
		}
		return false
	case "ecl":
		if match, err := regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, value); match && err == nil {
			return true
		}
		return false
	case "pid":
		if match, err := regexp.MatchString(`^[0-9]{9}$`, value); match && err == nil {
			return true
		}
		return false
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
			field, value := unpackString(passportField)

			// check if field is required
			if _, ok := requiredFields[field]; ok && checkValidField(field, value) {
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
