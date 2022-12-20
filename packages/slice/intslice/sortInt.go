package intslice

import "sort"

func SortInt(a []int) []int {
	s := make([]int, len(a))
	copy(s, a)
	sort.Ints(s)
	return s
}
