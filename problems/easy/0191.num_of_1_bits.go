package easy

func HammingWeight(n int) int {
    cnt := 0
    for n > 0 {
        rem := n % 2
        if rem == 1 {
            cnt++
        }
        n /= 2
    }
    return cnt 
}