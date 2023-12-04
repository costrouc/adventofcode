package days

import (
	"path/filepath"
	"testing"
)

func TestDay3Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day3-example1.txt")
	result := Day3Part1(path)
	t.Errorf("Day3Part1 %v", result)
}

func TestDay3Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day3-input1.txt")
	result := Day3Part1(path)
	t.Errorf("Day3Part1 %v", result)
}

func TestDay3Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day3-example1.txt")
	result := Day3Part2(path)
	t.Errorf("Day3Part2 %v", result)
}

func TestDay3Part2Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day3-input1.txt")
	result := Day3Part2(path)
	t.Errorf("Day3Part2 %v", result)
}
