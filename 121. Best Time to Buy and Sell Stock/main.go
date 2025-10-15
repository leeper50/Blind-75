// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 85.34%

package main

func maxProfit(prices []int) int {
	profit, lowestPrice := 0, prices[0]
	for _, currentPrice := range prices {
		if currentPrice < lowestPrice {
			lowestPrice = currentPrice
		}
		if profit < currentPrice-lowestPrice {
			profit = currentPrice - lowestPrice
		}
	}
	return profit
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	println(maxProfit([]int{7, 6, 4, 3, 1}) == 0)
	println(maxProfit([]int{1}) == 0)
}
