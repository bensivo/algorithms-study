package lc125_test

import (
	"testing"

	"bensivo.com/leetcode/lc125"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(
		t,
		lc125.IsPalindrome("A man, a plan, a canal: Panama"),
		true,
		"Example 1",
	)

	assert.Equal(
		t,
		lc125.IsPalindrome("race a car"),
		false,
		"Example 2",
	)

	assert.Equal(
		t,
		lc125.IsPalindrome(" "),
		true,
		"Example 3",
	)
}
