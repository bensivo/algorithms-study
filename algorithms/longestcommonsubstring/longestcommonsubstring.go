package longestcommonsubstring

import "fmt"

// Given 2 strings, find the longest common substring among them
//
// Basic dynamic programming algorithm
func LongestCommonSubstring(a string, b string) string {
	// DP solution
	// Starting with 2 empty strings, add one character to a and b, and track the length of the common suffixes
	//
	// Given the LCSuffix of "abc" and "xbc" is len 2, if we add a "d" to both strings, the new LCSuffix is len 3.
	lengths := make([][]int, len(a)+1) // Where lengths[i][j] means "length of LCSuffix if a stopped at i and b stoped at j"
	for i := 0; i < len(a)+1; i++ {
		lengths[i] = make([]int, len(b)+1)
	}

	maxLen := 0 // length of LCS
	maxI := 0   // index of last letter LCS (in string a)

	for i := 0; i < len(a)+1; i++ {
		for j := 0; j < len(b)+1; j++ {
			if i == 0 || j == 0 {
				// If any of the strings is len 0, len of LCS is 0
				lengths[i][j] = 0
				continue
			}

			if a[i-1] == b[j-1] {
				lengths[i][j] = lengths[i-1][j-1] + 1

				if lengths[i][j] > maxLen {
					maxLen = lengths[i][j]
					maxI = i
				}
			} else {
				lengths[i][j] = 0
			}
		}
	}

	fmt.Printf("Len of max substring %d\n", maxLen)
	fmt.Printf("LCS between %s and %s: %s\n", a, b, a[maxI-maxLen:maxI])
	return a[maxI-maxLen : maxI]
}
