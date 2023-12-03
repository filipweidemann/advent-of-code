package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input-test.txt
var testInput string
var testInputLines []string

//go:embed input.txt
var input string
var inputLines []string

var TOKENS = []string{"Game", "red", "green", "blue"}

type Game struct {
	ID    uint64
	Cubes Cubes
}

type GameToken struct {
	Token string
	Value interface{}
}

type Cubes struct {
	Red   []uint64
	Green []uint64
	Blue  []uint64
}

func (c *Cubes) MinimumSet() CubeSet {
	return CubeSet{
		slices.Max(c.Red),
		slices.Max(c.Green),
		slices.Max(c.Blue),
	}
}

type CubeSet struct {
	Red   uint64
	Green uint64
	Blue  uint64
}

func (cs *CubeSet) Power() uint64 {
	return cs.Red * cs.Green * cs.Blue
}

type Bag struct {
	Red   uint64
	Green uint64
	Blue  uint64
}

func (b *Bag) IsPossibleGame(g *Game) (uint64, bool) {
	maxCubeSet := g.Cubes.MinimumSet()

	if (b.Red >= maxCubeSet.Red) && (b.Green >= maxCubeSet.Green && (b.Blue >= maxCubeSet.Blue)) {
		return g.ID, true
	}

	return g.ID, false
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

func Part1(i []string) uint64 {
	b := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sumOfGames := uint64(0)

	for _, l := range i {
		if len(l) == 0 {
			continue
		}

		cl := CleanLine(l)
		tokens := TokenizeLine(cl)
		game := EvaluateGameTokens(tokens)
		gameId, wasPossible := b.IsPossibleGame(&game)

		if wasPossible {
			sumOfGames += gameId
		}
	}

	return sumOfGames
}

func Part2(i []string) uint64 {
	powerOfMinSets := uint64(0)

	for _, l := range i {
		if len(l) == 0 {
			continue
		}

		cl := CleanLine(l)
		tokens := TokenizeLine(cl)
		game := EvaluateGameTokens(tokens)
		minCubeSet := game.Cubes.MinimumSet()
		powerOfMinSets += minCubeSet.Power()
	}

	return powerOfMinSets

}

func EvaluateGameTokens(t []GameToken) Game {
	// do not actually treat them as tokens and apply grammars, we know it's a fixed structure
	// also, we don't check the type assertion, it's fine
	gameIdStr, _ := t[1].Value.(string)
	gameId, _ := strconv.ParseUint(gameIdStr, 10, 64)

	cubes := Cubes{}

	game := Game{ID: uint64(gameId), Cubes: cubes}

	for i := 3; i <= len(t); i += 2 {
		tok := t[i].Token
		tok = strings.TrimSpace(tok)
		value, _ := t[i-1].Value.(string)
		valUint, _ := strconv.ParseUint(value, 10, 64)

		switch tok {
		case "red":
			game.Cubes.Red = append(game.Cubes.Red, valUint)
		case "green":
			game.Cubes.Green = append(game.Cubes.Green, valUint)
		case "blue":
			game.Cubes.Blue = append(game.Cubes.Blue, valUint)
		}
	}

	return game

}

func CleanLine(l string) string {
	l = strings.ReplaceAll(l, ",", "")
	l = strings.ReplaceAll(l, ";", "")
	l = strings.Replace(l, ":", "", 1)
	return l
}

func TokenizeLine(l string) []GameToken {
	tokens := make([]GameToken, 0)
	splitted := strings.Split(l, " ")
	for _, t := range splitted {
		if slices.Contains(TOKENS, t) {
			tokens = append(tokens, GameToken{t, ""})
		} else {
			tokens = append(tokens, GameToken{"", t})
		}
	}

	return tokens
}
