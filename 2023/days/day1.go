package days

import (
	"bufio"
	"os"

	"github.com/costrouc/adventofcode/tools"
)

func Day1Part1(path string) int {
	dictionary := map[string]int{
		"0":  0,
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"\n": 10,
	}

	return ScoreLines(path, dictionary)
}

func Day1Part2(path string) int {
	dictionary := map[string]int{
		"0":     0,
		"zero":  0,
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
		"\n":    10,
	}

	return ScoreLines(path, dictionary)
}

func ScoreLines(path string, dictionary map[string]int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	trie := tools.NewTrie()
	oldTriePositions := []*tools.Trie{}
	newTriePositions := []*tools.Trie{}

	for word := range dictionary {
		trie.Set(word, dictionary[word])
	}

	totalSum := 0
	firstInt := -1
	lastInt := -1

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		t := trie.Next(rune(r))
		if t != nil {
			newTriePositions = append(newTriePositions, t)
		}
		for i := range oldTriePositions {
			t := oldTriePositions[i].Next(r)
			if t != nil {
				newTriePositions = append(newTriePositions, t)
			}
		}

		for i := range newTriePositions {
			if newTriePositions[i].Terminal {
				value := newTriePositions[i].Value
				switch value {
				case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
					if firstInt < 0 {
						firstInt = value
						lastInt = value
					} else {
						lastInt = value
					}
				case 10:
					totalSum += firstInt*10 + lastInt
					firstInt, lastInt = -1, -1
				}
			}
		}

		oldTriePositions = newTriePositions
		newTriePositions = []*tools.Trie{}
	}

	return totalSum
}
