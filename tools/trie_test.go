package tools

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTrieSetGet(t *testing.T) {
	dictionary := map[string]int{
		"word1": 1,
		"and":   2,
		"a":     3,
	}
	trie := NewTrie()
	for word := range dictionary {
		trie.Set(word, dictionary[word])
	}

	for word := range dictionary {
		value := dictionary[word]
		if trie.Get(word) != value {
			t.Errorf("Trie get for 'and' didn't get expected value %v, got %v", value, trie.Get(word))
		}
	}
}

func TestTrieSetGetNext(t *testing.T) {
	dictionary := map[string]int{
		"word1": 1,
		"and":   2,
		"a":     3,
	}
	trie := NewTrie()
	for word, value := range dictionary {
		trie.Set(word, value)
	}

	for word, value := range dictionary {
		var ct *Trie = trie
		for _, r := range word {
			ct = ct.Next(r)
			if ct == nil {
				t.Errorf("Trie got unexpected nil for word %v", word)
			}
		}
		if ct.Value != value {
			t.Errorf("Trie get for 'and' didn't get expected value %v, got %v", value, trie.Get(word))
		}
	}
}

func TestTrieSetGetNextMatch(t *testing.T) {
	dictionary := map[string]int{
		"the": 1,
		"and": 2,
		"a":   3,
	}
	trie := NewTrie()
	for word, value := range dictionary {
		trie.Set(word, value)
	}

	line := []byte("the quick brown fox and bunny")
	reader := bytes.NewReader(line)
	for {
		token := trie.NextMatch(reader)
		if token == -1 {
			break
		}
		fmt.Println(token)
	}
}
