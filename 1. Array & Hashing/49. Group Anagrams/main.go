package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for i := range strs {
		str := strs[i]
		r := []rune(str)
		slices.Sort(r)
		rs := string(r)
		if _, ok := m[rs]; ok {
			m[rs] = append(m[rs], str)
		} else {
			m[rs] = []string{str}
		}
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}
