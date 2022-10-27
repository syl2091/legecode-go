package code

func twoSum(ints []int, target int) []int {
	m := make(map[int]int)
	for k, v := range ints {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
