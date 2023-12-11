package days

import (
	"path/filepath"
	"testing"
)

func TestDay10Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-example1.txt")
	result := Day10Part1(path)
	t.Errorf("Day10Part1 Example1 %v", result)
}

func TestDay10Part1Example2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-example2.txt")
	result := Day10Part1(path)
	t.Errorf("Day10Part1 Example2 %v", result)
}

func TestDay10Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-input1.txt")
	result := Day10Part1(path)
	t.Errorf("Day10Part1 Input1 %v", result)
}

func TestDay10Part2Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-example1.txt")
	result := Day10Part2(path)
	t.Errorf("Day10Part2 Example2 %v", result)
}

func TestDay10Part2Example3(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-example3.txt")
	result := Day10Part2(path)
	t.Errorf("Day10Part2 Example3 %v", result)
}

func TestDay10Part2Example4(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-example4.txt")
	result := Day10Part2(path)
	t.Errorf("Day10Part2 Example4 %v", result)
}

func TestDay10Part2Example5(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-example5.txt")
	result := Day10Part2(path)
	t.Errorf("Day10Part2 Example5 %v", result)
}

func TestDay10Part2Input2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day10-input1.txt")
	result := Day10Part2(path)
	t.Errorf("Day10Part2 Input1 %v", result)
}
