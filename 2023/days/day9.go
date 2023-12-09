package days

import (
	"strconv"
	"strings"

	"github.com/costrouc/adventofcode/tools"
)

func ParseDay9(path string) [][]int {
	lines := tools.ReadFileLines(path)
	sequences := [][]int{}

	for _, line := range lines {
		sequence := []int{}
		for _, token := range strings.Split(line, " ") {
			i, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			sequence = append(sequence, i)
		}
		sequences = append(sequences, sequence)
	}

	return sequences
}

func AllZero(sequence []int) bool {
	for _, i := range sequence {
		if i != 0 {
			return false
		}
	}
	return true
}

func DiffSequence(sequence []int) []int {
	diffs := []int{}
	for i := 1; i < len(sequence); i++ {
		diffs = append(diffs, sequence[i]-sequence[i-1])
	}
	return diffs
}

func Diffs(sequence []int) [][]int {
	diffs := [][]int{}
	diffs = append(diffs, sequence)
	for !AllZero(diffs[len(diffs)-1]) {
		diffs = append(diffs, DiffSequence(diffs[len(diffs)-1]))
	}
	return diffs
}

func PreviousNumber(diffs [][]int) int {
	n := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		n = diffs[i][0] - n
		diffs[i] = append([]int{n}, diffs[i]...)
	}
	return n
}

func NextNumber(diffs [][]int) int {
	n := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		n = diffs[i][len(diffs[i])-1] + n
		diffs[i] = append(diffs[i], n)
	}
	return n
}

func Day9Part1(path string) int {
	sequences := ParseDay9(path)
	total := 0
	for _, sequence := range sequences {
		diffs := Diffs(sequence)
		next := NextNumber(diffs)
		total += next
	}
	return total
}

func Day9Part2(path string) int {
	sequences := ParseDay9(path)
	total := 0
	for _, sequence := range sequences {
		diffs := Diffs(sequence)
		next := PreviousNumber(diffs)
		total += next
	}
	return total
}
