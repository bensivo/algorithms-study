package lc242

// First solution
func IsAnagram_0(s string, t string) bool {
	runeMap := make(map[rune]uint16)

	// Iterate through first string, counting occurrences of each character
	for _, r := range s {
		count, exists := runeMap[r]
		if !exists {
			count = 0
		}

		runeMap[r] = count + 1
	}

	// Iterate through second string, decrementing the counts.
	// Fail if a char was not in the original count, or decrementing brings us below 0.
	for _, r := range t {
		count, exists := runeMap[r]
		if !exists {
			return false
		}

		if count == 0 {
			return false
		}

		runeMap[r] = count - 1
	}

	// One final iteration, see that there are no leftover characters in the count
	for _, v := range runeMap {
		if v != 0 {
			return false
		}
	}

	return true
}

// Second solution
// Adds basic dummy checks, and pre-initializes the array to 26 , because we know all letters are lowercase english letters
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make([]int, 26, 26)

	// Iterate through first string, counting occurrences of each character
	for _, r := range s {
		counts[r-'a']++ // where r-'a' turns into a number between 0 and 25
	}

	// Iterate through second string, decrementing the counts.
	// Fail if a char was not in the original count, or decrementing brings us below 0.
	for _, r := range t {
		counts[r-'a']--

	}

	// One final iteration, see that there are no leftover characters in the count
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
