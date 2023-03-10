package main

func convert(nums []int) [][]int {
	length := len(nums)
	table := make([][]int, length, length)
	for i, item := range nums {
		table[i] = make([]int, length, length)
		for j, num := range nums {
			table[i][j] = item * num
		}
	}

	return table
}
