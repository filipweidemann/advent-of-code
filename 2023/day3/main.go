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

func (s *Symbol) IsGear(l []Line, d []Digit) ([]Digit, bool) {
	// we assume equally long lines, so just grab one and get their length
	dc := DigitCollection{Digits: d}

	digits := dc.AdjacentDigitsFromSymbol(s, l)

	if len(digits) == 2 {
		return digits, true
	}

	return nil, false
}

type Digit struct {
	LineID  int
	Indices []int
	Runes   []rune
	Value   int
}

type DigitCollection struct {
	Digits []Digit
}

func (dc DigitCollection) AdjacentDigitsFromSymbol(s *Symbol, l []Line) []Digit {
	lineLength := len(l[0].Content)

	// vertical indices
	startLine := utils.MinWithBound(s.LineID, 0)
	endLine := utils.MaxWithBound(s.LineID, len(l)-1)

	// horizontal indices
	startIndex := utils.MinWithBound(s.Index, 0)
	endIndex := utils.MaxWithBound(s.Index, lineLength-1)

	fmt.Println("Start Line: ", startLine)
	fmt.Println("End Line: ", endLine)
	fmt.Println("Start Index: ", startIndex)
	fmt.Println("End Index: ", endIndex)

	filteredDigits := new([]Digit)
	for _, d := range dc.Digits {
		if (d.LineID >= startLine && d.LineID <= endLine) && (d.Indices[len(d.Indices)-1] >= startIndex && d.Indices[0] <= endIndex) {
			fmt.Printf("Found Digit for Symbol %v: Line %v, Indices %v \n", s, d.LineID, d.Indices)
			*filteredDigits = append(*filteredDigits, d)
		}
	}

	fmt.Println("\n\n")

	return *filteredDigits
}

func (d *Digit) HasAdjacentSymbol(l []Line) bool {
	// we assume equally long lines, so just grab one and get their length
	lineLength := len(l[0].Content)

	// vertical indices
	startLine := utils.MinWithBound(d.LineID, 0)
	endLine := utils.MaxWithBound(d.LineID, len(l)-1)

	// horizontal indices
	startIndex := utils.MinWithBound(d.Indices[0], 0)
	endIndex := utils.MaxWithBound(d.Indices[len(d.Indices)-1], lineLength-1)

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
	sum := Part1(inputLines)
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

func Part2(i []string) int {
	lines := LinesFromSlice(i)
	symbols, digits := Collect(lines)

	sum := 0
	for _, sym := range symbols {
		digits, ok := sym.IsGear(lines, digits)
		if ok {
			sum += (digits[0].Value * digits[1].Value)
		}
	}

	return sum
}

func Collect(l []Line) ([]Symbol, []Digit) {
	var symbols []Symbol
	var digits []Digit

	for lineIndex, line := range l {
		digit := Digit{LineID: lineIndex}

		for runeIndex, rune := range line.Content + "." {
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
			value, err := strconv.Atoi(string(digit.Runes))
			if err != nil {
				panic("Could not parse runes to digit")
			}

			digit.Value += value
			digits = append(digits, digit)

			// Clear digit struct
			digit = Digit{LineID: lineIndex}

		}
	}

	return symbols, digits
}
