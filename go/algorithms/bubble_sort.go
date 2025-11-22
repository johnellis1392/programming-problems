package algorithms

import "cmp"

func BubbleSort[T cmp.Ordered](s []T) {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				temp := s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}
}
