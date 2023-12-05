package days

import (
	"path/filepath"
	"testing"
)

func TestDay5Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day5-example1.txt")
	result := Day5Part1(path)
	t.Errorf("Day5Part1 Example1: %v", result)
}

func TestDay5Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day5-input1.txt")
	result := Day5Part1(path)
	t.Errorf("Day5Part1 Input1: %v", result)
}

func TestDay5Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day5-example1.txt")
	result := Day5Part2(path)
	t.Errorf("Day5Part2 Example1: %v", result)
}

func TestDay5Part2Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day5-input1.txt")
	result := Day5Part2(path)
	t.Errorf("Day5Part2 Input1: %v", result)
}
