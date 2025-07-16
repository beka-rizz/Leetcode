package easy

func ContainsDuplicate(nums []int) bool {
    mp := make(map[int]int)
    for _, num := range nums {
        mp[num]++
        if mp[num] > 1 {
            return true
        }
    }
    return false
}