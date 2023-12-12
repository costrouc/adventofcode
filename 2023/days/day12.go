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

func ValidateRecord(record Record, combination int) bool {
	currentDamagedStreak := 0
	location := 0
	index := 0

	for i := 0; i < len(record.Entries); i++ {
		switch record.Entries[i] {
		case '.':
			if currentDamagedStreak > 0 {
				if currentDamagedStreak != record.Groups[index] {
					return false
				}
				currentDamagedStreak = 0
				index++
			}
		case '#':
			if index == len(record.Groups) {
				return false
			}

			currentDamagedStreak++

			if currentDamagedStreak > record.Groups[index] {
				return false
			}
		case '?':
			if combination&(1<<location) == 0 { // .
				if currentDamagedStreak > 0 {
					if currentDamagedStreak != record.Groups[index] {
						return false
					}
					currentDamagedStreak = 0
					index++
				}
			} else { // #
				if index == len(record.Groups) {
					return false
				}

				currentDamagedStreak++

				if currentDamagedStreak > record.Groups[index] {
					return false
				}
			}

			location++
		}
	}

	// must have gone through every group
	if currentDamagedStreak == 0 && index != len(record.Groups) {
		return false
	}

	// if there is a damaged streak at the end, it must match the last group
	if currentDamagedStreak > 0 && (index != len(record.Groups)-1 || currentDamagedStreak != record.Groups[index]) {
		return false
	}

	return true
}

func RecordCombinations(record Record) int {
	locations := []int{}
	for i := 0; i < len(record.Entries); i++ {
		switch record.Entries[i] {
		case '?':
			locations = append(locations, i)
		}
	}

	validCombinations := 0
	totalCombinations := 1 << len(locations)
	for i := 0; i < totalCombinations; i++ {
		valid := ValidateRecord(record, i)
		if valid {
			validCombinations++
		}
	}
	return validCombinations
}

func Day12Part1(path string) int {
	records := ParseRecords(path)
	total := 0
	for _, record := range records {
		total += RecordCombinations(record)
	}
	return total
}

func MulFiveRecord(record Record) Record {
	groups := []int{}
	for i := 0; i < 5; i++ {
		groups = append(groups, record.Groups...)
	}
	return Record{
		Entries: []rune(strings.Repeat(string(record.Entries), 5)),
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
