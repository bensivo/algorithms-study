package lc217_test

import (
	"testing"

	"bensivo.com/leetcode/lc217"
	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	assert.Equal(
		t,
		lc217.ContainsDuplicate([]int{1, 2, 3, 4, 5}),
		false,
		"no duplicates",
	)

	assert.Equal(
		t,
		lc217.ContainsDuplicate([]int{1, 1}),
		true,
		"duplicates",
	)

	assert.Equal(
		t,
		lc217.ContainsDuplicate([]int{}),
		false,
		"Empty - returns false",
	)
}
