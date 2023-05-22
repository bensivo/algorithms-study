package lc125

import (
	"unicode"
)

func IsPalindrome_0(s string) bool {
	alphanumeric := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			alphanumeric = append(alphanumeric, unicode.ToLower(r))
		}
	}

	i := 0
	j := len(alphanumeric) - 1
	for {
		if i > j {
			break
		}

		if alphanumeric[i] != alphanumeric[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// Second solution, doesn't create a new representation fo the string in memory at all, just iterates on the original
func IsPalindrome(s string) bool {

	i := 0
	j := len(s) - 1
	for {
		if i > j {
			break
		}

		rune_i := rune(s[i])
		if !unicode.IsDigit(rune_i) && !unicode.IsLetter(rune_i) {
			i++
			continue
		}

		rune_j := rune(s[j])
		if !unicode.IsDigit(rune_j) && !unicode.IsLetter(rune_j) {
			j--
			continue
		}

		if unicode.ToLower(rune_i) == unicode.ToLower(rune_j) {
			i++
			j--
			continue
		}

		return false
	}

	return true
}
