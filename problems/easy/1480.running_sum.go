package easy

func RunningSum(nums []int) []int {
	res := []int{}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := 0; j <= i; j++ {
			sum += nums[j]
		}
		res = append(res, sum)
	}
	return res
}
