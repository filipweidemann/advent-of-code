package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/filipweidemann/advent-of-code/utils"
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

func main() {
	sum := Part1(inputLines)
	fmt.Printf("Part 1: %v \n", sum)

	power := Part2(inputLines)
	fmt.Printf("Part 2: %v", power)
}

type Card struct {
	WinningNumbers []string
	DrawnNumbers   []string
	Instances      uint
	CardNumber     int
}

func (c *Card) FromLine(l string) {
	vals := strings.Split(l, ":")[1]
	nums := strings.Split(vals, "|")

	replacedWinningNumbers := strings.TrimSpace(strings.ReplaceAll(nums[0], "  ", " "))
	replacedDrawnNumbers := strings.TrimSpace(strings.ReplaceAll(nums[1], "  ", " "))
	c.WinningNumbers = strings.Split(replacedWinningNumbers, " ")
	c.DrawnNumbers = strings.Split(replacedDrawnNumbers, " ")
}

func (c *Card) Worth() float64 {
	worth := 0
	for _, v := range c.WinningNumbers {
		if slices.Contains(c.DrawnNumbers, v) {
			worth += 1
		}
	}

	// do some math
	if worth != 0 {
		return math.Pow(2, float64(worth-1))
	}

	// sad
	return 0
}

func (c *Card) PayoutNextCards(cards *[]Card) {
	worth := int(c.Worth())
	if worth == 0 {
		return
	}

	nextCards := utils.MaxWithBound(worth, len(*cards)-c.CardNumber)
	fmt.Println("Evaluating Card ", c.CardNumber, " with worth ", worth)

	i := 1
	for i < nextCards {
		if c.CardNumber < 5 {
			fmt.Printf("Current: Card %v, has %v instances, adding %v to %v instances of next card! \n", c.CardNumber, c.Instances, c.Instances, (*cards)[c.CardNumber+i].Instances)
		}
		(*cards)[c.CardNumber+i].Instances += c.Instances
		i += 1
	}

}

func ParseInput(lns []string) *[]Card {
	cards := new([]Card)

	for idx, l := range lns {
		c := Card{Instances: 1, CardNumber: idx}
		c.FromLine(l)
		*cards = append(*cards, c)
	}

	return cards
}

func Part1(i []string) int {
	cards := ParseInput(i)

	combinedWorth := float64(0)
	for _, c := range *cards {
		combinedWorth += c.Worth()
	}

	return int(combinedWorth)
}

func Part2(i []string) uint {
	cards := ParseInput(i)

	for _, c := range *cards {
		c.PayoutNextCards(cards)
	}

	return 0
}
