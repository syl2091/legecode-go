package code

func longestCommonPrefix(strings []string) string {
	prefix := strings[0]
	for i := 1; i < len(strings); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strings[i]) <= j || strings[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}
