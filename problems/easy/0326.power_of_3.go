package easy

func IsPowerOfThree(n int) bool {
    if n < 0 {
        return false
    }

    for i := 1; i <= n; i*=3 {
        if i == n {
            return true
        }
    }
    return false
}