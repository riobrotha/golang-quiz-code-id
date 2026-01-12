package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input Jumlah Baris Pirmaid : ")
	fmt.Scan(&n)
	pyramidNumber(n)
}

func pyramidNumber(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < (n*2) - 1; j++ {
			result := (n-1-j-i)+1
			if result < 1 {
				result = (result * -1) + 2
			}
			if result > n - i {
				fmt.Print(" ")
			}else {
				fmt.Printf("%d ", result)
			}
		}

		fmt.Println()
	}
}