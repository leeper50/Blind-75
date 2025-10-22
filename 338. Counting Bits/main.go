// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 42.95%

package main

import (
	"fmt"
	"strconv"
)

func countBits(n int) []int {
	array := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		array[i] = array[i/2] + (i % 2)
	}
	return array
}

func binary(i int) string {
	return strconv.FormatInt(int64(i), 2)
}

func main() {
	for i := range 9 {
		fmt.Printf("%v\t%v\n", binary(i), countBits(i))
	}
}
