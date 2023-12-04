package tools

import (
	"reflect"
	"testing"
)

func TestGridIndexPeriodic(t *testing.T) {
	grid := NewGrid[int](3, 3)
	p := Pair{I: 0, J: 0}
	if grid.Index(p) != 0 { // 0, 0
		t.Errorf("0, 0 expected 0, got %v", p)
	}
	p = Pair{I: -1, J: -1}
	if grid.Index(p) != 8 { // 2, 2
		t.Errorf("-1, -1 expected 8, got %v", grid.Index(p))
	}
	p = Pair{I: -1, J: 0}
	if grid.Index(p) != 6 { // 2, 0
		t.Errorf("-1, 0 expected 6, got %v", grid.Index(p))
	}
	p = Pair{I: 0, J: -1}
	if grid.Index(p) != 2 { // 0, 2
		t.Errorf("0, -1 expected 2, got %v", grid.Index(p))
	}
	p = Pair{I: 3, J: 3}
	if grid.Index(p) != 0 { // 0, 0
		t.Errorf("3, 3 expected 0, got %v", grid.Index(p))
	}
	p = Pair{I: 3, J: 0}
	if grid.Index(p) != 0 { // 0, 0
		t.Errorf("3, 0 expected 0, got %v", grid.Index(p))
	}
	p = Pair{I: 0, J: 3}
	if grid.Index(p) != 0 { // 0, 0
		t.Errorf("0, 3 expected 0, got %v", grid.Index(p))
	}
	p = Pair{I: -1, J: 3}
	if grid.Index(p) != 6 { // 2, 0
		t.Errorf("2, 0 expected 6, got %v", grid.Index(p))
	}
	p = Pair{I: 3, J: -1}
	if grid.Index(p) != 2 { // 0, 2
		t.Errorf("0, 2 expected 2, got %v", grid.Index(p))
	}
}

func TestGridSetGet(t *testing.T) {
	grid := NewGrid[int](3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid.Set(Pair{I: i, J: j}, i*3+j)
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			value := grid.Get(Pair{I: i, J: j})
			expectedValue := i*3 + j
			if value != expectedValue {
				t.Errorf("Grid item: %v, %v expected %v got %v", i, j, value, expectedValue)
			}
		}
	}
}

func TestGridGetNeighbors(t *testing.T) {
	grid := NewGrid[int](3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid.Set(Pair{I: i, J: j}, i*3+j)
		}
	}

	if !reflect.DeepEqual(
		grid.Neighbors(Pair{I: 0, J: 0}, true, true, true),
		[]Pair{
			{I: -1, J: 0},
			{I: -1, J: -1},
			{I: -1, J: 1},
			{I: 0, J: -1},
			{I: 0, J: 1},
			{I: 1, J: 0},
			{I: 1, J: -1},
			{I: 1, J: 1},
		},
	) {
		t.Errorf("Missing or incorrect pair in neighbors (0, 0) diagonal, orthogonal, and periodic expected")
	}

	if !reflect.DeepEqual(
		grid.Neighbors(Pair{I: 0, J: 0}, true, false, true),
		[]Pair{
			{I: -1, J: 0},
			{I: 0, J: -1},
			{I: 0, J: 1},
			{I: 1, J: 0},
		},
	) {
		t.Errorf("Missing or incorrect pair in neighbors (0, 0) not diagonal, orthogonal, and periodic expected")
	}

	if !reflect.DeepEqual(
		grid.Neighbors(Pair{I: 0, J: 0}, false, true, true),
		[]Pair{
			{I: -1, J: -1},
			{I: -1, J: 1},
			{I: 1, J: -1},
			{I: 1, J: 1},
		},
	) {
		t.Errorf("Missing or incorrect pair in neighbors (0, 0) diagonal, not orthogonal, and periodic expected")
	}

	if !reflect.DeepEqual(
		grid.Neighbors(Pair{I: 0, J: 0}, true, true, false),
		[]Pair{
			{I: 0, J: 1},
			{I: 1, J: 0},
			{I: 1, J: 1},
		},
	) {
		t.Errorf("Missing or incorrect pair in neighbors (0, 0) diagonal, orthogonal, and not periodic expected")
	}
}
