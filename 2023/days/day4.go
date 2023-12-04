package days

import (
	"math"
	"strconv"
	"strings"

	"github.com/costrouc/adventofcode/tools"
)

type Game struct {
	Winning tools.Set[int]
	Numbers tools.Set[int]
}

func ParseGames(path string) []Game {
	lines := tools.ReadFileLines(path)
	games := []Game{}

	for _, line := range lines {
		game := Game{Winning: tools.NewSet[int](), Numbers: tools.NewSet[int]()}
		inNumbers := false
		for _, token := range strings.Split(line, " ") {
			i, err := strconv.Atoi(token)
			if err != nil {
				if token == "|" {
					inNumbers = true
				}
			} else {
				if inNumbers {
					game.Numbers.Add(i)
				} else {
					game.Winning.Add(i)
				}
			}
		}
		games = append(games, game)
	}
	return games
}

func GameMatches(games []Game) []int {
	matches := []int{}
	for _, game := range games {
		numMatches := game.Winning.Intersection(game.Numbers).Cardinality()
		matches = append(matches, numMatches)
	}
	return matches
}

func Day4Part1(path string) int {
	games := ParseGames(path)
	matches := GameMatches(games)
	total := 0
	for i := 0; i < len(games); i++ {
		total += int(math.Pow(2, float64(matches[i]-1)))
	}
	return total
}

func Day4Part2(path string) int {
	games := ParseGames(path)
	matches := GameMatches(games)
	numCards := make([]int, len(matches))
	for i := 0; i < len(numCards); i++ {
		numCards[i] = numCards[i] + 1
	}

	for i := 0; i < len(matches); i++ {
		for j := 0; (j < matches[i]) && (i+j+1 < len(matches)); j++ {
			numCards[i+j+1] = numCards[i+j+1] + numCards[i]
		}
	}

	return tools.Sum(numCards...)
}
