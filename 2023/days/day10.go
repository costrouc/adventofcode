package days

import (
	"fmt"

	"github.com/costrouc/adventofcode/tools"
)

type Maze struct {
	Start tools.Pair
	Grid  *tools.Grid[rune]
}

type MazeNode struct {
	Distance int
	Point    tools.Pair
}

func ParseDay10(path string) Maze {
	lines := tools.ReadFileLines(path)
	start := tools.Pair{}

	grid := tools.NewGrid[rune](len(lines), len(lines[0]))
	for i, line := range lines {
		for j, r := range line {
			if r == 'S' {
				start = tools.Pair{I: i, J: j}
			}

			grid.Set(tools.Pair{I: i, J: j}, r)
		}
	}

	maze := Maze{Start: start, Grid: grid}

	return maze
}

func (m *Maze) String() string {
	s := ""
	for i := 0; i < m.Grid.Rows; i++ {
		for j := 0; j < m.Grid.Columns; j++ {
			s += string(m.Grid.Get(tools.Pair{I: i, J: j}))
		}
		s += "\n"
	}
	return s
}

func (m *Maze) Neighbor(node MazeNode) []MazeNode {
	currentPipe := m.Grid.Get(node.Point)
	nodes := []MazeNode{}

	switch currentPipe {
	case '|': // N-S
		nodes = []MazeNode{
			{Distance: node.Distance + 1, Point: tools.Pair{I: 1, J: 0}},
			{Distance: node.Distance + 1, Point: tools.Pair{I: -1, J: 0}},
		}
	case '-': // E-W
		nodes = []MazeNode{
			{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: 1}},
			{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: -1}},
		}
	case 'L': // N-E
		nodes = []MazeNode{
			{Distance: node.Distance + 1, Point: tools.Pair{I: -1, J: 0}},
			{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: 1}},
		}
	case 'J': // N-W
		nodes = []MazeNode{
			{Distance: node.Distance + 1, Point: tools.Pair{I: -1, J: 0}},
			{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: -1}},
		}
	case '7': // S-W
		nodes = []MazeNode{
			{Distance: node.Distance + 1, Point: tools.Pair{I: 1, J: 0}},
			{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: -1}},
		}
	case 'F': // S-E
		nodes = []MazeNode{
			{Distance: node.Distance + 1, Point: tools.Pair{I: 1, J: 0}},
			{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: 1}},
		}
	case 'S': // Start
		northPoint := m.Grid.Get(tools.Pair{I: node.Point.I - 1, J: node.Point.J})
		if northPoint == '|' || northPoint == '7' || northPoint == 'F' {
			nodes = append(nodes, MazeNode{Distance: node.Distance + 1, Point: tools.Pair{I: -1, J: 0}})
		}

		southPoint := m.Grid.Get(tools.Pair{I: node.Point.I + 1, J: node.Point.J})
		if southPoint == '|' || southPoint == 'L' || southPoint == 'J' {
			nodes = append(nodes, MazeNode{Distance: node.Distance + 1, Point: tools.Pair{I: 1, J: 0}})
		}

		westPoint := m.Grid.Get(tools.Pair{I: node.Point.I, J: node.Point.J - 1})
		if westPoint == '-' || westPoint == 'F' || westPoint == 'L' {
			nodes = append(nodes, MazeNode{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: -1}})
		}

		eastPoint := m.Grid.Get(tools.Pair{I: node.Point.I, J: node.Point.J + 1})
		if eastPoint == '-' || eastPoint == 'J' || eastPoint == '7' {
			nodes = append(nodes, MazeNode{Distance: node.Distance + 1, Point: tools.Pair{I: 0, J: 1}})
		}
	case '.': // point
		nodes = []MazeNode{}
	}

	for i := range nodes {
		nodes[i] = MazeNode{Distance: nodes[i].Distance, Point: node.Point.Add(nodes[i].Point)}
	}

	return nodes
}

func Day10Part1(path string) int {
	maze := ParseDay10(path)

	visited := make(map[tools.Pair]bool)
	visitedNodes := []MazeNode{}
	nodes := []MazeNode{{Distance: 0, Point: maze.Start}}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		if visited[node.Point] {
			continue
		}

		visited[node.Point] = true
		visitedNodes = append(visitedNodes, node)
		neighbors := maze.Neighbor(node)
		for _, neighbor := range neighbors {
			if visited[neighbor.Point] {
				continue
			}
			if neighbor.Point.I < 0 || neighbor.Point.I >= maze.Grid.Rows || neighbor.Point.J < 0 || neighbor.Point.J >= maze.Grid.Columns {
				continue
			}

			nodes = append(nodes, neighbor)
		}
	}

	maxDistance := 0
	for _, node := range visitedNodes {
		if node.Distance > maxDistance {
			maxDistance = node.Distance
		}
	}

	return maxDistance
}

func Day10Part2(path string) int {
	maze := ParseDay10(path)

	visited := make(map[tools.Pair]bool)
	visitedNodes := []MazeNode{}
	nodes := []MazeNode{{Distance: 0, Point: maze.Start}}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		if visited[node.Point] {
			continue
		}

		visited[node.Point] = true
		visitedNodes = append(visitedNodes, node)
		neighbors := maze.Neighbor(node)
		for _, neighbor := range neighbors {
			if visited[neighbor.Point] {
				continue
			}
			if neighbor.Point.I < 0 || neighbor.Point.I >= maze.Grid.Rows || neighbor.Point.J < 0 || neighbor.Point.J >= maze.Grid.Columns {
				continue
			}

			nodes = append(nodes, neighbor)
		}
	}

	insidePoints := 0
	for i := 0; i < maze.Grid.Rows; i++ {
		crossings := 0
		lastUp := 0
		for j := 0; j < maze.Grid.Columns; j++ {
			symbol := maze.Grid.Get(tools.Pair{I: i, J: j})

			_, ok := visited[tools.Pair{I: i, J: j}]

			if !ok && crossings%2 == 1 {
				fmt.Print("I")
			} else if symbol == '.' {
				fmt.Print("O")
			} else {
				fmt.Print(string(symbol))
			}

			if !ok && symbol != '.' {
				if crossings%2 == 1 {
					insidePoints++
				}
				continue
			}

			switch symbol {
			case '|': // ignore -
				crossings++
			case 'L':
				lastUp = 1
			case 'F':
				lastUp = -1
			case 'J':
				if lastUp == -1 {
					crossings++
				}
			case '7':
				if lastUp == 1 {
					crossings++
				}
			case '.':
				if crossings%2 == 1 {
					insidePoints++
				}
			case 'S': // is a '|' in this case
				crossings++
			}
		}
		fmt.Print("\n")
	}

	return insidePoints
}
