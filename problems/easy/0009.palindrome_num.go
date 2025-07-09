package problems

import "fmt"

func IsPalindrome(x int) bool {
    if x < 0 {
		return false
	} else {
		numString := fmt.Sprint(x)
		for i, el := range numString {
			mirror_el := numString[len(numString)-1-i]
			if  mirror_el != byte(el) {
				return false
			}
		}
	}
	return true
}