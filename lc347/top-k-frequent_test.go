package lc347_test

import (
	"testing"

	"bensivo.com/leetcode/lc347"
	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	assert.Equal(
		t,
		lc347.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2),
		[]int{1, 2},
		"Example 1",
	)

	assert.Equal(
		t,
		lc347.TopKFrequent([]int{1}, 1),
		[]int{1},
		"Example 2",
	)
}
