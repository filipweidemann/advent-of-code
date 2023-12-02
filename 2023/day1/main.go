package main

import (
	_ "embed"
	"fmt"
	"slices"
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

//go:embed input-test-two.txt
var testInputTwo string
var testInputTwoLines []string

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func init() {
	inputLines = strings.Split(input, "\n")
	testInputLines = strings.Split(testInput, "\n")
	testInputTwoLines = strings.Split(testInputTwo, "\n")
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

func (lc *LineCode) FromLine(l string, dm map[string]string) LineCode {
	runes := make([]string, 0)
	words := make([]string, len(dm))
	lineLen := len(l)

	i := 0
	for w := range dm {
		words[i] = w
		i++
	}

	for i, rune := range l {
		remainingChars := lineLen - i

		if !unicode.IsDigit(rune) {
			if len(dm) == 0 {
				continue
			}

			if remainingChars > 2 && slices.Contains(words, l[i:i+3]) {
				runes = append(runes, dm[l[i:i+3]])
			}
			if remainingChars > 3 && slices.Contains(words, l[i:i+4]) {
				runes = append(runes, dm[l[i:i+4]])
			}
			if remainingChars > 4 && slices.Contains(words, l[i:i+5]) {
				runes = append(runes, dm[l[i:i+5]])
			}
		} else {
			runes = append(runes, string(rune))
		}
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

	part2 := Part2(inputLines)
	fmt.Printf("Day 1, Part 2: %v", part2)
}

func Part1(i []string) int {
	pairs := make([]LineCode, 0)

	for _, line := range i {
		lc := LineCode{}
		dm := make(map[string]string, 0)
		pairs = append(pairs, lc.FromLine(line, dm))
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

func Part2(i []string) int {
	pairs := make([]LineCode, 0)
	for _, line := range i {
		lc := LineCode{}
		pairs = append(pairs, lc.FromLine(line, digitMap))
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
