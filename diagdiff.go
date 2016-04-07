package main

import "fmt"

func main() {
	var size int

	matrix := [][]int{}

	fmt.Scanf("%v", &size)

	for i := 0; i < size; i++ {
		matrix = append(matrix, []int{})
	}

	var value int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scanf("%v", &value)
			matrix[i] = append(matrix[i], value)
		}
	}

	rightDiagonalSum := 0
	leftDiagonalSum := 0

	for row, col := 0, 0; row < size && col < size; row, col = row+1, col+1 {
		rightDiagonalSum += matrix[row][col]
	}

	for row, col := 0, size-1; row < size && col >= 0; row, col = row+1, col-1 {
		leftDiagonalSum += matrix[row][col]
	}

	fmt.Println(abs(rightDiagonalSum - leftDiagonalSum))
}

func abs(value int) int {
	if value < 0 {
		value *= -1
	}

	return value
}
