package main

func findCommonDivisor(a, b int) (int, bool) {
	var min, max int
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}

	if max%min == 0 {
		return min, true
	}

	for i := min / 2; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i, true
		}
	}

	return 0, false
}
