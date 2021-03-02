package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inst struct {
	instType string
	offset   int
}

func parseInstructions(filename string) []Inst {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	insts := []Inst{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		instType := line[0]
		offset, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		insts = append(insts, Inst{instType: instType, offset: offset})
	}

	return insts
}

func checkTermination(instructions []Inst) bool {
	executed := make(map[int]struct{}) // indexes of already executed instructions
	acc := 0
	curInstIdx := 0 // current instruction index

	for curInstIdx < len(instructions) {
		inst := instructions[curInstIdx]

		if _, ok := executed[curInstIdx]; ok {
			return false
		}
		executed[curInstIdx] = struct{}{}

		switch inst.instType {
		case "acc":
			acc = acc + inst.offset
			curInstIdx++
		case "jmp":
			curInstIdx = curInstIdx + inst.offset
		default: // nop
			curInstIdx++
		}
	}

	fmt.Println(acc)
	return true
}

func main() {
	instructions := parseInstructions("input.txt")
	for idx, _ := range instructions {
		inst := &instructions[idx]
		switch inst.instType {
		case "jmp":
			inst.instType = "nop"
			if checkTermination(instructions) {
				return
			}
			inst.instType = "jmp"
		case "nop":
			inst.instType = "jmp"
			if checkTermination(instructions) {
				return
			}
			inst.instType = "nop"
		}
	}
}
