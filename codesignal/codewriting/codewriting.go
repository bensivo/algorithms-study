package codewriting

import (
	"fmt"
	"strconv"
)

func CodeWriting(message string) int {

	if message == "" {
		return 1
	}

	modulo := 1000000007

	// Interpretation: combos[n][0] = number of combinations with n chars, which end in a single character encoding
	//                 combos[n][1] = number of combinations with n chars, which end in a double character encoding
	combos := make([][]int, len(message)+1)
	for n := 0; n < len(message)+1; n++ {
		combos[n] = []int{0, 0}
	}

	// NOTE: we start at n=1 because the 0 row is all 0's anways
	for n := 1; n < len(message)+1; n++ {
		r := message[n-1]

		// Find how many combinations there are that end in a single-character
		if r != '0' {
			if n == 1 { // Base case, first letter
				combos[n][0] = 1
			} else {
				combos[n][0] = (combos[n-1][0] + combos[n-1][1]) % modulo // Take all previous combos, and add a letter to the end
			}
		} else {
			combos[n][0] = 0
		}

		// Find how many combinations we can make that end in a double-character
		if n >= 2 {
			parsed, _ := strconv.ParseInt(fmt.Sprintf("%c%c", message[n-2], message[n-1]), 10, 0)
			if parsed >= 10 && parsed <= 26 {
				// Take all combos from n-1 which ended in a single character and turn them into our double character
				combos[n][1] = combos[n-1][0]
			} else {
				combos[n][1] = 0
			}
		}

		// fmt.Printf("%s - %d, %d\n", message[:n], combos[n][0], combos[n][1])
	}

	return combos[len(message)][0] + combos[len(message)][1]
}

// 4 character combinations
// i i i i
// ii i i
// i ii i
// i i ii
// ii ii

// Adding the 5th character
// - all the options from before, with just a single added at the end
// i i i i i
// ii i i i
// i ii i i
// i i ii i
// ii ii i
//
// - all the options from before, but anything that ends in a single, becomes a double
// i i i ii
// ii i ii
// i ii ii
//
//  can only happen if the double is valid - 2 letters are somewhere in 10-26
