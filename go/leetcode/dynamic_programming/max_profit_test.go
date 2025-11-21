package dynamicprogramming

import "testing"

func TestMaxProfit(t *testing.T) {
	assertEquals := func(expected, actual int) {
		if expected != actual {
			t.Errorf("Failed: %v != %v", actual, expected)
		}
	}

	t.Run("Example 1", func(t *testing.T) {
		assertEquals(5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	})

	t.Run("Example 2", func(t *testing.T) {
		assertEquals(0, maxProfit([]int{7, 6, 4, 3, 1}))
	})
}
