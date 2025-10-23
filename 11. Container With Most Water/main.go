// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 88.02%

package main

func maxAreaSlow(height []int) int {
	left, right := 0, 0
	for i := range height {
		for j := i; j < len(height); j++ {
			m := min(i, j)
			if left*right < m*m {
				left, right = i, j
			}
		}
	}
	return height[left] * height[right]
}

func maxAreaFast(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func main() {
	println(maxAreaSlow([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	println(maxAreaSlow([]int{1, 1}), 1)
	println(maxAreaFast([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	println(maxAreaFast([]int{1, 1}), 1)
}
