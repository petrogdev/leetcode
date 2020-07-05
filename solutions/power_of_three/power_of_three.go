package power_of_three

// isIntPowerOfThree
func isIntPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}