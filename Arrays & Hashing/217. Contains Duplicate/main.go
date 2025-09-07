package main

func containsDuplicate(nums []int) bool {
	n := len(nums)

	if n <= 1 {
		return false
	}

	m := make(map[int]struct{})
	v := struct{}{}

	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		ni := nums[i]
		if _, ok := m[ni]; ok {
			return true
		}
		m[ni] = v

		if i == j {
			break
		}

		nj := nums[j]
		if _, ok := m[nj]; ok {
			return true
		}
		m[nj] = v
	}

	return false
}
