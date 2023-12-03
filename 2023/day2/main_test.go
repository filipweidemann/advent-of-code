package main

import (
	"testing"
)

func TestCleanLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example-input",
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  "Game 1 3 blue 4 red 1 red 2 green 6 blue 2 green",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanLine(tt.input); got != tt.want {
				t.Errorf("Part1: got %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  uint64
	}{
		{
			name:  "example-input",
			input: testInputLines,
			want:  8,
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

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  uint64
	}{
		{
			name:  "example-input",
			input: testInputLines,
			want:  2286,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.input); got != tt.want {
				t.Errorf("Part1: got %v, expected %v", got, tt.want)
			}
		})
	}
}
