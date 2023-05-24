package editdistance_test

import (
	"testing"

	"bensivo.com/leetcode/algorithms/editdistance"
	"github.com/stretchr/testify/assert"
)

func TestEditDistance_EmptyStrings(t *testing.T) {
	assert.Equal(t,
		0,
		editdistance.EditDistance("", ""),
	)
}

func TestEditDistance_Same(t *testing.T) {
	assert.Equal(t,
		0,
		editdistance.EditDistance("abcdef", "abcdef"),
	)
}

func TestEditDistance_SingleSub(t *testing.T) {
	assert.Equal(t,
		1,
		editdistance.EditDistance("aaaaaa", "aaabaa"),
	)
}

func TestEditDistance_SingleInsert(t *testing.T) {
	assert.Equal(t,
		1,
		editdistance.EditDistance("aaa", "aaaa"),
	)
}

func TestEditDistance_SingleDelete(t *testing.T) {
	assert.Equal(t,
		1,
		editdistance.EditDistance("aaa", "aa"),
	)
}

func TestEditDistance_Complex(t *testing.T) {
	assert.Equal(t,
		5,
		editdistance.EditDistance("levenshtein", "locenstdn"),
	)
}
