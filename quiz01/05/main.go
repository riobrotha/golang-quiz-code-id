package main

import "fmt"


func main() {
	n := 9
	sequenceNumber(n)
}

func sequenceNumber(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				fmt.Printf("%d ", i+1)
			} else {
				fmt.Printf("%d ", n-i)
			}
		}
		fmt.Println()
	}
}
