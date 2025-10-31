// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 75.26%

package main

import "fmt"

func setZeroesBig(matrix [][]int) {
	// the bad one
	// create map to store indexes
	zeroMap := make(map[int][]int)

	// Loop through and collect indices with 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zeroMap[i] = append(zeroMap[i], j)
			}
		}
	}
	// Loop through index map
	for i, l := range zeroMap {
		// 0 out rows and columns
		// Row
		for k := range matrix[0] {
			matrix[i][k] = 0
		}
		for _, j := range l {
			// Column
			for k := range matrix {
				matrix[k][j] = 0
			}
		}
	}
}

func setZeroesSmall(matrix [][]int) {
	// the small one
	zeroInFirstColumn := false
	for row := range matrix {
		// Check if there is a 0 somewhere in the first column
		if matrix[row][0] == 0 {
			zeroInFirstColumn = true
		}
		// If 0 is anywhere in array
		// Overwrite first element of row and column with 0
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}
	// Iterate through matrix backwards,
	// setting each element to 0 if row/column are set to 0
	for row := len(matrix) - 1; row >= 0; row-- {
		for col := len(matrix[0]) - 1; col >= 1; col-- {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
		// Update first column if 0 was present from first scan
		if zeroInFirstColumn {
			matrix[row][0] = 0
		}
	}
}

func main() {
	var data [][]int

	data = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroesBig(data)
	fmt.Printf("%v == [[1 0 1] [0 0 0] [1 0 1]]\n", data)

	data = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroesBig(data)
	fmt.Printf("%v == [[0 0 0 0] [0 4 5 0] [0 3 1 0]]\n", data)

	data = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroesSmall(data)
	fmt.Printf("%v == [[1 0 1] [0 0 0] [1 0 1]]\n", data)

	data = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroesSmall(data)
	fmt.Printf("%v == [[0 0 0 0] [0 4 5 0] [0 3 1 0]]\n", data)
}
