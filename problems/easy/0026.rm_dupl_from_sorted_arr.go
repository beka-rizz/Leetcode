package easy

func RemoveDuplicatesFromSortedArray(nums []int) int {
    initSize := len(nums)
    k := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
    }
    k = len(nums)
    for range initSize-k {
        nums = append(nums, -1)
    }
    return k
}