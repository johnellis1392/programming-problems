package dynamicprogramming

import "testing"

func TestPascalsTriangle2(t *testing.T) {
	assertArraysEqual := func(a, b []int) {
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
		assertArraysEqual([]int{1, 3, 3, 1}, getRow(3))
	})

	t.Run("Example 2", func(t *testing.T) {
		assertArraysEqual([]int{1}, getRow(0))
	})

	t.Run("Example 3", func(t *testing.T) {
		assertArraysEqual([]int{1, 1}, getRow(1))
	})
}
