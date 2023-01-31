package code

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2] <= target; i++ {

	}

	return nil
}
