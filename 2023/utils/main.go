package utils

func MinWithBound(i int, min int) int {
	if i > min {
		return i - 1
	}

	return min
}
func MaxWithBound(i int, max int) int {
	if i < max {
		return i + 1
	}

	return max
}
