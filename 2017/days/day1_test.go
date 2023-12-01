package days

import (
	"path/filepath"
	"testing"
)

func TestDay1Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day1-sample1.txt")
	Day1Part12(path)
	t.Errorf("failed")
}

func TestDay1Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day1-input1.txt")
	Day1Part12(path)
	t.Errorf("failed")
}
