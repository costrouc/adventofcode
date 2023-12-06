package days

import (
	"path/filepath"
	"testing"
)

func TestDay6Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day6-example1.txt")
	result := Day6Part1(path)
	t.Errorf("Day6Part1 Example1 %v", result)
}

func TestDay6PArt1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day6-input1.txt")
	result := Day6Part1(path)
	t.Errorf("Day6Part1 Input1 %v", result)
}

func TestDay6Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day6-example1.txt")
	result := Day6Part2(path)
	t.Errorf("Day6Part2 Example1 %v", result)
}

func TestDay6PArt1Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day6-input1.txt")
	result := Day6Part2(path)
	t.Errorf("Day6Part2 Input1 %v", result)
}
