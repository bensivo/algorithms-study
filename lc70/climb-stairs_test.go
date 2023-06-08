package lc70_test

import (
	"testing"

	"bensivo.com/leetcode/lc70"
	"github.com/stretchr/testify/assert"
)

func TestClimbStairs_Example1(t *testing.T) {
	assert.Equal(t, 2, lc70.ClimbStairs(2))
}

func TestClimbStairs_Example2(t *testing.T) {
	assert.Equal(t, 3, lc70.ClimbStairs(3))
}

func TestClimbStairs_Example3(t *testing.T) {
	// 1
	// 1 1, 2
	// 1 2, 1 1 1, 2 1
	// 1 1 2, 2 2, 1 2 1, 1 1 1 1, 2 1 1
	assert.Equal(t, 5, lc70.ClimbStairs(4))
}
