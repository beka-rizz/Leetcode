package easy

// The idea is to use bit manipulation based on the observation that if n is a power of 2, then its binary representation has exactly one bit set to 1, and n-1 will have all bits to the right of that bit set to 1.
// Therefore, n & (n-1) will always be 0 for powers of 2.

func IsPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}
