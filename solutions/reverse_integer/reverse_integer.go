package reverse_integer

const (
	Max = (1<<31) - 1
	Min = -1<<31
)

func reverse(x int) int {
	var sign = 1

	if x < 0 {
		sign = -1
	}

	x *= sign

	var res = 0
	for x > 0 {
		res = res*10 + x%10
		if isNumWithinRange(res) {
			return 0
		}
		x /= 10
	}
	return res * sign
}

func isNumWithinRange(x int) bool {
	if x > Max || x < Min {
		return true
	} else {
		return false
	}
}