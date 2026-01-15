package main

import "fmt"

func main() {
	var matrix [8][8]int

	matrix = initMatrix(matrix)

	displayMatrix(matrix)
}

func initMatrix(matrix [8][8]int) [8][8]int {
	counter := 0
	colSum := 0
	rowSum := make([]int, 8)
	total := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == len(matrix[i])-1 {
				if i == len(matrix)-1 {
					matrix[i][j] = total
					} else {
					matrix[i][j] = colSum
				}
			} else if i == len(matrix)-1 {
				matrix[i][j] = rowSum[j]
				total += rowSum[j]
			} else {
				matrix[i][j] = counter
				colSum += counter
				rowSum[i] += counter
				counter++
			}

		}
		counter = i + 1
		colSum = 0
		total = 0
	}

	return matrix
}

func displayMatrix(matrix [8][8]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			fmt.Printf("%5d ", matrix[i][j])

		}

		fmt.Println()
	}
}
