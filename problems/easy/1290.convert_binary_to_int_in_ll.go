package easy

import "math"


type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	var exps []int
	res := 0
	idx := 0
	for head != nil {
		if head.Val == 1 {
			exps = append(exps, idx)
		}
		idx++
		head = head.Next
	}
	for _, el := range exps {
		res += int(math.Pow(2, float64(idx-1-el)))
	}
	return res
}