package main

import "fmt"

func main() {
	n := 9
	sequenceNumber(n)
}

func sequenceNumber(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i%2 == 0 && j%2 != 0) || (i%2 != 0 && j%2 == 0) {
				fmt.Printf("%d ", j+1)
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}
