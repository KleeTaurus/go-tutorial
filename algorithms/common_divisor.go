package algorithms

func FindCommonDivisor(a, b int) int {
	var min, max int
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}

	mod := max % min
	if mod == 0 {
		return min
	}
	return FindCommonDivisor(min, mod)
}
