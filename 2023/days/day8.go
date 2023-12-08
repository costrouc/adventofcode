package days

import (
	"regexp"

	"github.com/costrouc/adventofcode/tools"
)

func ParseMap(path string) ([]int, map[string][]string) {
	lines := tools.ReadFileLines(path)

	desertmap := make(map[string][]string)

	instructions := []int{}
	for _, r := range lines[0] {
		if r == 'R' {
			instructions = append(instructions, 1)
		} else if r == 'L' {
			instructions = append(instructions, 0)
		}
	}

	lineRegex := regexp.MustCompile(`([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)`)
	for _, line := range lines[2:] {
		matches := lineRegex.FindStringSubmatch(line)
		desertmap[matches[1]] = matches[2:4]
	}

	return instructions, desertmap
}

func Day8Part1(path string) int {
	instructions, desertmap := ParseMap(path)

	currentNode := "AAA"
	steps := 0
	for i := 0; currentNode != "ZZZ"; i = (i + 1) % len(instructions) {
		currentNode = desertmap[currentNode][instructions[i]]
		steps++
	}

	return steps
}

func AllNodesMatch(nodes []string) bool {
	for _, node := range nodes {
		if node[len(node)-1] != 'Z' {
			return false
		}
	}
	return true
}

func AllNotZero(nodes []int) bool {
	for _, node := range nodes {
		if node == 0 {
			return false
		}
	}
	return true
}

func Day8Part2(path string) int {
	instructions, desertmap := ParseMap(path)

	currentNodes := []string{}
	periods := []int{}
	for node := range desertmap {
		if node[len(node)-1] == 'A' {
			currentNodes = append(currentNodes, node)
			periods = append(periods, 0)
		}
	}

	steps := 0
	for i := 0; !AllNotZero(periods); i = (i + 1) % len(instructions) {
		for j, node := range currentNodes {
			if node[len(node)-1] == 'Z' && periods[j] == 0 {
				periods[j] = steps
			}

			currentNodes[j] = desertmap[node][instructions[i]]
		}
		steps++
	}

	return tools.LCM(periods[0], periods[1], periods[2:]...)
}
