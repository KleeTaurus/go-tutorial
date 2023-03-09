package main

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := nums[len(nums)/2]
	lower := []int{}
	higher := []int{}
	for i, x := range nums {
		if i == len(nums)/2 {
			continue
		}
		if x <= mid {
			lower = append(lower, x)
		} else {
			higher = append(higher, x)
		}
	}

	target := []int{}
	target = append(target, quickSort(lower)...)
	target = append(target, mid)
	target = append(target, quickSort(higher)...)
	return target
}
