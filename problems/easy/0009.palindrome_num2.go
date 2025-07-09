package easy

import (
	"fmt"
)

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		var reversedString string
		var originalNumber = fmt.Sprint(x)
		for x > 0 {
			remainder := x % 10
			reversedString += fmt.Sprint(remainder)
			x /= 10
		}
		return originalNumber == reversedString
	}
}
