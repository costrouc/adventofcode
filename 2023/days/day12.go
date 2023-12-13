package days

import (
	"strconv"
	"strings"

	"github.com/costrouc/adventofcode/tools"
)

type Record struct {
	Entries []rune
	Groups  []int
}

func ParseRecords(path string) []Record {
	records := []Record{}
	lines := tools.ReadFileLines(path)

	for _, line := range lines {
		tokens := strings.Split(line, " ")

		groups := []int{}
		for _, token := range strings.Split(tokens[1], ",") {
			i, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			groups = append(groups, i)
		}

		records = append(records, Record{
			Entries: []rune(tokens[0]),
			Groups:  groups,
		})
	}

	return records
}

func ValidateRecord(record *Record, recordGuess []int) (bool, int) {
	currentDamagedStreak := 0
	groupIndex := 0
	guessLocation := 0

	for i := 0; i < len(record.Entries); i++ {
		switch record.Entries[i] {
		case '.':
			if currentDamagedStreak > 0 {
				if currentDamagedStreak != record.Groups[groupIndex] {
					return false, guessLocation
				}
				currentDamagedStreak = 0
				groupIndex++
			}
		case '#':
			if groupIndex == len(record.Groups) {
				return false, guessLocation
			}

			currentDamagedStreak++

			if currentDamagedStreak > record.Groups[groupIndex] {
				return false, guessLocation
			}
		case '?':
			if recordGuess[guessLocation] == 0 { // .
				if currentDamagedStreak > 0 {
					if currentDamagedStreak != record.Groups[groupIndex] {
						return false, guessLocation
					}
					currentDamagedStreak = 0
					groupIndex++
				}
			} else { // #
				if groupIndex == len(record.Groups) {
					return false, guessLocation
				}

				currentDamagedStreak++

				if currentDamagedStreak > record.Groups[groupIndex] {
					return false, guessLocation
				}
			}

			guessLocation++
		}
	}

	// must have gone through every group
	if currentDamagedStreak == 0 && groupIndex != len(record.Groups) {
		return false, guessLocation
	}

	// if there is a damaged streak at the end, it must match the last group
	if currentDamagedStreak > 0 && (groupIndex != len(record.Groups)-1 || currentDamagedStreak != record.Groups[groupIndex]) {
		return false, guessLocation
	}

	return true, guessLocation
}

func PartialRecordCombinations(record *Record, guess []int, index int) int {
	valid, location := ValidateRecord(record, guess)
	if index == len(guess) && valid {
		return 1
	} else if index == len(guess) {
		return 0
	} else if location < index {
		return 0
	}

	total := 0
	guess[index] = 0
	total += PartialRecordCombinations(record, guess, index+1)

	guess[index] = 1
	total += PartialRecordCombinations(record, guess, index+1)
	return total
}

func RecordCombinations(record Record) int {
	locations := 0
	for i := 0; i < len(record.Entries); i++ {
		switch record.Entries[i] {
		case '?':
			locations++
		}
	}
	recordGuess := make([]int, locations)
	return PartialRecordCombinations(&record, recordGuess, 0)
}

func Day12Part1(path string) int {
	records := ParseRecords(path)
	total := 0
	for _, record := range records {
		result := RecordCombinations(record)
		total += result
	}
	return total
}

func MulFiveRecord(record Record) Record {
	n := 5
	groups := []int{}
	entities := []string{}
	for i := 0; i < n; i++ {
		groups = append(groups, record.Groups...)
		entities = append(entities, string(record.Entries))
	}
	return Record{
		Entries: []rune(strings.Join(entities, "?")),
		Groups:  groups,
	}
}

func Day12Part2(path string) int {
	records := ParseRecords(path)
	total := 0
	for _, record := range records {
		fiveRecord := MulFiveRecord(record)
		total += RecordCombinations(fiveRecord)
	}
	return total
}
