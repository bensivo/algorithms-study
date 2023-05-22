package lc1

func TwoSum_0(nums []int, target int) []int {

	for a := 0; a < len(nums)-1; a++ {
		for b := a + 1; b < len(nums); b++ {
			if nums[a]+nums[b] == target {
				return []int{a, b}
			}
		}
	}

	return nil
}

// Second solution, uses a map to track expected values
func TwoSum(nums []int, target int) []int {

	// Key = diff, Value = Index such that diff + nums[index] = target
	diffMap := make(map[int]int)

	for index, value := range nums {
		if diffIndex, exists := diffMap[value]; exists {
			return []int{diffIndex, index}
		}

		diff := target - value
		diffMap[diff] = index
	}

	return nil
}
