package main

func countdown(n int, target []int) {
	target = append(target, n)
	if n <= 0 {
		return
	}
	countdown(n-1, target)
}
