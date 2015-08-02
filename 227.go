package main

import "fmt"

func solve(matrix [][]float64, result []float64) {
	// There are additional tests we would need to do in a generic solver (such
	// as checking determinant is not 0) but we won't do that here.

	// Initiate lower triangular matrix.
	size := len(matrix)
	lt := make([][]float64, size)
	for i := 0; i < size; i++ {
		lt[i] = make([]float64, size)
	}

	// Split matrix into lower and upper triangular matrices.
	for rowK, rowV := range matrix {
		for colK, colV := range rowV[:rowK] {
			// If the value is 0, no action needs to be done.
			if matrix[colK][colK] != 0 {
				lt[rowK][colK] = colV / matrix[colK][colK]

				// Multiply all remaining values in row by multiplier.
				for i := colK; i < size; i++ {
					// Note colK determines the row we are multiplying by.
					matrix[rowK][i] -= lt[rowK][colK] * matrix[colK][i]
				}
			}
		}
	}

	// Set diagonal of lower triangle to 1.
	for rowK, rowV := range lt {
		rowV[rowK] = 1
	}

	// Solve for y (result) using lower half.
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			result[i] -= result[j] * lt[i][j]
		}
	}

	// Solve for answer (result) using upper half.
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j > i; j-- {
			result[i] -= result[j] * matrix[i][j]
		}
		result[i] /= matrix[i][i]
	}
}

func mod(i int) int {
	if i < 0 {
		return -i
	}
	if i > 50 {
		return 100 - i
	}
	return i
}

func main() {
	matrix := make([][]float64, 51)
	for i := 0; i < 51; i++ {
		matrix[i] = make([]float64, 51)
	}

	result := make([]float64, 51)
	for i := 1; i < 51; i++ {
		result[i] = 36
		matrix[i][i] = 18
		matrix[i][mod(i-1)] -= 8
		matrix[i][mod(i+1)] -= 8
		matrix[i][mod(i-2)] -= 1
		matrix[i][mod(i+2)] -= 1
	}
	matrix[0][0] = 1
	result[0] = 0

	solve(matrix, result)

	fmt.Printf("%10f\n", result[50])
}
