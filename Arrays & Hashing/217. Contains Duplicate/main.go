package main

func containsDuplicate(nums []int) bool {
	n := len(nums)
	m := make(map[int]struct{}, n)
	v := struct{}{}
	for i := range nums {
		num := nums[i]
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = v
	}

	return false
}
