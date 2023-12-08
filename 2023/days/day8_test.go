package days

import (
	"path/filepath"
	"testing"
)

func TestDay8Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day8-example1.txt")
	result := Day8Part1(path)
	t.Errorf("Day8Part1 Example1 %v", result)
}

func TestDay8PArt1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day8-input1.txt")
	result := Day8Part1(path)
	t.Errorf("Day8Part1 Input1 %v", result)
}

func TestDay8Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day8-example2.txt")
	result := Day8Part2(path)
	t.Errorf("Day8Part2 Example2 %v", result)
}

func TestDay8Part2Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day8-input1.txt")
	result := Day8Part2(path)
	t.Errorf("Day8Part2 Input1 %v", result)
}
