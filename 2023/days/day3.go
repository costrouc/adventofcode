package days

import (
	"github.com/costrouc/adventofcode/tools"
)

func isNumber(r rune) bool {
	return int(r) >= int('0') && int(r) <= int('9')
}

func isSymbol(r rune) bool {
	return !isNumber(r) && r != '.'
}

type Number struct {
	R     int
	C_i   int
	C_j   int
	Value int
}

type Symbol struct {
	R int
	C int
	S rune
}

func ParseSchematic(path string) ([]Number, []Symbol) {
	numbers := []Number{}
	symbols := []Symbol{}

	lines := tools.ReadFileLines(path)
	for i, line := range lines {
		startIndex := -1
		currentNumber := 0

		for j, r := range line {
			if isNumber(r) {
				if startIndex == -1 {
					startIndex = j
				}
				currentNumber = currentNumber*10 + int(r) - int('0')
			} else if isSymbol(r) {
				if startIndex != -1 {
					numbers = append(numbers, Number{R: i, C_i: startIndex, C_j: j - 1, Value: currentNumber})
				}
				symbols = append(symbols, Symbol{R: i, C: j, S: r})
				startIndex = -1
				currentNumber = 0
			} else { // '.'
				if startIndex != -1 {
					numbers = append(numbers, Number{R: i, C_i: startIndex, C_j: j - 1, Value: currentNumber})
				}
				startIndex = -1
				currentNumber = 0
			}
		}

		if currentNumber != 0 {
			numbers = append(numbers, Number{R: i, C_i: startIndex, C_j: len(line) - 1, Value: currentNumber})
		}
	}

	return numbers, symbols
}

func Day3Part1(path string) int {
	numbers, symbols := ParseSchematic(path)

	total := 0
	for _, number := range numbers {
		intersect := false
		for _, symbol := range symbols {
			if (symbol.R >= number.R-1 && symbol.R <= number.R+1) && (symbol.C >= number.C_i-1 && symbol.C <= number.C_j+1) {
				intersect = true
				break
			}
		}
		if intersect {
			total += number.Value
		}
	}

	return total
}

func Day3Part2(path string) int {
	numbers, symbols := ParseSchematic(path)

	total := 0
	for _, symbol := range symbols {
		intersections := 0
		gearRatio := 1
		for _, number := range numbers {
			if symbol.S == '*' && (symbol.R >= number.R-1 && symbol.R <= number.R+1) && (symbol.C >= number.C_i-1 && symbol.C <= number.C_j+1) {
				intersections += 1
				gearRatio *= number.Value
			}
		}
		if intersections == 2 {
			total += gearRatio
		}
	}

	return total
}
