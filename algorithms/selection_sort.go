package main

func findSmallest(items []int) int {
	if len(items) == 0 {
		return -1
	}

	sidx := 0
	for idx, item := range items {
		if item < items[sidx] {
			sidx = idx
		}
	}
	return sidx
}

func selectionSort(items []int) []int {
	target := []int{}

	for len(items) > 0 {
		sidx := findSmallest(items)
		if sidx == -1 {
			break
		}
		target = append(target, items[sidx])
		if sidx+1 == len(items) {
			items = items[:sidx]
		} else {
			items = append(items[:sidx], items[sidx+1:]...)
		}
	}
	return target
}
