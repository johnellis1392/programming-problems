package dynamicprogramming

func getRow(rowIndex int) []int {
	dp := make([]int, rowIndex+1)
	dp[0] = 1
	for r := 1; r <= rowIndex; r++ {
		dp[r] = 1
		for c := r - 1; c >= 1; c-- {
			dp[c] = dp[c] + dp[c-1]
		}
	}
	return dp
}
