/*
Given an integer x, return true if x is a
palindrome
, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/
package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return x == reversed || x == reversed/10
}
func main() {
	fmt.Println(isPalindrome(121))  // Output: true
	fmt.Println(isPalindrome(-121)) // Output: false
	fmt.Println(isPalindrome(10))   // Output: false
}
