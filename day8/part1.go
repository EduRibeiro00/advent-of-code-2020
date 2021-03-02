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

func main() {
	instructions := parseInstructions("input.txt")
	executed := make(map[int]struct{}) // indexes of already executed instructions
	acc := 0
	curInstIdx := 0 // current instruction index

	for curInstIdx < len(instructions) {
		inst := instructions[curInstIdx]

		if _, ok := executed[curInstIdx]; ok {
			fmt.Println(acc)
			return
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
}
