package main

import (
	"errors"
	"fmt"
	"math"
)

// Prints vector or matrix
func vm_print(s interface{}) {
	switch v := s.(type) {
	case []float64:
		for _, val := range v {
			fmt.Printf("%.2f ", val)
		}
		fmt.Println()
	case [][]float64:
		for _, row := range v {
			for _, val := range row {
				fmt.Printf("%.2f ", val)
			}
			fmt.Println()
		}
	}
}

// Gaussian Elimination
// Input: Augmented matrix (A|b)
// Output: Solution vector, final matrix state, error
func gaussian_elimination(input_matrix [][]float64) ([]float64, [][]float64, error) {
	n := len(input_matrix)
	m := len(input_matrix[0]) - 1

	// Make a copy of the matrix to modify
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, len(input_matrix[i]))
		copy(matrix[i], input_matrix[i])
	}

	// Forward elimination with partial pivoting
	for col := range m {
		maxRow := col
		for row := col + 1; row < n; row++ {
			if math.Abs(matrix[row][col]) > math.Abs(matrix[maxRow][col]) {
				maxRow = row
			}
		}

		if math.Abs(matrix[maxRow][col]) < 1e-10 {
			return nil, matrix, errors.New("Error: Last element is zero")
		}

		// Swap current row with maxRow
		matrix[col], matrix[maxRow] = matrix[maxRow], matrix[col]

		// Eliminate current column in rows below
		for row := col + 1; row < n; row++ {
			factor := matrix[row][col] / matrix[col][col]
			for k := col; k <= m; k++ {
				matrix[row][k] -= factor * matrix[col][k]
			}
		}
	}

	// Back substitution
	solution := make([]float64, m)
	for row := m - 1; row >= 0; row-- {
		solution[row] = matrix[row][m]
		for col := row + 1; col < m; col++ {
			solution[row] -= matrix[row][col] * solution[col]
		}
		solution[row] /= matrix[row][row]
	}

	return solution, matrix, nil
}

func main() {
	// Matrixes used for gaussian_elimination
	var matrix_1 = [][]float64{
		{0, 1, 1},
		{0, 1, 2},
		{0, 2, 2},
	}

	var matrix_2 = [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	var matrix_3 = [][]float64{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}

	var matrix_4 = [][]float64{
		{1, 2, 3},
		{4, 12, 20},
	}

	var matrix_5 = [][]float64{
		{1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4},
		{0, 0, 1, 2, 3},
		{0, 0, 0, 1, 2},
	}

	// Verification of failure cases:
	// The coefficient matrix A should fail immediately with the matrix unchanged (for any b).
	var solution, finalMatrix, error = gaussian_elimination(matrix_1)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf("Returned Solution: ")
		vm_print(solution)
	}
	fmt.Println("Final matrix:")
	vm_print(finalMatrix)
	fmt.Println()

	// The coefficient matrix A should fail on it’s final check, with the resulting matrix (for any b)
	solution, finalMatrix, error = gaussian_elimination(matrix_2)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf("Returned Solution: ")
		vm_print(solution)
	}
	fmt.Println("Final matrix:")
	vm_print(finalMatrix)
	fmt.Println()

	// The coefficient matrix A should fail when seeking it’s third pivot,
	// with the resulting matrix (for any b)
	solution, finalMatrix, error = gaussian_elimination(matrix_3)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf("Returned Solution: ")
		vm_print(solution)
	}
	fmt.Println("Final matrix:")
	vm_print(finalMatrix)
	fmt.Println()

	// Verification of successful implementation
	// The augmented matrix A∗ should find the solution (−1, 2) with ending matrix
	solution, finalMatrix, error = gaussian_elimination(matrix_4)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf("Returned Solution: ")
		vm_print(solution)
	}
	fmt.Println("Final matrix:")
	vm_print(finalMatrix)
	fmt.Println()

	// The augmented matrix A∗ should find the solution (0, 0, −1, 2) with the matrix unchanged.
	solution, finalMatrix, error = gaussian_elimination(matrix_5)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf("Returned Solution: ")
		vm_print(solution)
	}
	fmt.Println("Final matrix:")
	vm_print(finalMatrix)
	fmt.Println()
}
