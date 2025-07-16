package easy

func IsPowerOfThree2(n int) bool {
    if n <= 0 {
        return false
    }
    // max int that is power of 3 => 1_162_261_467 = 3^19
    return 1_162_261_467 % n == 0
}