package longestcommonsubstring_test

import (
	"testing"

	"bensivo.com/leetcode/algorithms/longestcommonsubstring"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubstring(t *testing.T) {
	assert.Equal(t,
		"def",
		longestcommonsubstring.LongestCommonSubstring("abcdef", "defghi"),
	)
}

func TestLongestCommonSubstring_Empty(t *testing.T) {
	assert.Equal(t,
		"",
		longestcommonsubstring.LongestCommonSubstring("abcdef", ""),
	)
}

func TestLongestCommonSubstring_NoLCS(t *testing.T) {
	assert.Equal(t,
		"",
		longestcommonsubstring.LongestCommonSubstring("abc", "def"),
	)
}
