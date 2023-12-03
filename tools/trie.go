package tools

import (
	"io"
)

type Trie struct {
	Runes    map[rune]*Trie
	Terminal bool
	Value    int
}

func NewTrie() *Trie {
	return &Trie{Runes: make(map[rune]*Trie, 0)}
}

func (t *Trie) Set(key string, value int) {
	var ct *Trie = t

	for _, r := range key {
		_, ok := ct.Runes[r]
		if !ok {
			ct.Runes[r] = NewTrie()
		}
		ct = ct.Runes[r]
	}

	ct.Terminal = true
	ct.Value = value
}

func (t *Trie) Get(key string) int {
	ct := t
	for _, r := range key {
		ct = ct.Runes[r]
		if ct == nil {
			return -1
		}
	}
	if ct.Terminal {
		return ct.Value
	}
	return -1
}

func (t *Trie) Next(r rune) *Trie {
	return t.Runes[r]
}

// Note this function always give the shortest match
// So this function currently has issues if any word
// is a subset of the others
func (t *Trie) NextMatch(r io.RuneReader) int {
	ct := t
	for !ct.Terminal {
		r, _, err := r.ReadRune()
		if err == io.EOF {
			return -1
		} else if err != nil {
			panic(err)
		}

		ct = ct.Next(r)
		if ct == nil {
			ct = t
		}
	}

	return ct.Value
}
