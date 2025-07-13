package easy

import "math"

func ReverseBits(num uint32) uint32 {
    var cnt uint32 = 0
    for i := 31; i >= 0; i-- {
        lastBit := num & 1
        if lastBit == 1 {
            cnt += uint32(math.Pow(2, float64(i)))
        }
        num >>= 1
    }
    return cnt
}