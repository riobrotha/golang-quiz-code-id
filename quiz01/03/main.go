package main

import "fmt"

func main() {
	n := 5
	triAngle1(n)
	fmt.Println("============================")
	triAngle2(n)
}

func triAngle1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func triAngle2(n int) {
	limit := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= limit {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
		limit--
	}
}
