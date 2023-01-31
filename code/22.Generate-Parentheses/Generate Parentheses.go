package code

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex int, rindex int, s string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, s)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, s+"(", res)
	}
	if rindex > 0 {
		findGenerateParenthesis(lindex, rindex-1, s+")", res)
	}
}
