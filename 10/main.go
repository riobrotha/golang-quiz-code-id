package main

import "fmt"

func main() {
	fmt.Println(checkBraces("())"))
}

func checkBraces(s string) bool {

	open := 0
	close := 0

	for _, char := range s {
		if char == '(' {
			open++
		} else if char == ')' {
			if open == 0 {
				return false
			}
			close++
		}
	}

	return open == close
}
