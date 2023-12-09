package days

import (
	"path/filepath"
	"testing"
)

func TestDay9Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day9-example1.txt")
	result := Day9Part1(path)
	t.Errorf("Day9Part1 Example1 %v", result)
}

func TestDay9PArt1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day9-input1.txt")
	result := Day9Part1(path)
	t.Errorf("Day9Part1 Input1 %v", result)
}

func TestDay9Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day9-example1.txt")
	result := Day9Part2(path)
	t.Errorf("Day9Part2 Example2 %v", result)
}

func TestDay9Part2Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day9-input1.txt")
	result := Day9Part2(path)
	t.Errorf("Day9Part2 Input1 %v", result)
}
