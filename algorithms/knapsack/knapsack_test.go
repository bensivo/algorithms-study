package knapsack_test

import (
	"testing"

	"bensivo.com/leetcode/algorithms/knapsack"
	"github.com/stretchr/testify/assert"
)

func TestKnapsack_Empty(t *testing.T) {
	weights := []int{}
	values := []int{}
	capacity := 5
	assert.Equal(
		t,
		0,
		knapsack.Knapsack(capacity, weights, values),
	)
}
func TestKnapsack_OneItem(t *testing.T) {
	weights := []int{5}
	values := []int{1}
	capacity := 5

	assert.Equal(
		t,
		1,
		knapsack.Knapsack(capacity, weights, values),
	)
}
func TestKnapsack_Decision(t *testing.T) {
	weights := []int{5, 5}
	values := []int{1, 2}
	capacity := 5

	assert.Equal(
		t,
		2,
		knapsack.Knapsack(capacity, weights, values),
	)
}

func TestKnapsack_Complex(t *testing.T) {
	weights := []int{12, 32, 6, 12, 1, 64, 12, 32, 5, 61}
	values := []int{8, 111, 34, 4525, 25, 321, 46, 3, 223, 43}
	capacity := 100

	value := knapsack.Knapsack(capacity, weights, values)

	assert.Equal(t, 5174, value)
}
