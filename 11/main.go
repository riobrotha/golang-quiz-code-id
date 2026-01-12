package main

import "fmt"

func main() {
	fmt.Println(isNumberPalindrome(5555))
}

func isNumberPalindrome(n int) bool {

	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n = n / 10
	}

	if original == reversed {
		return true
	}
	return false
}