package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	left := make([]int, 0)
	right := make([]int, 0)

	leftSet := make(map[int]int)
	rightSet := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "   ")

		if len(tokens) != 2 {
			fmt.Printf("Error expecting 2 tokens got %d\n", len(tokens))
			return
		}

		i, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Error converting integer string %s: %w", tokens[0], err)
		}
		left = append(left, i)
		leftSet[i] = leftSet[i] + 1

		j, err := strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Printf("Error converting integer string %s: %w", tokens[1], err)
		}
		right = append(right, j)
		rightSet[j] = rightSet[j] + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum1 := 0
	for i := range left {
		sum1 = sum1 + abs(left[i]-right[i])
	}
	fmt.Println(sum1)

	sum2 := 0
	for _, j := range left {
		// if i doesn't exist it will be zero the default value
		sum2 = sum2 + j*rightSet[j]
	}
	fmt.Println(sum2)
}
