package days

import (
	"path/filepath"
	"testing"
)

func TestDay12Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day12-example1.txt")
	result := Day12Part1(path)
	t.Errorf("Day12Part1 Example1 %v", result)
}

func TestDay12Part1Example2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day12-example2.txt")
	result := Day12Part1(path)
	t.Errorf("Day12Part1 Example2 %v", result)
}

func TestDay12PArt1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day12-input1.txt")
	result := Day12Part1(path)
	t.Errorf("Day12Part1 Input1 %v", result)
}

func TestDay12Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day12-example1.txt")
	result := Day12Part2(path)
	t.Errorf("Day12Part2 Example1 %v", result)
}

func TestDay12Part2Example2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day12-example2.txt")
	result := Day12Part2(path)
	t.Errorf("Day12Part2 Example2 %v", result)
}

func TestDay12Part2Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day12-input1.txt")
	result := Day12Part2(path)
	t.Errorf("Day12Part2 Input1 %v", result)
}
