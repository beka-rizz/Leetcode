package easy

func TwoSum(nums []int, target int) []int {
	res := make(map[int]int)
	for index, num := range nums {
		if _, ok := res[target-num]; ok {
			return []int{res[target-num], index}
		} else {
			res[num] = index
		}
	}
	return []int{}
}
