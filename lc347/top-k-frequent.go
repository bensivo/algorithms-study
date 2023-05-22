package lc347

import (
	"sort"
)

type NumCount struct {
	Num   int
	Count int
}

// First solution - 38% on runtime and 91% on memory
func TopKFrequent(nums []int, k int) []int {
	// Create a map, from value to count of value
	counts := make(map[int]int)
	for _, v := range nums {
		count, exists := counts[v]
		if !exists {
			count = 0
		}
		counts[v] = count + 1
	}

	// Turn that map into an array of structs
	//
	// POTENTIAL IMPROVEMENT -  Use a maxHeap or a priority queue when converting from the map to the arr of structs
	// 	This will save you the effort of sorting the array after.
	numCounts := make([]NumCount, len(counts))
	cnt := 0
	for k, v := range counts {
		numCounts[cnt] = NumCount{
			Num:   k,
			Count: v,
		}
		cnt++
	}

	// Sort the struct arr by the "Count" field, in reverse order, so max counts is first
	sort.Slice(numCounts, func(a int, b int) bool {
		return numCounts[b].Count < numCounts[a].Count
	})

	// Take the first k numbers into your result
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = numCounts[i].Num
	}
	return res
}
