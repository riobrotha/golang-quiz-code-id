package main

import "fmt"

func main() {
	findDivisor(7)
}

func findDivisor(n int) {
	for i := 0; i < n; i++ {
		if n%(i+1) == 0 && i < n-1 {
			fmt.Printf("%d ", i+1)
		}
	}
}