package code

// 无重复字符的最长子串

//解法一 位图

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

//解法二 滑动窗口

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

/*
func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	index := make(map[byte]int, len(s))
	for left < len(s) {

		if idx, ok := index[s[left]]; ok && idx >= right {

		}
	}
	return 0
}*/
