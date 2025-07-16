package easy

func IsUgly(n int) bool {
	if n == 1 {
		return true
	}
	for i := 2; i <= n; i *= 2 {
		for j := 3; j <= n; j *= 3 {
			for k := 5; k <= n; k *= 5 {
				if i*j*k == n || j*k == n || k == n || i*k == n {
					return true
				}
			}
			if i*j == n || j == n {
				return true
			}
		}
		if i == n {
			return true
		}
	}
	return false
}
