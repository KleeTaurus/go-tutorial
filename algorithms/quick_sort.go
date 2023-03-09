package main

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[len(nums)/2]
	lower := []int{}
	higher := []int{}
	for i, x := range nums {
		if i == len(nums)/2 {
			continue
		}
		if x <= pivot {
			lower = append(lower, x)
		} else {
			higher = append(higher, x)
		}
	}

	target := append(quickSort(lower), pivot)
	target = append(target, quickSort(higher)...)
	return target
}
