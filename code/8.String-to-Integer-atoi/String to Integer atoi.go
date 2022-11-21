package code

import "math"

func myAtoi(s string) int {
	signAllowed, whitespaceAllowed, sign, digits, res := true, true, 1, []int{}, 0
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whitespaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}
		whitespaceAllowed, signAllowed = false, false
		digits = append(digits, int(c-48))
	}
	j := 1
	for i := len(digits); i > 0; i-- {
		res += int(math.Pow(10, float64(j-1))) * digits[i-1]
		j++
	}
	return res * sign
}
