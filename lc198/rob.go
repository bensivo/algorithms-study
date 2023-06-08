package lc198

func Rob(nums []int) int {

	// Basic idea, start with 1 house, and add houses until you get to n
	// At each step i, you can either not rob that house, and take the value of n-1. Or robn that house and take n-2 + value

	values := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			values[i] = nums[i]
			continue
		}
		if i == 1 {
			values[i] = max(nums[0], nums[1])
			continue
		}

		values[i] = max(values[i-2]+nums[i], values[i-1])
	}

	return values[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
