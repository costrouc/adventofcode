package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/costrouc/adventofcode/tools"
)

type DestinationSourceRange struct {
	DestinationStart int
	SourceStart      int
	RangeLength      int
}

type SeedMap struct {
	From   string
	To     string
	Ranges []DestinationSourceRange
}

func ParseAlminac(path string) ([]SeedMap, []int) {
	lines := tools.ReadFileLines(path)

	seeds := []int{}
	seedMaps := []SeedMap{}
	for _, v := range strings.Split(lines[0], " ")[1:] {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, i)
	}

	blankLine := true
	currentSeedMap := SeedMap{}
	for _, line := range lines[2:] {
		if line == "" {
			seedMaps = append(seedMaps, currentSeedMap)
			blankLine = true
			currentSeedMap = SeedMap{}
			continue
		}

		if blankLine {
			nameMapping := strings.Split(strings.Split(line, " ")[0], "-")
			currentSeedMap.From = nameMapping[0]
			currentSeedMap.To = nameMapping[2]
			blankLine = false
		} else {
			tokens := strings.Split(line, " ")
			destinationStart, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			sourceStart, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			rangeLength, err := strconv.Atoi(tokens[2])
			if err != nil {
				panic(err)
			}
			currentSeedMap.Ranges = append(currentSeedMap.Ranges, DestinationSourceRange{
				DestinationStart: destinationStart,
				SourceStart:      sourceStart,
				RangeLength:      rangeLength,
			})
		}
	}

	return seedMaps, seeds
}

func CalculateMapping(ranges []DestinationSourceRange, seed int) int {
	for _, rng := range ranges {
		if seed >= rng.SourceStart && seed <= rng.SourceStart+rng.RangeLength {
			return rng.DestinationStart + (seed - rng.SourceStart)
		}
	}
	return seed
}

func Day5Part1(path string) int {
	seedMaps, seeds := ParseAlminac(path)
	// fmt.Println(seedMaps, seeds)

	minLocation := math.MaxInt64
	for _, seed := range seeds {
		v := []int{seed}
		for i, seedMap := range seedMaps {
			v = append(v, CalculateMapping(seedMap.Ranges, v[i]))
		}

		if v[len(v)-1] < minLocation {
			minLocation = v[len(v)-1]
		}
	}

	return minLocation
}

func Day5Part2(path string) int {
	seedMaps, seeds := ParseAlminac(path)

	minLocation := math.MaxInt64
	for i := 0; i < len(seeds)/2; i++ {
		fmt.Println("seed", seeds[i*2])
		for seed := seeds[i*2]; seed < seeds[i*2]+seeds[i*2+1]; seed++ {
			l := seed
			for _, seedMap := range seedMaps {
				l = CalculateMapping(seedMap.Ranges, l)
			}

			if l < minLocation {
				minLocation = l
			}
		}
	}

	return minLocation
}
