package lc70

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
func ClimbStairs(n int) int {

	steps := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			steps[i] = 1
			continue
		}
		if i == 1 {
			steps[i] = 2
			continue
		}

		// 2 ways to get to step i, go to i-2, and take 2 steps. Or go to i-1 and take 1 step.
		steps[i] = steps[i-2] + steps[i-1]
	}

	return steps[n-1]
}
