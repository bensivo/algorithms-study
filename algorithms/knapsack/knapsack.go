package knapsack

// DP implementation of Knapsack, only returnign the max value achievable
func Knapsack(capacity int, weights []int, values []int) int {

	// Initialize DP matrix, capacity + 1 x items + 1
	maxValue := make([][]int, capacity+1)
	for c := 0; c < capacity+1; c++ {
		maxValue[c] = make([]int, len(weights)+1)
	}

	for c := 0; c < capacity+1; c++ {
		for n := 0; n < len(weights)+1; n++ {
			if c == 0 || n == 0 {
				maxValue[c][n] = 0
				continue
			}

			ind := n - 1 // Index of this item in the original weights, values arrays
			weight := weights[ind]
			value := values[ind]

			// Item won't fit in the bag, it might as well not be an option
			// Here, we just copy the value from the previous n iteration
			if weight > c {
				maxValue[c][n] = maxValue[c][n-1]
				continue
			}

			// See new value if we included the item, using the reduced capacity value
			valueIfIncluded := maxValue[c-weight][n-1] + value

			// See new value if we don't include the item, which is just the last n iteration (as if this item wasn't an option)
			valueIfNotIncluded := maxValue[c][n-1]

			if valueIfIncluded > valueIfNotIncluded {
				maxValue[c][n] = valueIfIncluded
			} else {
				maxValue[c][n] = valueIfNotIncluded
			}
		}
	}

	return maxValue[capacity][len(weights)]
}
