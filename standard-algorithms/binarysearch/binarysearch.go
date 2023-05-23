package binarysearch

import (
	"errors"
	"fmt"
)

// Given a slice of ints, finds the index of the target value
// Assumes the input slice is sorted
func BinarySearch(nums []int, target int) (int, error) {
	var index int
	bottom := 0
	top := len(nums) - 1

	for bottom <= top {
		index = bottom + ((top - bottom) / 2)
		fmt.Printf("Window %d,%d,%d - nums[%d]=%d\n", bottom, index, top, index, nums[index])

		if nums[index] == target {
			return index, nil
		} else if nums[index] > target {
			top = index - 1
		} else {
			bottom = index + 1
		}
	}

	return -1, errors.New("target not found")
}
