package editdistance

import "fmt"

// Edit distance using levenshtein - number of edits required to change string a into string b
// Edits being: insert, delete, substitute
func EditDistance(a string, b string) int {

	// Interpreted as - "Edit distance between a and b if a was only i characters long, and b was only j characters long"
	dist := make([][]int, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		dist[i] = make([]int, len(b)+1)
	}

	for i := 0; i < len(a)+1; i++ {
		for j := 0; j < len(b)+1; j++ {
			// Base case, if either string is empty, the distance is just the len of the other string
			if i == 0 {
				dist[i][j] = j
				continue
			}
			if j == 0 {
				dist[i][j] = i
				continue
			}

			// If both characters are the same, edit distance is equal to the edit distance of the string with both characters removed
			if a[i-1] == b[j-1] {
				dist[i][j] = dist[i-1][j-1]

				continue
			}

			// If both characters are different, take min edit distance if either char was removed, then add one.
			// Example: dist("abcdef", "def")
			// 		3 options:
			//			 - start at abcde_ + de_, and substitute the _ for an f
			//			 - start at abcde + def, and add an f to the first string
			//			 - start at abcdef + de, and add an f to the second string
			minDist := min3(dist[i-1][j], dist[i][j-1], dist[i-1][j-1])
			dist[i][j] = minDist + 1
		}
	}

	fmt.Printf("Matrix %s - %s\n", a, b)
	printMatrix(dist)

	return dist[len(a)][len(b)]
}

func printMatrix(n [][]int) {
	for _, arr := range n {
		for _, value := range arr {
			fmt.Printf("%2d ", value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func min3(a int, b int, c int) int {
	min := a
	if a < b {
		min = a
	}
	if c < min {
		min = c
	}
	return min
}
