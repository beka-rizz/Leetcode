package easy

// Brian Kernighan's Algorithm

func CountBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		j := i
		cnt := 0
		for j > 0 {
			j = j & (j - 1)
			cnt++
		}
		res[i] = cnt
	}
	return res
}
