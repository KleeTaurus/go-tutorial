package algorithms

func Countdown(n int, target []int) {
	target = append(target, n)
	if n <= 0 {
		return
	}
	Countdown(n-1, target)
}
