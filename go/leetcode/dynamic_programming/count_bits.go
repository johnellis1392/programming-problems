package dynamicprogramming

import "math"

/*
0 0
1 1
2 10
3 11
4 100
5 101
6 110
7 111
8 1000
9 1001
10 1010
11 1011
12 1100
13 1101
14 1110
15 1111
16 1 0000
*/

func countBits(n int) []int {
	switch n {
	case 0:
		return []int{0}
	}
	dp := make([]int, n+1)
	dp[0] = 0
	checkpoint := 1
	mask := checkpoint - 1

	for i := 1; i <= n; i++ {
		if i == checkpoint*2 {
			dp[i] = 1
			checkpoint = checkpoint * 2
			mask = checkpoint - 1
		} else {
			dp[i] = 1 + dp[mask&i]
		}
	}

	return dp
}

func countBits_OLD(n int) []int {
	switch n {
	case 0:
		return []int{0}
	case 1:
		return []int{0, 1}
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 2; i <= n; i++ {
		bracket := math.Floor(math.Log2(float64(i)))
		mask := math.Round(math.Pow(2.0, bracket)) - 1
		dp[i] = dp[i&int(mask)] + 1
	}
	return dp
}
