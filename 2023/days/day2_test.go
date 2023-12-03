package days

import (
	"path/filepath"
	"testing"
)

func TestDay2Part1Example1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day2-example1.txt")
	games := Day2Part1(path)

	bag := Bag{Red: 12, Green: 13, Blue: 14}
	totalSum := 0
	totalPower := 0
	for key, value := range games {
		if value.Blue <= bag.Blue && value.Red <= bag.Red && value.Green <= bag.Green {
			totalSum += key
		}
		totalPower += value.Blue * value.Red * value.Green
	}
	t.Errorf("total sum %v %v", totalSum, totalPower)
}

func TestDay2Part1Input1(t *testing.T) {
	path := filepath.Join("..", "puzzles", "day2-input1.txt")
	games := Day2Part1(path)

	bag := Bag{Red: 12, Green: 13, Blue: 14}
	totalSum := 0
	totalPower := 0
	for key, value := range games {
		if value.Blue <= bag.Blue && value.Red <= bag.Red && value.Green <= bag.Green {
			totalSum += key
		}
		totalPower += value.Blue * value.Red * value.Green
	}
	t.Errorf("total sum %v %v", totalSum, totalPower)
}
