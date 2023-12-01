package days

import (
	"fmt"

	"github.com/costrouc/adventofcode/tools"
)

func Day1Part12(path string) {
	lines := tools.ReadFileLines(path)
	for _, line := range lines {
		totalP1 := 0
		totalP2 := 0
		for i, r := range line {
			if line[i] == line[(i+1)%len(line)] {
				totalP1 += int(r) - int('0')
			}
			if line[i] == line[(i+len(line)/2)%len(line)] {
				totalP2 += int(r) - int('0')
			}
		}
		fmt.Println("Part1", totalP1)
		fmt.Println("Part2", totalP2)
	}
}
