package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/filipweidemann/advent-of-code/utils"
)

//go:embed input-test.txt
var testInput string
var testInputLines []string

//go:embed input.txt
var input string
var inputLines []string

type Symbol struct {
	LineID int
	Index  int
}

type ParsedLine struct {
	LineID  int
	Symbols []Symbol
}

func init() {
	testInputLines = strings.Split(testInput, "\n")
	inputLines = strings.Split(input, "\n")
}

func main() {
	sum := Part1(inputLines)
	fmt.Printf("Part 1: %v \n", sum)

	power := Part2(inputLines)
	fmt.Printf("Part 2: %v", power)
}

func Part1(i []string) int {
	symbols := make([]Symbol, 0)

	for ind, l := range i {
		lineSymbols := ParseLine(l, ind)
		for _, ls := range lineSymbols {
			symbols = append(symbols, ls)
		}
	}

	sum := 0
	for symbolIndex, symbol := range symbols {
		fmt.Printf("Symbol: %v --- ", symbol)
		fmt.Printf("Running for symbol #%v \n", symbolIndex)

		sum += GetAdjacentNumbersSum(symbol, &i)
	}

	return sum
}

func Part2(i []string) uint64 {
	return 0
}

func GetAdjacentNumbersSum(s Symbol, l *[]string) int {
	startLine := utils.MinWithBound(s.LineID-1, 0)
	endLine := utils.MaxWithBound(s.LineID+1, len((*l))-1)
	startIndex := utils.MinWithBound(int(s.Index), 0)
	endIndex := utils.MaxWithBound(int(s.Index)+1, len((*l)[0])-1)

	digits := make([]int, 0)
	il := len((*l))

	for i := startLine; i <= endLine; i++ {
		if i == il {
			break
		}

		lineDigits := ""
		ll := len((*l)[0])
		for j := startIndex; j <= endIndex; j++ {
			if j == ll || len((*l)[i]) == 0 {
				break
			}

			if c := (*l)[i][j]; unicode.IsDigit(rune(c)) {
				lineDigits = lineDigits + string(c)
			} else {
				lineDigits = lineDigits + " "
			}
		}

		// No digit in string
		if strings.TrimSpace(lineDigits) == "" {
			continue
		}

		// If left side contains a digit, scan further to the left
		if !strings.HasPrefix(lineDigits, " ") {
			leftExtension := ""
			for li := startIndex - 1; li >= 0; li-- {
				if c := (*l)[i][li]; unicode.IsDigit(rune(c)) {
					leftExtension = string(c) + leftExtension
				} else {
					break
				}
			}
			lineDigits = leftExtension + lineDigits
		}

		// If right side contains a digit, scan further to the left
		if !strings.HasSuffix(lineDigits, " ") {
			rightExtension := ""
			for ri := endIndex + 1; ri < len(*l); ri++ {
				if ri == ll {
					break
				}
				if c := (*l)[i][ri]; unicode.IsDigit(rune(c)) {
					rightExtension = string(c) + rightExtension
				} else {
					break
				}
			}
			lineDigits = lineDigits + rightExtension
		}

		// Line is completely evaluated, convert the digits to actual uints
		ts := strings.TrimSpace(lineDigits)
		for _, d := range strings.Split(ts, " ") {
			pd, err := strconv.Atoi(d)
			if err == nil {
				digits = append(digits, pd)
			}
		}
	}

	sum := 0
	for _, val := range digits {
		sum += val
	}

	return sum
}

func ParseLine(l string, ind int) []Symbol {
	symbols := make([]Symbol, 0)

	for i, r := range l {
		if unicode.IsDigit(r) || r == '.' {
			continue
		}

		symbol := Symbol{
			LineID: ind,
			Index:  i,
		}

		symbols = append(symbols, symbol)
	}

	return symbols
}
