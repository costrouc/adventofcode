package days

import (
	"path/filepath"
	"testing"
)

func TestDay4Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day4-example1.txt")
	result := Day4Part1(path)
	t.Errorf("Day4Part1 Example1 %v", result)
}

func TestDay4Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day4-input1.txt")
	result := Day4Part1(path)
	t.Errorf("Day4Part1 Input1 %v", result)
}

func TestDay4Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day4-example1.txt")
	result := Day4Part2(path)
	t.Errorf("Day4Part2 Example1 %v", result)
}

func TestDay4Part2Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day4-input1.txt")
	result := Day4Part2(path)
	t.Errorf("Day4Part2 Example1 %v", result)
}
