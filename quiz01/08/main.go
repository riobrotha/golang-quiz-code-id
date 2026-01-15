package main

import "fmt"
import "strings"

func main() {
	fmt.Println(isPalindrome("tamaT"))
}

func isPalindrome(s string) bool {
	convString := strings.ToLower(s)
	left := 0
	right := len(convString) - 1
	for left < right {
		if convString[left] != convString[right] {
			return false
		}
		left++
		right--
	}
	return true
}