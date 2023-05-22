package lc242_test

import (
	"testing"

	"bensivo.com/leetcode/lc242"
	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(
		t,
		lc242.IsAnagram("anagram", "anagram"),
		true,
		"should be true",
	)

	assert.Equal(
		t,
		lc242.IsAnagram("anagram", "naagram"),
		true,
		"should be true",
	)
	assert.Equal(
		t,
		lc242.IsAnagram("a", "b"),
		false,
		"should be false",
	)
}
