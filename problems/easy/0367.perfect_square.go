package easy

func IsPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}

	for i := 2; i <= num/2; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
