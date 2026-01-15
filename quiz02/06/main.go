package main

import "fmt"

func main() {
	var matrix [5][5]int


	matrix = initMatrix(matrix)

	displayMatrix(matrix)
}

func initMatrix(matrix [5][5]int) [5][5]int {
	value := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = value
			} else if i < j {
				matrix[i][j] = 10
			} else {
				matrix[i][j] = 20
			}
			// matrix[i][j] = value
		}
		value++
	}

	return matrix
}

func displayMatrix(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%5d ", matrix[i][j])
		}

		fmt.Println()
	}
}