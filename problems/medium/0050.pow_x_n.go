package medium

func MyPow(x float64, n int) float64 {
	if n == 0 || x == 1.0 {
		return 1
	}

	if n <= -20 && x != -1 {
		return 0
	}

	if x == -1 {
		switch n & 1 {
		case 0:
			return 1
		case 1:
			return -1
		}
	}

	res := 1.0
	if n > 0 {
		for range n {
			res *= x
		}
		return res
	}

	k := -1 * n
	for range k {
		res *= 1 / x
	}
	return res
}
