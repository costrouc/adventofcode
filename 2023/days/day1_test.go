package days

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestDay1Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day1-example1.txt")
	result := Day1Part1(path)
	fmt.Printf("day1 example1: %v\n", result)
}

func TestDay1Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day1-input1.txt")
	result := Day1Part1(path)
	fmt.Printf("day1 part1 %v\n", result)
}

func TestDay1Part2Example2(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day1-example2.txt")
	result := Day1Part2(path)
	fmt.Printf("day1 part2 example2 %v\n", result)
}

func TestDay1Part2Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day1-input1.txt")
	result := Day1Part2(path)
	fmt.Printf("day1 part2 input1 %v\n", result)
}

func BenchmarkDay1Part2Input1(b *testing.B) {
	path := filepath.Join("..", "puzzles", "day1-input1.txt")
	for i := 0; i < b.N; i++ {
		Day1Part2(path)
	}
}
