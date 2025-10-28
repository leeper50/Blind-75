// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 60.94%

package main

import "fmt"

func rotate90(matrix [][]int) {
	n := len(matrix)
	for i := range matrix {
		// Swap values
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		// Flip each row
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

func rotate180(matrix [][]int) {
	n := len(matrix)
	for i := range matrix {
		// Flip each row
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	// Flip cols
	for i := 0; i < n; i++ {
		matrix[i], matrix[n-1] = matrix[n-1], matrix[i]
		n--
	}
}

func rotate270(matrix [][]int) {
	n := len(matrix)
	for i := range matrix {
		// Swap values
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// Flip cols
	for i := 0; i < n; i++ {
		matrix[i], matrix[n-1] = matrix[n-1], matrix[i]
		n--
	}
}

func main() {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate90(data)
	fmt.Printf("%v\n", data)
	rotate90(data)
	fmt.Printf("%v\n", data)
	rotate90(data)
	fmt.Printf("%v\n", data)
	rotate90(data)
	fmt.Printf("%v\n", data)
	data = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate90(data)
	fmt.Printf("%v\n", data)
	data = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate180(data)
	fmt.Printf("%v\n", data)
	data = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate270(data)
	fmt.Printf("%v\n", data)
}
