// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 49.28%

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	rowlen, collen := len(matrix[0]), len(matrix)
	result := make([]int, rowlen*collen)
	tb, bb, lb, rb := 0, collen-1, 0, rowlen-1
	op, coord := 0, 0
	for tb <= bb && lb <= rb {
		switch op {
		case 0: // Forwards
			for i := lb; i <= rb; i++ {
				result[coord] = matrix[tb][i]
				coord++
			}
			tb++
		case 1: // Down
			for i := tb; i <= bb; i++ {
				result[coord] = matrix[i][rb]
				coord++
			}
			rb--
		case 2: // Backwards
			for i := rb; i >= lb; i-- {
				result[coord] = matrix[bb][i]
				coord++
			}
			bb--
		case 3: // Up
			for i := bb; i >= tb; i-- {
				result[coord] = matrix[i][lb]
				coord++
			}
			lb++
		}
		op = (op + 1) % 4
	}
	return result
}

func main() {
	var data [][]int
	data = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("%v == [1 2 3 6 9 8 7 4 5]\n", spiralOrder(data))

	data = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Printf("%v == [1 2 3 4 8 12 11 10 9 5 6 7]\n", spiralOrder(data))

	data = [][]int{{1}}
	fmt.Printf("%v == [1]\n", spiralOrder(data))

	data = [][]int{{1, 2, 3, 4, 5}}
	fmt.Printf("%v == [1 2 3 4 5]\n", spiralOrder(data))

	data = [][]int{{1}, {2}, {3}, {4}, {5}}
	fmt.Printf("%v == [1 2 3 4 5]\n", spiralOrder(data))

	data = [][]int{{1, 2}, {3, 4}}
	fmt.Printf("%v == [1 2 4 3]\n", spiralOrder(data))
}
