package lc49_test

import (
	"testing"

	"bensivo.com/leetcode/lc49"
	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {

	assert.Equal(
		t,
		lc49.GroupAnagrams([]string{
			"eat", "tea", "tan", "ate", "nat", "bat",
		}),
		[][]string{ // NOTE: this test fails if the order is wrong, but the actual Leetcode problem should pass
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		},
		"Example 1",
	)

	assert.Equal(
		t,
		lc49.GroupAnagrams([]string{
			"",
		}),
		[][]string{
			{""},
		},
		"Example 2",
	)

	assert.Equal(
		t,
		lc49.GroupAnagrams([]string{
			"a",
		}),
		[][]string{
			{"a"},
		},
		"Example 3",
	)
}
