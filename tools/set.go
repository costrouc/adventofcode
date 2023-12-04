package tools

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T], 0)
}

func (s Set[T]) Add(t T) bool {
	prevLength := len(s)
	s[t] = struct{}{}
	return prevLength != len(s)
}

func (s Set[T]) Cardinality() int {
	return len(s)
}

func (s Set[T]) Empty() bool {
	return len(s) == 0
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		_, ok := other[k]
		if ok {
			result.Add(k)
		}
	}
	return result
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		result.Add(k)
	}
	for k := range other {
		result.Add(k)
	}
	return result
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}

	for k := range s {
		_, ok := other[k]
		if !ok {
			return false
		}
	}
	return true
}
