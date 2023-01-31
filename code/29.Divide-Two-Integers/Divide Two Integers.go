package code

import "math"

func divide(dividend int, divisor int) int {
	sign, res := -1, 0
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}
	if dividend > math.MaxInt32 {
		dividend = math.MaxInt32
	}
	res = binarySearchQuotient(0, abs(dividend), abs(divisor), abs(dividend))
	if res > math.MaxInt32 {
		return sign * math.MaxInt32
	}
	if res < math.MinInt32 {
		return sign * math.MinInt32
	}
	return sign * res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func binarySearchQuotient(low, high, val, dividend int) int {
	quotient := low + ((high - low) >> 1)
	if ((quotient+1)*val > dividend && quotient*val <= dividend) || ((quotient+1)*val >= dividend && quotient*val < dividend) {
		if (quotient+1)*val == dividend {
			return quotient + 1
		}
		return quotient
	}
	if (quotient+1)*val > dividend && quotient*val > dividend {
		return binarySearchQuotient(low, quotient-1, val, dividend)
	}
	if (quotient+1)*val < dividend && quotient*val < dividend {
		return binarySearchQuotient(quotient+1, high, val, dividend)
	}
	return 0
}
