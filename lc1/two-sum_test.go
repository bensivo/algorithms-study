package lc1_test

import (
	"testing"

	"bensivo.com/leetcode/lc1"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	assert.Equal(
		t,
		lc1.TwoSum([]int{2, 7, 11, 15}, 9),
		[]int{0, 1},
		"Example 1",
	)

	assert.Equal(
		t,
		lc1.TwoSum([]int{3, 2, 4}, 6),
		[]int{1, 2},
		"Example 2",
	)

	assert.Equal(
		t,
		lc1.TwoSum([]int{3, 3}, 6),
		[]int{0, 1},
		"Example 3",
	)
}
