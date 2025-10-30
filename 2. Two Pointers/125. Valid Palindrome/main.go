package main

func isPalindrome(s string) bool {
	b := make([]byte, 0)

	for i := range len(s) {
		if s[i] >= 'A' && s[i] <= 'Z' {
			b = append(b, s[i]+32)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			b = append(b, s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			b = append(b, s[i])
		}
	}

	i, j := 0, len(b)-1
	for {
		if i >= j {
			break
		}
		if b[i] != b[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true

}
