package main

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for _, n := range numbers {
		m[n] = target - n
	}

	var firstIdx, secondNumber int
	for i, n := range numbers {
		secondNumber = m[n]
		if _, ok := m[secondNumber]; ok {
			firstIdx = i + 1
			break
		}
	}

	for i, n := range numbers {
		if i <= firstIdx-1 {
			continue
		}
		if n == secondNumber {
			return []int{firstIdx, i + 1}
		}
	}

	return nil
}
