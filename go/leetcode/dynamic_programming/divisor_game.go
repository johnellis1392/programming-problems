package dynamicprogramming

import "math"

// 0ms, beats 100%
// 3.93MB, beats 12.68%
func divisorGame1(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false
	dp[1] = false
	for i := 2; i <= n; i++ {
		if dp[i-1] == false {
			dp[i] = true
		}
		for d := 2; d <= i/d; d++ {
			if i%d != 0 {
				continue
			}
			q := i / d
			if dp[i-d] == false || dp[i-q] == false {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// 11732 ns/op
func divisorGame2(n int) bool {
	brackets := int(math.Ceil(float64(n) / 8))
	dp := make([]byte, brackets)
	dp[0] |= 0b1

	getBit := func(i int) bool {
		return dp[i/8]&(0b1<<(i%8)) > 0
	}

	setBit := func(i int, b bool) {
		v := byte(0b1 << (i % 8))
		if b {
			dp[i/8] |= v
		} else {
			dp[i/8] &= ^v
		}
	}

	for i := 2; i <= n; i++ {
		if getBit(i-1) == false {
			dp[i/8] |= (0b1 << (i % 8))
		}
		for d := 2; d <= i/d; d++ {
			if i%d != 0 {
				continue
			}
			q := i / d
			if getBit(d) || getBit(q) {
				setBit(i, true)
			}
		}
	}
	return getBit(n)
}

// 1.815 ns/op (obviously)
func divisorGame(n int) bool {
	return n%2 == 0
}
