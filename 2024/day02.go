package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validSeries(d []int) bool {
	if len(d) < 2 {
		return false
	}

	if (d[0] - d[1]) > 0 {
		for i := range len(d) - 1 {
			diff := d[i] - d[i+1]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	} else {
		for i := range len(d) - 1 {
			diff := d[i] - d[i+1]
			if diff < -3 || diff > -1 {
				return false
			}
		}
	}

	return true
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// Index out of range, return a copy of the original slice
		return append([]int(nil), slice...)
	}
	// Create a new slice without the element at the given index
	return append(append([]int(nil), slice[:index]...), slice[index+1:]...)
}

func main() {
	file, err := os.Open("inputs/day02.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numValid := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)

		line := scanner.Text()
		tokens := strings.Split(line, " ")

		for _, token := range tokens {
			i, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Unable to parse token %v", err.Error())
				return
			}

			row = append(row, i)
		}

		if validSeries(row) {
			numValid = numValid + 1
		} else {
			for i := range len(row) {
				if validSeries(removeElement(row, i)) {
					numValid = numValid + 1
					break
				}
			}
		}
	}

	fmt.Println("part1", numValid)
}
