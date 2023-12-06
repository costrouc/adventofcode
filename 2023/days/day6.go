package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/costrouc/adventofcode/tools"
)

func ParseDay6Part1(path string) ([]int, []int) {
	lines := tools.ReadFileLines(path)
	times := []int{}
	distances := []int{}

	numberRegex := regexp.MustCompile("[0-9]+")
	for _, token := range numberRegex.FindAllString(lines[0], -1) {
		v, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		times = append(times, v)
	}

	for _, token := range numberRegex.FindAllString(lines[1], -1) {
		v, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		distances = append(distances, v)
	}

	return times, distances
}

func CalculateDistance(t, h int) int {
	return t*h - h*h
}

func Day6Part1(path string) int {
	times, distances := ParseDay6Part1(path)
	total := 1

	for i, time := range times {
		combinations := 0
		for holdTime := 1; holdTime < time; holdTime++ {
			if CalculateDistance(time, holdTime) > distances[i] {
				combinations++
			}
		}
		total *= combinations
	}
	return total
}

func Day6Part2(path string) int {
	times, distances := ParseDay6Part1(path)
	timesString := []string{}
	distancesString := []string{}
	for i := 0; i < len(times); i++ {
		timesString = append(timesString, fmt.Sprintf("%d", times[i]))
		distancesString = append(distancesString, fmt.Sprintf("%d", distances[i]))
	}
	time, err := strconv.Atoi(strings.Join(timesString, ""))
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.Join(distancesString, ""))
	if err != nil {
		panic(err)
	}
	fmt.Println(time, distance)

	total := 1
	combinations := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		if CalculateDistance(time, holdTime) > distance {
			combinations++
		}
	}
	total *= combinations
	return total
}
