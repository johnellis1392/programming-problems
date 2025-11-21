package dynamicprogramming

func maxProfit(prices []int) int {
	n := len(prices)
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(prices[i]-minPrice, maxProfit)
	}
	return maxProfit
}
