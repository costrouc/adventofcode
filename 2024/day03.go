package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TokenPattern struct {
	Label   string
	Pattern string
}

type Token struct {
	Type  string
	Value string
}

type FunctionNode struct {
	Name string
	Args []int
}

func tokenize(input string) []Token {
	tokenPatterns := []TokenPattern{
		{"FUNCTION", `[A-Za-z']+`},
		{"LPAREN", `\(`},
		{"COMMA", `,`},
		{"INTEGER", `\d+`},
		{"RPAREN", `\)`},
		{"UNKNOWN", `[^A-Za-z0-9\(\),]`},
	}

	// Compile the regex patterns into a single regex
	var combinedPattern string
	for _, pattern := range tokenPatterns {
		combinedPattern += fmt.Sprintf(`(?P<%s>%s)|`, pattern.Label, pattern.Pattern)
	}
	// Remove the trailing '|' from the combined pattern
	combinedPattern = combinedPattern[:len(combinedPattern)-1]

	re := regexp.MustCompile(combinedPattern)

	var tokens []Token
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		for i, value := range match {
			if i == 0 || value == "" {
				continue
			}
			tokens = append(tokens, Token{
				Type:  re.SubexpNames()[i],
				Value: value,
			})
		}
	}

	return tokens
}

func parse(tokens []Token) []FunctionNode {
	stmts := make([]FunctionNode, 0)

	loc := 0
	function_name := ""
	args := make([]int, 0)

	for _, token := range tokens {
		if loc == 0 && token.Type == "FUNCTION" {
			if strings.HasSuffix(token.Value, "do") {
				function_name = "do"
				loc = loc + 1
			} else if strings.HasSuffix(token.Value, "don't") {
				function_name = "don't"
				loc = loc + 1
			} else if strings.HasSuffix(token.Value, "mul") {
				function_name = "mul"
				loc = loc + 1
			}
		} else if loc == 1 && token.Type == "LPAREN" {
			loc = loc + 1
		} else if loc > 1 && token.Type == "INTEGER" {
			value, _ := strconv.Atoi(token.Value)
			args = append(args, value)
			loc = loc + 1
		} else if loc > 2 && loc%2 == 1 && token.Type == "COMMA" {
			loc = loc + 1
		} else if (loc == 2 || (loc > 2 && loc%2 == 1)) && token.Type == "RPAREN" {
			stmts = append(stmts, FunctionNode{Name: function_name, Args: args})
			function_name = ""
			args = make([]int, 0)
			loc = 0
		} else {
			function_name = ""
			args = make([]int, 0)
			loc = 0
		}
	}

	return stmts
}

func evaluate(stmts []FunctionNode, ignoreDo bool) int {
	do := true

	sum := 0
	for _, stmt := range stmts {
		if stmt.Name == "do" && len(stmt.Args) == 0 {
			do = true
		} else if stmt.Name == "don't" && len(stmt.Args) == 0 {
			do = false
		} else if stmt.Name == "mul" && len(stmt.Args) == 2 && (do || ignoreDo) {
			sum = sum + stmt.Args[0]*stmt.Args[1]
		}
	}

	return sum
}

func main() {
	file, err := os.Open("inputs/day03.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tokens := tokenize(string(input))
	stmts := parse(tokens)

	fmt.Println("part1", evaluate(stmts, true))
	fmt.Println("part2", evaluate(stmts, false))
}
