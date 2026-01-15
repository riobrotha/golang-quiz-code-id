package main

import "fmt"

func main() {
	var matrix [7][7]int


	matrix = initMatrix(matrix)

	displayMatrix(matrix)
}

func initMatrix(matrix [7][7]int) [7][7]int {
	counter := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[i])-1 {
				matrix[i][j] = counter
			} else {
				matrix[i][j] = 0
			}
			counter++
		}
		counter = i + 1
	}

	return matrix
}

func displayMatrix(matrix [7][7]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 && i != 0  {
				fmt.Printf("%5s ", " ")
			} else {
				fmt.Printf("%5d ", matrix[i][j])
			}
		}

		fmt.Println()
	}
}