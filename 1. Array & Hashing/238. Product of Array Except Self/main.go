package main

func productExceptSelf(nums []int) []int {
	n := len(nums)

	result := make([]int, n)
	l, r := 1, 1
	for i := range n {
		result[i] = l
		l *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		result[i] *= r
		r *= nums[i]
	}

	return result
}
