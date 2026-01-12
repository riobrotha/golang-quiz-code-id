package main

import "fmt"

func main() {
	n := 12234
	extractDigit(n)
}

func extractDigit(n int) {
	reversed := 0
	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		fmt.Printf("%d ", digit)
		n = n / 10
	}
}