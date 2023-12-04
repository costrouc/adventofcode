package tools

import "testing"

func TestSetAdd(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	if s.Cardinality() != 3 {
		t.Errorf("Expected set of size 3, got %v", s.Cardinality())
	}
}
