package tools

type Pair struct {
	I int
	J int
}

type Grid[T any] struct {
	Array   []T
	Rows    int
	Columns int
}

func NewGrid[T any](rows int, columns int) *Grid[T] {
	return &Grid[T]{
		Array:   make([]T, rows*columns),
		Rows:    rows,
		Columns: columns,
	}
}

func (g *Grid[T]) Index(p Pair) int {
	row := (((p.I % g.Rows) + g.Rows) % g.Rows)
	column := (((p.J % g.Columns) + g.Columns) % g.Columns)
	return row*g.Rows + column
}

func (g *Grid[T]) Get(p Pair) T {
	return g.Array[g.Index(p)]
}

func (g *Grid[T]) Set(p Pair, value T) {
	g.Array[g.Index(p)] = value
}

func (g *Grid[T]) Neighbors(p Pair, orthogonal bool, diagonal bool, periodic bool) []Pair {
	atTop := p.I == 0
	atBottom := p.I == (g.Rows - 1)
	atLeft := p.J == 0
	atRight := p.J == (g.Columns - 1)

	values := []Pair{}

	if !atTop || (atTop && periodic) {
		if orthogonal {
			values = append(values, Pair{I: p.I - 1, J: p.J}) // Up
		}
		if diagonal && (!atLeft || (atLeft && periodic)) {
			values = append(values, Pair{I: p.I - 1, J: p.J - 1}) // Up Left
		}
		if diagonal && (!atRight || (atRight && periodic)) {
			values = append(values, Pair{I: p.I - 1, J: p.J + 1}) // Up Right
		}
	}

	if orthogonal && (!atLeft || (atLeft && periodic)) {
		values = append(values, Pair{I: p.I, J: p.J - 1}) // Left
	}

	if orthogonal && (!atRight || (atRight && periodic)) {
		values = append(values, Pair{I: p.I, J: p.J + 1}) // Right
	}

	if !atBottom || (atBottom && periodic) {
		if orthogonal {
			values = append(values, Pair{I: p.I + 1, J: p.J}) // Down
		}
		if diagonal && (!atLeft || (atLeft && periodic)) {
			values = append(values, Pair{I: p.I + 1, J: p.J - 1}) // Down Left
		}
		if diagonal && (!atRight || (atRight && periodic)) {
			values = append(values, Pair{I: p.I + 1, J: p.J + 1}) // Down Right
		}
	}

	return values
}
