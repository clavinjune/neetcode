package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = keys[i]
	}

	return result
}
