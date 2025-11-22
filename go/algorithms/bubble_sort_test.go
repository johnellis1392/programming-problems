package algorithms

import (
	"math/rand"
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort test 1", func(t *testing.T) {
		input := []int{3, 2, 1}
		expected := []int{1, 2, 3}
		BubbleSort(input)
		if !slices.Equal(expected, input) {
			t.Errorf("Failed to sort: %v", input)
		}
	})

	t.Run("Benchmark Bubblesort", func(t *testing.T) {
		n := 30
		input := make([]int, n)
		for i := range n {
			input[i] = rand.Intn(100)
		}
		for range 10 {
			rand.Shuffle(n, func(i, j int) {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			})
			BubbleSort(input)
			if !slices.IsSorted(input) {
				t.Errorf("Failed to sort: %v", input)
			}
		}
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("Benchmark Bubblesort", func(b *testing.B) {
		n := 30
		input := make([]int, n)
		for i := range n {
			input[i] = rand.Intn(100)
		}
		for b.Loop() {
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			})
			b.StartTimer()
			BubbleSort(input)
			if !slices.IsSorted(input) {
				b.Errorf("Failed to sort: %v", input)
			}
		}
	})
}
