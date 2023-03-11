package algorithms

func RecursionSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + RecursionSum(nums[1:])
}
