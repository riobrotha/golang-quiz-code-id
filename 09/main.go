package main

import "fmt"

func main() {
	fmt.Print(reverse("XYnb"))
}

func reverse(s string) string {
	sRune := []rune(s)
	left := 0
	right := len(sRune) - 1
	for left < right {
		sRune[left], sRune[right] = sRune[right], sRune[left]
		left++
		right--
	}
	return string(sRune)
}