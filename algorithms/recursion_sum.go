package main

func recursionSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + recursionSum(nums[1:])
}
