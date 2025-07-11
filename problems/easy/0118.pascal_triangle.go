package easy

func GeneratePT(numRows int) {
	res := make([][]int, numRows)
    for i := range res {
		for j := 0; j <= i; j++ {
			// last element of row
			if j == i || j == 0 {
				res[i] = append(res[i], 1)
			} else {
				res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
			}
		}
    }
}