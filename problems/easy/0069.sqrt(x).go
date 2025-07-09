package easy

func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	for i := 0; i <= x; i++ {
		if i*i > x {
			return i - 1
		} else if i*i == x {
			return i
		}
	}
	return -1
}
