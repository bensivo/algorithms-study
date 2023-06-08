package lc198_test

import (
	"testing"

	"bensivo.com/leetcode/lc198"
	"github.com/stretchr/testify/assert"
)

func TestRob_Example1(t *testing.T) {
	assert.Equal(t, 4, lc198.Rob([]int{1, 2, 3, 1}))
}

func TestRob_Example2(t *testing.T) {
	assert.Equal(t, 12, lc198.Rob([]int{2, 7, 9, 3, 1}))
}
