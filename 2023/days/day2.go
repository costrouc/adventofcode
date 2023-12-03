package days

import (
	"bufio"
	"os"

	"github.com/costrouc/adventofcode/tools"
)

const (
	GAME      int = 10
	BLUE      int = 11
	GREEN     int = 12
	RED       int = 13
	COLON     int = 14
	COMMA     int = 15
	SEMICOLON int = 16
	NEWLINE   int = 17
)

type Bag struct {
	Red   int
	Blue  int
	Green int
}

func Max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func Day2Part1(path string) map[int]Bag {
	dictionary := map[string]int{
		"Game":  GAME,
		"blue":  BLUE,
		"green": GREEN,
		"red":   RED,
		":":     COLON,
		",":     COMMA,
		";":     SEMICOLON,
		"\n":    NEWLINE,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	trie := tools.NewTrie()
	for token, value := range dictionary {
		trie.Set(token, value)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	games := make(map[int]Bag)
	currentGame := 0
	lastValue := 0
	currentBag := Bag{}
	var EOF bool
	for !EOF {
		token := trie.NextMatch(reader)
		switch token {
		case -1:
			EOF = true
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			lastValue = lastValue*10 + token
		case BLUE:
			currentBag.Blue = Max(currentBag.Blue, lastValue)
		case GREEN:
			currentBag.Green = Max(currentBag.Green, lastValue)
		case RED:
			currentBag.Red = Max(currentBag.Red, lastValue)
		case COLON:
			currentGame = lastValue
			lastValue = 0
		case COMMA, SEMICOLON:
			lastValue = 0
		case NEWLINE:
			games[currentGame] = currentBag
			lastValue = 0
		case GAME:
			currentBag = Bag{}
		}
	}

	return games
}
