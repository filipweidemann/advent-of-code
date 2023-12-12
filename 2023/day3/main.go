package main

import (
	_ "embed"
	"fmt"
	"github.com/filipweidemann/advent-of-code/utils"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input-test.txt
var testInput string
var testInputLines []string

//go:embed input.txt
var input string
var inputLines []string

func init() {
	testInputLines = strings.Split(testInput, "\n")
	testInputLines = testInputLines[:len(testInputLines)-1]
	inputLines = strings.Split(input, "\n")
	inputLines = inputLines[:len(inputLines)-1]
}

type Symbol struct {
	LineID int
	Index  int
}

type Digit struct {
	LineID  int
	Indices []int
	Runes   []rune
	Value   int
}

func (d *Digit) HasAdjacentSymbol(l []Line) bool {
	// we assume equally long lines, so just grab one and get their length
	lineLength := len(l[0].Content)

	// vertical indices
	startLine := utils.MinWithBound(d.LineID-1, 0)
	endLine := utils.MaxWithBound(d.LineID+1, lineLength-1)

	// horizontal indices
	startIndex := utils.MinWithBound(d.Indices[0], 0)
	endIndex := utils.MaxWithBound(d.Indices[len(d.Indices)-1], lineLength)

	for i := startLine; i <= endLine; i++ {
		for j := startIndex; j <= endIndex; j++ {
			c := l[i].Content[j]
			if string(c) != "." && !unicode.IsDigit(rune(c)) {
				return true
			}
		}
	}

	return false
}

type Line struct {
	LineID  int
	Content string
}

func LinesFromSlice(s []string) []Line {
	lines := []Line{}

	for i, l := range s {
		lines = append(lines, Line{i, l})
	}

	return lines
}

func main() {
	sum := Part1(testInputLines)
	fmt.Printf("Part 1: %v \n", sum)

	power := Part2(inputLines)
	fmt.Printf("Part 2: %v", power)
}

func Part1(i []string) int {
	// var symbols []Symbol
	var digits []Digit

	lines := LinesFromSlice(i)
	_, digits = Collect(lines)

	sum := 0
	for _, d := range digits {
		s := d.HasAdjacentSymbol(lines)
		if s {
			sum += d.Value
		}
	}

	return sum
}

func Part2(i []string) uint64 {
	return 0
}

func Collect(l []Line) ([]Symbol, []Digit) {
	var symbols []Symbol
	var digits []Digit

	for lineIndex, line := range l {
		digit := Digit{LineID: lineIndex}

		for runeIndex, rune := range line.Content {
			if unicode.IsDigit(rune) {
				digit.Runes = append(digit.Runes, rune)
				digit.Indices = append(digit.Indices, runeIndex)
				continue
			}

			// Now, also check for symbols
			if rune != '.' {
				symbols = append(symbols, Symbol{lineIndex, runeIndex})
			}

			// If no prior digit is written, just go to next char
			if len(digit.Runes) == 0 {
				continue
			}

			// If prior digit was written, determine its value and append it
			fmt.Println("Converting runes to Str: ", string(digit.Runes))
			value, err := strconv.Atoi(string(digit.Runes))
			if err != nil {
				panic("Could not parse runes to digit")
			}

			digit.Value += value
			fmt.Println("Appending Digit: ", digit)
			digits = append(digits, digit)
			digit.Runes = nil
			digit.Indices = nil

		}
	}

	return symbols, digits
}
