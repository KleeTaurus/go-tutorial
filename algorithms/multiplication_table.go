package algorithms

func Convert(nums []int) [][]int {
	table := make([][]int, len(nums))
	for i, item := range nums {
		table[i] = make([]int, len(nums))
		for j, num := range nums {
			table[i][j] = item * num
		}
	}

	return table
}
