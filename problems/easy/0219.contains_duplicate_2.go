package easy

func ContainsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]int)
    for i, num := range nums {
        if val, ok := mp[num]; ok {
            if i - val <= k {
                return true
            }
        }
        mp[num] = i
    }
    return false
}