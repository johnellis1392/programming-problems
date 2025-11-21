package dynamicprogramming

func ClimbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}
	dp := make([]int, n+1)

	// Base Cases
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
