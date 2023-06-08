package trie_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrie_Search(t *testing.T) {
	tri := trie.NewTrie()

	assert.Equal(t, false, tri.Search("hello"))

	tri.Insert("hello")

	assert.Equal(t, true, tri.Search("hello"))
}

func TestTrie_StartsWith(t *testing.T) {
	tri := trie.NewTrie()

	tri.Insert("hello")

	assert.Equal(t, true, tri.StartsWith("hell"))
	assert.Equal(t, false, tri.StartsWith("heli"))
}
