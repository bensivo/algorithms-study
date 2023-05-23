package binarysearch_test

import (
	"testing"

	"bensivo.com/leetcode/standard-algorithms/binarysearch"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch_0(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index, _ := binarysearch.BinarySearch(nums, 4)
	assert.Equal(t, 3, index)
}

func TestBinarySearch_1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index, _ := binarysearch.BinarySearch(nums, 8)
	assert.Equal(t, 7, index)
}

func TestBinarySearch_NotFound(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	_, err := binarysearch.BinarySearch(nums, 11)
	assert.EqualError(t, err, "target not found")
}
func TestBinarySearch_Empty(t *testing.T) {
	nums := []int{}

	_, err := binarysearch.BinarySearch(nums, 1)
	assert.EqualError(t, err, "target not found")
}
