package days

import (
	"path/filepath"
	"testing"
)

func TestDay7Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day7-example1.txt")
	result := Day7Part1(path)
	t.Errorf("Day7Part1 Example1 %v", result)
}

func TestDay7PArt1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day7-input1.txt")
	result := Day7Part1(path)
	t.Errorf("Day7Part1 Input1 %v", result)
}

func TestDay7Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day7-example1.txt")
	result := Day7Part2(path)
	t.Errorf("Day7Part2 Example1 %v", result)
}

func TestDay7Part2Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day7-input1.txt")
	result := Day7Part2(path)
	t.Errorf("Day7Part2 Input1 %v", result)
}
