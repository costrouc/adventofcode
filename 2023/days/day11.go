package days

import (
	"github.com/costrouc/adventofcode/tools"
)

type Galaxy struct {
	Stars []*tools.Pair
	N     int
	M     int
}

func ParseGalaxy(path string) Galaxy {
	lines := tools.ReadFileLines(path)
	stars := []*tools.Pair{}

	for i, line := range lines {
		for j, token := range line {
			if token == '#' {
				stars = append(stars, &tools.Pair{I: i, J: j})
			}
		}
	}
	return Galaxy{Stars: stars, N: len(lines), M: len(lines[0])}
}

func ExpandGalaxy(galaxy *Galaxy, increment int) {
	emptyRows := []int{}
	for i := galaxy.N - 1; i >= 0; i-- {
		emptyRow := true
		for _, star := range galaxy.Stars {
			if star.I == i {
				emptyRow = false
				break
			}
		}

		if emptyRow {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyColumns := []int{}
	for j := galaxy.M - 1; j >= 0; j-- {
		emptyColumn := true
		for _, star := range galaxy.Stars {
			if star.J == j {
				emptyColumn = false
				break
			}
		}

		if emptyColumn {
			emptyColumns = append(emptyColumns, j)
		}
	}

	galaxy.N += len(emptyRows) * increment
	for _, emptyRow := range emptyRows {
		for _, star := range galaxy.Stars {
			if star.I > emptyRow {
				star.I += increment
			}
		}
	}

	galaxy.M += len(emptyColumns) * increment
	for _, emptyColumn := range emptyColumns {
		for _, star := range galaxy.Stars {
			if star.J > emptyColumn {
				star.J += increment
			}
		}
	}
}

func PairwiseDistance(galaxy Galaxy) [][]int {
	distances := [][]int{}
	for i := 0; i < len(galaxy.Stars); i++ {
		row := []int{}
		for j := 0; j < len(galaxy.Stars); j++ {
			distance := galaxy.Stars[i].Subtract(*galaxy.Stars[j])
			distance = distance.Abs()
			row = append(row, distance.I+distance.J)
		}
		distances = append(distances, row)
	}
	return distances
}

func Day11Part1(path string) int {
	galaxy := ParseGalaxy(path)
	ExpandGalaxy(&galaxy, 1)

	distances := PairwiseDistance(galaxy)

	total := 0
	for i := 0; i < len(distances); i++ {
		for j := i + 1; j < len(distances[0]); j++ {
			total += distances[i][j]
		}
	}

	return total
}

func Day11Part2(path string) int {
	galaxy := ParseGalaxy(path)
	ExpandGalaxy(&galaxy, 999999)

	distances := PairwiseDistance(galaxy)

	total := 0
	for i := 0; i < len(distances); i++ {
		for j := i + 1; j < len(distances[0]); j++ {
			total += distances[i][j]
		}
	}

	return total
}
