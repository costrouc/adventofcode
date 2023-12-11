package days

import (
	"path/filepath"
	"testing"
)

func TestDay11Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day11-example1.txt")
	result := Day11Part1(path)
	t.Errorf("Day11Part1 Example1 %v", result)
}

func TestDay11PArt1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day11-input1.txt")
	result := Day11Part1(path)
	t.Errorf("Day11Part1 Input1 %v", result)
}

func TestDay11Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day11-example1.txt")
	result := Day11Part2(path)
	t.Errorf("Day11Part2 Example1 %v", result)
}

func TestDay11Part2Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day11-input1.txt")
	result := Day11Part2(path)
	t.Errorf("Day11Part2 Input1 %v", result)
}
