package dynamicprogramming

import "testing"

func TestCountBits(t *testing.T) {
	assertEquals := func(a, b []int) {
		if len(a) != len(b) {
			t.Errorf("a != b\n- len(a) = %v\n- len(b) = %v", len(a), len(b))
			return
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				t.Errorf("a[%d] != b[%d]\n- a = %v\n- b = %v", i, i, a, b)
				return
			}
		}
	}

	t.Run("Example 1", func(t *testing.T) {
		assertEquals([]int{0, 1, 1}, countBits(2))
	})

	t.Run("Example 2", func(t *testing.T) {
		assertEquals([]int{0, 1, 1, 2, 1, 2}, countBits(5))
	})

	t.Run("Example 3", func(t *testing.T) {
		assertEquals([]int{0, 1, 1, 2, 1, 2, 2, 3}, countBits(7))
	})

	t.Run("Example 4", func(t *testing.T) {
		assertEquals([]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1}, countBits(16))
	})
}
