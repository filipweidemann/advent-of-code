package main

import (
	"testing"
)

var unitTestInput = []string{
	"10.*10...20...",
	"5...40..5$..$$",
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example-input",
			input: unitTestInput,
			want:  75,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.input); got != tt.want {
				t.Errorf("Part1: got %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestLineParser(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  uint
	}{
		{
			name:  "example-input",
			input: unitTestInput,
			want:  70,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			symbols := ParseLine(tt.input[1], 1)
			if symbols[0].Index != 9 {
				t.Errorf("ParseLine returned wrong parsed line, expected 10, got %v", symbols[0].Index)
			}
		})
	}

}

func TestGetAdjacentNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input1 Symbol
		input2 []string
		want   int
	}{
		//{
		//	name:   "example-input",
		//	input1: Symbol{Index: 9, LineID: 1},
		//	input2: unitTestInput,
		//	want:   25,
		//},
		//{
		//	name:   "example-input",
		//	input1: Symbol{Index: 13, LineID: 1},
		//	input2: unitTestInput,
		//	want:   0,
		//},
		{
			name:   "example-input",
			input1: Symbol{Index: 3, LineID: 0},
			input2: unitTestInput,
			want:   50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := GetAdjacentNumbersSum(tt.input1, &tt.input2)
			if nums != tt.want {
				t.Errorf("GetAdjacentNumbersSum, expected %v, got %v", tt.want, nums)
			}
		})
	}

}
