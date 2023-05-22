package lc217

/**
 * Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
 */
func ContainsDuplicate(nums []int) bool {
	existsMap := make(map[int]byte)

	for _, value := range nums {
		_, ok := existsMap[value]
		if ok {
			return true
		}
		existsMap[value] = 0
	}
	return false
}
