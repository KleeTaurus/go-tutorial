package algorithms

func FindMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		}
		return nums[0]
	}

	max := FindMax(nums[1:])
	if nums[0] < max {
		return max
	}
	return nums[0]
}
