package days

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	filename := filepath.Join("..", "puzzles", "test.txt")
	lines := ReadFileLines(filename)
	if !reflect.DeepEqual(lines, []string{"line1", "line2", "line3"}) {
		t.Errorf("expected list of lines did not match")
	}
}

func BenchmarkReadFileLines(b *testing.B) {
	filename := filepath.Join("..", "puzzles", "test.txt")
	for i := 0; i < b.N; i++ {
		ReadFileLines(filename)
	}
}
