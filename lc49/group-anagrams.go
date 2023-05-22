package lc49

import (
	"strconv"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	buckets := make(map[string][]string)

	for _, str := range strs {
		key := anagramKey(str)
		bucket, exists := buckets[key]
		if !exists {
			bucket = []string{}
		}

		buckets[key] = append(bucket, str)
	}

	res := make([][]string, len(buckets))
	index := 0
	for _, bucket := range buckets {
		res[index] = bucket
		index++
	}

	return res
}

// Convert any string into a key such that any anagrams map to the same value
//
// Format is as follows: "1,2,0,1", where each item is the number of times that letter appears
// Example: "abcdd" would map to "1,1,1,2,0,0,0,0...."
func anagramKey(str string) string {
	counts := make([]int, 26) // Potential memory improvement, use a [26]byte instead of a string as output. Fixed-length arrays can be used as hash keys in golang.  Only works if num of occurences can fit in 1 byte.
	for _, r := range str {
		counts[r-'a']++
	}

	countsStr := make([]string, 26)
	for i, v := range counts {
		countsStr[i] = strconv.Itoa(v)
	}
	return strings.Join(countsStr, ",")
}
