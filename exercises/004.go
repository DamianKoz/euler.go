package exercises

import (
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func Solve004() int {
	for i := 999; i >= 0; i-- {
		for j := 999; j >= 800; j-- {
			sum := i * j
			if check_palindrome(sum) {
				return sum
			}
		}
	}
	return 1
}

func check_palindrome(intNum int) bool {
	reversedString := ""
	num := strconv.Itoa(intNum)
	for i := len(num) - 1; i >= 0; i-- {
		reversedString += string(num[i])
	}
	return num == reversedString
}
