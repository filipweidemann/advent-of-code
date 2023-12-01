package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string
var inputLines []string

//go:embed input-test.txt
var testInput string
var testInputLines []string

func init() {
	inputLines = strings.Split(input, "\n")
	testInputLines = strings.Split(testInput, "\n")
}

type LineCode struct {
	first  string
	second string
}

func (lc *LineCode) IsEmpty() bool {
	return lc.first == "" && lc.second == ""
}

func (lc *LineCode) IsFull() bool {
	return lc.first != "" && lc.second != ""
}

func (lc *LineCode) FromLine(l string) LineCode {
	runes := make([]rune, 0)

	for _, rune := range l {
		if !unicode.IsDigit(rune) {
			continue
		}

		runes = append(runes, rune)
	}

	if runeLength := len(runes); runeLength == 0 {
		lc.first = "0"
		lc.second = "0"
		return *lc
	} else if runeLength == 1 {
		val := string(runes[0])
		lc.first = val
		lc.second = val
	} else {
		lc.first = string(runes[0])
		lc.second = string(runes[runeLength-1])
	}

	return *lc
}

func main() {
	part1 := Part1(inputLines)
	fmt.Printf("Day 1, Part 1: %v", part1)
}

func Part1(i []string) int {
	pairs := make([]LineCode, 0)

	for _, line := range i {
		lc := LineCode{}
		pairs = append(pairs, lc.FromLine(line))
	}

	sum := 0
	for _, lc := range pairs {
		num, err := strconv.Atoi((lc.first + lc.second))
		if err != nil {
			panic("Invalid character in input string.")
		}

		sum += num
	}

	return sum
}
