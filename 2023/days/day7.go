package days

import (
	"cmp"
	"slices"
	"strconv"
	"strings"

	"github.com/costrouc/adventofcode/tools"
)

type Hand struct {
	Cards []rune
	Bid   int
}

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_PAIR
	FULL_HOUSE
	FOUR_KIND
	FIVE_KIND
)

var cardOrderPart1 = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardOrderPart2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func IdentifyPart1Hand(h Hand) int {
	countMap := make(map[rune]int)
	for _, r := range h.Cards {
		countMap[r]++
	}
	counts := []int{}
	for _, i := range countMap {
		counts = append(counts, i)
	}
	slices.Sort(counts)

	if len(counts) == 5 {
		return HIGH_CARD
	} else if len(counts) == 4 {
		return ONE_PAIR
	} else if len(counts) == 3 && counts[2] == 2 {
		return TWO_PAIR
	} else if len(counts) == 3 && counts[2] == 3 {
		return THREE_PAIR
	} else if len(counts) == 2 && counts[1] == 3 {
		return FULL_HOUSE
	} else if len(counts) == 2 && counts[1] == 4 {
		return FOUR_KIND
	} else {
		return FIVE_KIND
	}
}

func IdentifyPart2Hand(h Hand) int {
	countMap := make(map[rune]int)
	jokers := 0
	for _, r := range h.Cards {
		if r == 'J' {
			jokers++
		} else {
			countMap[r]++
		}
	}
	counts := []int{}
	for _, i := range countMap {
		counts = append(counts, i)
	}
	slices.Sort(counts)
	if len(counts) == 0 {
		counts = append(counts, jokers)
	} else {
		counts[len(counts)-1] = counts[len(counts)-1] + jokers
	}

	if len(counts) == 5 {
		return HIGH_CARD
	} else if len(counts) == 4 {
		return ONE_PAIR
	} else if len(counts) == 3 && counts[2] == 2 {
		return TWO_PAIR
	} else if len(counts) == 3 && counts[2] == 3 {
		return THREE_PAIR
	} else if len(counts) == 2 && counts[1] == 3 {
		return FULL_HOUSE
	} else if len(counts) == 2 && counts[1] == 4 {
		return FOUR_KIND
	} else {
		return FIVE_KIND
	}
}

func ComparePart1Hands(a Hand, b Hand) int {
	v1 := IdentifyPart1Hand(a)
	v2 := IdentifyPart1Hand(b)
	if v1 == v2 {
		for i := 0; i < len(a.Cards); i++ {
			if cardOrderPart1[a.Cards[i]] != cardOrderPart1[b.Cards[i]] {
				return cmp.Compare[int](cardOrderPart1[a.Cards[i]], cardOrderPart1[b.Cards[i]])
			}
		}
	}
	return cmp.Compare[int](v1, v2)
}

func ComparePart2Hands(a Hand, b Hand) int {
	v1 := IdentifyPart2Hand(a)
	v2 := IdentifyPart2Hand(b)
	if v1 == v2 {
		for i := 0; i < len(a.Cards); i++ {
			if cardOrderPart2[a.Cards[i]] != cardOrderPart2[b.Cards[i]] {
				return cmp.Compare[int](cardOrderPart2[a.Cards[i]], cardOrderPart2[b.Cards[i]])
			}
		}
	}
	return cmp.Compare[int](v1, v2)
}

func ParseHands(path string) []Hand {
	lines := tools.ReadFileLines(path)
	hands := []Hand{}
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		v, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{
			Cards: []rune(tokens[0]),
			Bid:   v,
		})
	}
	return hands
}

func Day7Part1(path string) int {
	hands := ParseHands(path)
	slices.SortFunc(hands, ComparePart1Hands)

	total := 0
	for i, h := range hands {
		total += h.Bid * (i + 1)
	}
	return total
}

func Day7Part2(path string) int {
	hands := ParseHands(path)
	slices.SortFunc(hands, ComparePart2Hands)

	total := 0
	for i, h := range hands {
		total += h.Bid * (i + 1)
	}
	return total
}
