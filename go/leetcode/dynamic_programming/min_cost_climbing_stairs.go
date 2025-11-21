package dynamicprogramming

func minCostClimbingStairs_OLD(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func minCostClimbingStairs(cost []int) int {
	switch len(cost) {
	case 1:
		return cost[0]
	case 2:
		return min(cost[0], cost[1])
	}
	a := cost[0]
	b := cost[1]
	for i := 2; i < len(cost); i++ {
		c := min(a, b) + cost[i]
		a = b
		b = c
	}
	return min(a, b)
}
