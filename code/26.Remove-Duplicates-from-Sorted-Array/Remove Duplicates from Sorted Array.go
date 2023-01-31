package code

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
